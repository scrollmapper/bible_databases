"""Reject new XML-like markup in plain sources without deleting bracketed text.

Raw *-osis.json and ZIP archives intentionally retain original markup.
Existing unrelated findings are recorded explicitly for separate review.
"""
import argparse
from collections import Counter
import html
import json
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parents[1]
# Require actual XML-like syntax; Coptic editorial brackets and headings such
# as <DEUX APPENDICES > are text, not tags with attribute assignments.
TAG = re.compile(r'''</?[A-Za-z][A-Za-z0-9:_.-]*(?:\s+[A-Za-z_:][A-Za-z0-9:_.-]*\s*=\s*(?:"[^"]*"|'[^']*'))*\s*/?>''')


def findings(source_dir):
    result = Counter()
    for path in sorted(Path(source_dir).glob("*/*/*.json")):
        if path.name.endswith("-osis.json"):
            continue
        data = json.loads(path.read_text(encoding="utf-8"))
        for book in data.get("books", []):
            for chapter in book.get("chapters", []):
                for verse in chapter.get("verses", []):
                    decoded = verse["text"]
                    # Inspect nested entity escaping without changing the text.
                    for _ in range(3):
                        unescaped = html.unescape(decoded)
                        if unescaped == decoded:
                            break
                        decoded = unescaped
                    for tag in TAG.findall(decoded):
                        key = (str(path.relative_to(source_dir)), book["name"], chapter["chapter"], verse["verse"], tag)
                        result[key] += 1
    return result


def unexpected_findings(actual, baseline):
    allowed = Counter({tuple(item["location"]): item["count"] for item in baseline})
    return actual - allowed


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--sources", type=Path, default=ROOT / "sources")
    args = parser.parse_args()
    baseline = json.loads((ROOT / "tests/known_plain_text_markup.json").read_text())
    actual = findings(args.sources)
    unexpected = unexpected_findings(actual, baseline)
    if unexpected:
        for location, count in unexpected.items():
            print(f"Unexpected plain-text markup ({count}): {location}")
        raise SystemExit(1)
    print(f"No new plain-text markup leaks. {sum(actual.values())} pre-existing markers remain tracked for separate review.")


if __name__ == "__main__":
    main()
