from pysword.modules import SwordModules
import argparse
import html
import json
import os
import re
import sys
import xml.etree.ElementTree as ET

if sys.version_info > (3, 0):
    xrange = range


# These tags represent structural/layout boundaries in OSIS.
# They should create word boundaries in plain text.
BOUNDARY_TAGS = {
    # Anonymous block: generic block/container element.
    "ab",

    # Chapter milestone: marks the start or end of a chapter.
    "chapter",

    # Closing block: closing material such as a letter ending or formal conclusion.
    "closer",

    # Division block: larger document/book/section division.
    "div",

    # List item: item inside an OSIS list.
    "item",

    # Line: poetic or line-based text unit.
    "l",

    # Line break: explicit inline line break.
    "lb",

    # Line group: group of poetic or line-based text units.
    "lg",

    # List block: container for list items.
    "list",

    # Generic milestone: empty marker used for starts, ends, or structural positions.
    "milestone",

    # Milestone end: explicit end marker for a milestone range.
    "milestoneEnd",

    # Milestone start: explicit start marker for a milestone range.
    "milestoneStart",

    # Paragraph: normal paragraph boundary.
    "p",

    # Quotation: quoted speech or quotation block/inline wrapper.
    "q",

    # Table row: row within an OSIS table.
    "row",

    # Salutation: greeting/opening formula, often in epistles.
    "salute",

    # Signature: signed/authorship closing material.
    "signed",

    # Speech: spoken discourse container.
    "speech",

    # Speaker: identifies the person speaking.
    "speaker",

    # Table: tabular text container.
    "table",

    # Verse milestone: marks the start or end of a verse.
    "verse",

    # Abbreviation: shortened form of a word or title.
    "abbr",

    # Date: date expression or calendar reference.
    "date",

    # Divine name: special wrapper for names/titles of God.
    "divineName",

    # Foreign text: word or phrase in another language.
    "foreign",

    # Highlighted text: emphasis, italics, small caps, or similar rendering.
    "hi",

    # Mentioned term: word/phrase being referenced as a term, not merely used normally.
    "mentioned",

    # Name: personal, place, divine, or other named entity.
    "name",

    # Reference: scripture citation, cross-reference, or linked reference text.
    "reference",

    # Segment: generic inline segment wrapper.
    "seg",

    # Translator change: supplied, added, changed, or otherwise translator-marked text.
    "transChange",
}

# These tags are not verse-body text for the plain source output.
# Raw versions are preserved in <version>-osis.json.
REMOVE_TAGS = {
    "note",
}

# These are usually editorial/non-canonical. Keep only if canonical="true".
CONDITIONAL_REMOVE_TAGS = {
    "head",
    "title",
}

# This legacy marker is escaped as text inside the ChiUn SWORD module.
# XML parsing therefore exposes it as literal text rather than an element.
# Keep this list explicit: angle brackets also mark real editorial text in
# translations such as the Sahidic Coptic Bible.
ESCAPED_SOURCE_MARKERS = ("<WAHb>",)


def get_osis_output_file(output_file):
    root, ext = os.path.splitext(output_file)
    return root + "-osis" + ext


def local_name(tag):
    if "}" in tag:
        return tag.split("}", 1)[1]
    return tag


def is_canonical(node):
    return node.attrib.get("canonical", "").lower() == "true"


def append_space(parts):
    if parts and not parts[-1].endswith(" "):
        parts.append(" ")


def walk_osis_node(node, parts):
    tag = local_name(node.tag)

    if tag in REMOVE_TAGS:
        if node.tail:
            parts.append(node.tail)
        return

    if tag in CONDITIONAL_REMOVE_TAGS and not is_canonical(node):
        if node.tail:
            parts.append(node.tail)
        return

    if tag in BOUNDARY_TAGS:
        append_space(parts)

    if node.text:
        parts.append(node.text)

    for child in node:
        walk_osis_node(child, parts)

    if tag in BOUNDARY_TAGS:
        append_space(parts)

    if node.tail:
        parts.append(node.tail)


def strip_tags_fallback(osis_text):
    """
    Fallback for malformed XML fragments.

    This is intentionally simple and only used if ElementTree cannot parse
    the wrapped OSIS fragment.
    """

    text = osis_text

    # Remove notes.
    text = re.sub(r"<note\b[^>]*>.*?</note>", " ", text, flags=re.DOTALL)

    # Remove non-canonical title/head blocks when possible.
    text = re.sub(r"<title\b(?![^>]*canonical=['\"]true['\"])[^>]*>.*?</title>", " ", text, flags=re.DOTALL)
    text = re.sub(r"<head\b(?![^>]*canonical=['\"]true['\"])[^>]*>.*?</head>", " ", text, flags=re.DOTALL)

    # Preserve structural boundaries before stripping tags.
    boundary = (
        "ab|chapter|closer|div|item|l|lb|lg|list|milestone|"
        "milestoneEnd|milestoneStart|p|q|row|salute|signed|"
        "speech|speaker|table|verse"
    )

    text = re.sub(r"<(?:" + boundary + r")\b[^>]*/>", " ", text)
    text = re.sub(r"</?(?:" + boundary + r")\b[^>]*>", " ", text)

    # Strip all remaining tags while preserving their inner text.
    text = re.sub(r"<[^>]+>", "", text)

    return text


def osis_fragment_to_plain_text(osis_text):
    """
    Convert a raw OSIS/SWORD XML fragment into plain verse text.

    This replaces pysword clean=True because clean=True can strip OSIS
    structural tags without preserving word boundaries.

    Raw OSIS is still preserved separately in <version>-osis.json.
    """

    if osis_text is None:
        return ""

    try:
        # pysword returns fragments, not full OSIS documents.
        # Wrap in a fake root so the XML parser can handle it.
        root = ET.fromstring("<root>" + osis_text + "</root>")
        parts = []

        if root.text:
            parts.append(root.text)

        for child in root:
            walk_osis_node(child, parts)

        text = "".join(parts)

    except ET.ParseError:
        text = strip_tags_fallback(osis_text)

    text = html.unescape(text)
    for marker in ESCAPED_SOURCE_MARKERS:
        text = text.replace(marker, "")

    # Normalize whitespace.
    text = re.sub(r"\s+", " ", text).strip()

    # Remove accidental spaces before punctuation.
    text = re.sub(r"\s+([,.;:?!])", r"\1", text)

    # Light cleanup around brackets/quotes.
    text = re.sub(r"([“‘(\[])\s+", r"\1", text)
    text = re.sub(r"\s+([”’)\]])", r"\1", text)

    return text


def generate_dicts(source_file, bible_version):
    modules = SwordModules(source_file)
    modules.parse_modules()
    bible = modules.get_bible_from_module(bible_version)

    books = bible.get_structure()._books["ot"] + bible.get_structure()._books["nt"]

    plain_bib = {"books": []}
    osis_bib = {"books": []}

    for book in books:
        plain_chapters = []
        osis_chapters = []

        for chapter in xrange(1, book.num_chapters + 1):
            plain_verses = []
            osis_verses = []

            for verse in xrange(1, len(book.get_indicies(chapter)) + 1):
                name = book.name + " " + str(chapter) + ":" + str(verse)

                osis_text = bible.get(
                    books=[book.name],
                    chapters=[chapter],
                    verses=[verse],
                    clean=False
                )

                plain_text = osis_fragment_to_plain_text(osis_text)

                plain_verses.append({
                    "verse": verse,
                    "chapter": chapter,
                    "name": name,
                    "text": plain_text
                })

                osis_verses.append({
                    "verse": verse,
                    "chapter": chapter,
                    "name": name,
                    "text": osis_text
                })

            plain_chapters.append({
                "chapter": chapter,
                "name": book.name + " " + str(chapter),
                "verses": plain_verses
            })

            osis_chapters.append({
                "chapter": chapter,
                "name": book.name + " " + str(chapter),
                "verses": osis_verses
            })

        plain_bib["books"].append({
            "name": book.name,
            "chapters": plain_chapters
        })

        osis_bib["books"].append({
            "name": book.name,
            "chapters": osis_chapters
        })

    return plain_bib, osis_bib


def write_json(bible_dict, output_file):
    # Opening with "w" intentionally overwrites the file if it already exists.
    with open(output_file, "w", encoding="utf-8") as outfile:
        json.dump(bible_dict, outfile, indent=4, ensure_ascii=False)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--source_file", required=True)
    parser.add_argument("--bible_version", required=True)
    parser.add_argument("--output_file", required=True)
    parser.add_argument("--osis_output_file")
    args = parser.parse_args()

    osis_output_file = args.osis_output_file
    if not osis_output_file:
        osis_output_file = get_osis_output_file(args.output_file)

    plain_bible_dict, osis_bible_dict = generate_dicts(
        args.source_file,
        args.bible_version
    )

    write_json(plain_bible_dict, args.output_file)
    write_json(osis_bible_dict, osis_output_file)

    print("Wrote source:")
    print(args.output_file)
    print("Wrote OSIS source:")
    print(osis_output_file)


if __name__ == "__main__":
    main()
