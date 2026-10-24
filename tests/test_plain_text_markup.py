import json
from pathlib import Path
import sqlite3
import tempfile
import unittest

import pyarrow.parquet as pq

from scripts.sword_to_json import generate_dicts, osis_fragment_to_plain_text
from scripts.verify_plain_text_markup import findings, unexpected_findings


ROOT = Path(__file__).resolve().parents[1]
SOURCE = ROOT / "sources/zh-hant/ChiUn"


def verse_at(data, book_name, chapter_number, verse_number):
    return next(v for b in data["books"] if b["name"] == book_name
                for c in b["chapters"] if c["chapter"] == chapter_number
                for v in c["verses"] if v["verse"] == verse_number)


class ConversionTests(unittest.TestCase):
    def test_encoded_and_double_encoded_source_markers(self):
        for marker in ("&lt;WAHb&gt;", "&amp;lt;WAHb&amp;gt;"):
            with self.subTest(marker=marker):
                self.assertEqual(osis_fragment_to_plain_text(f"誠心{marker}實意地{marker}"), "誠心實意地")

    def test_issue_126_real_osis_fragment_preserves_chinese_text(self):
        raw = json.loads((SOURCE / "ChiUn-osis.json").read_text())
        source = json.loads((SOURCE / "ChiUn.json").read_text())
        original = verse_at(raw, "Joshua", 24, 14)["text"]
        corrected = verse_at(source, "Joshua", 24, 14)["text"]
        self.assertEqual(original.count("&lt;WAHb&gt;"), 2)
        self.assertEqual(osis_fragment_to_plain_text(original), corrected)
        self.assertIn("誠心實意地", corrected)
        self.assertNotIn("WAHb", corrected)

    def test_editorial_brackets_and_other_literal_text_survive(self):
        for encoded, literal in [
            ("ⲁ &lt;ⲛ&gt; ⲃ", "ⲁ <ⲛ> ⲃ"),
            ("&lt;DEUX APPENDICES &gt;", "<DEUX APPENDICES >"),
            ("&lt;example&gt;", "<example>"),
            ("a &lt; b &amp; c &gt; d", "a < b & c > d"),
        ]:
            with self.subTest(encoded=encoded):
                self.assertEqual(osis_fragment_to_plain_text(encoded), literal)

    def test_existing_word_boundaries_and_note_removal(self):
        self.assertEqual(osis_fragment_to_plain_text('alpha<lb/>beta'), "alpha beta")
        self.assertEqual(osis_fragment_to_plain_text('alpha<note>footnote</note> beta'), "alpha beta")

    def test_unclosed_literal_markers_in_malformed_fragment(self):
        self.assertEqual(osis_fragment_to_plain_text('誠心<WAHb>實意地<WAHb>'), "誠心實意地")

    def test_complete_archive_regeneration_matches_corrected_source(self):
        regenerated, _ = generate_dicts(str(SOURCE / "ChiUn.zip"), "ChiUn")
        committed = json.loads((SOURCE / "ChiUn.json").read_text())
        self.assertEqual(regenerated, committed)


class MarkupGuardTests(unittest.TestCase):
    def test_all_plain_sources_have_no_new_markup(self):
        baseline = json.loads((ROOT / "tests/known_plain_text_markup.json").read_text())
        self.assertFalse(any("WAHb" in item["location"][-1] for item in baseline))
        self.assertFalse(unexpected_findings(findings(ROOT / "sources"), baseline))

    def test_unknown_and_encoded_tags_are_flagged_without_editing(self):
        with tempfile.TemporaryDirectory() as directory:
            base = Path(directory)
            path = base / "en/Test/Test.json"
            path.parent.mkdir(parents=True)
            document = {"books": [{"name": "Genesis", "chapters": [{"chapter": 1, "verses": [
                {"verse": 1, "text": "text &lt;newTag&gt; more <other/> &amp;lt;nestedTag&amp;gt;"},
                {"verse": 2, "text": "ⲁ <ⲛ> ⲃ <DEUX APPENDICES >"},
            ]}]}]}
            original = json.dumps(document)
            path.write_text(original)
            found = findings(base)
            self.assertEqual({key[-1] for key in found}, {"<newTag>", "<other/>", "<nestedTag>"})
            self.assertEqual(sum(unexpected_findings(found, []).values()), 3)
            self.assertEqual(path.read_text(), original)

    def test_existing_findings_cannot_hide_new_locations_or_repetitions(self):
        key = ("en/Test/Test.json", "Genesis", 1, 1, "<note>")
        from collections import Counter
        baseline = [{"location": list(key), "count": 1}]
        self.assertEqual(unexpected_findings(Counter({key: 2}), baseline)[key], 1)
        new_key = (*key[:3], 2, key[-1])
        self.assertTrue(unexpected_findings(Counter({new_key: 1}), baseline))


class PublishedOutputTests(unittest.TestCase):
    def test_all_text_exports_are_corrected(self):
        for folder, suffix in [("csv", "csv"), ("json", "json"), ("md", "md"),
                               ("sql", "sql"), ("txt", "txt"), ("yaml", "yaml")]:
            with self.subTest(folder=folder):
                text = (ROOT / f"formats/{folder}/ChiUn.{suffix}").read_text()
                self.assertNotIn("WAHb", text)
                self.assertIn("誠心實意地", text)

    def test_sqlite_rows_are_corrected_and_database_is_valid(self):
        path = ROOT / "formats/sqlite/ChiUn.db"
        with sqlite3.connect(f"file:{path}?mode=ro", uri=True) as db:
            self.assertEqual(db.execute("PRAGMA integrity_check").fetchone()[0], "ok")
            self.assertEqual(db.execute("SELECT COUNT(*) FROM ChiUn_verses WHERE text LIKE '%WAHb%'").fetchone()[0], 0)
            rows = db.execute("SELECT text FROM ChiUn_verses WHERE book_id=6 AND chapter=24 AND verse=14").fetchall()
            self.assertTrue(rows)
            expected = verse_at(json.loads((SOURCE / "ChiUn.json").read_text()), "Joshua", 24, 14)["text"]
            self.assertTrue(all(text.strip() == expected for (text,) in rows))

    def test_parquet_rows_are_corrected(self):
        table = pq.read_table(ROOT / "formats/parquet/ChiUn.parquet")
        rows = table.to_pylist()
        self.assertFalse(any("WAHb" in row["text"] for row in rows))
        verse = next(row for row in rows if row["book_name"] == "Joshua" and row["chapter"] == 24 and row["verse"] == 14)
        self.assertIn("誠心實意地", verse["text"])


if __name__ == "__main__":
    unittest.main()
