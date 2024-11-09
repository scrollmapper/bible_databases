import os
import json

class CrossReferencesGenerator:
    def __init__(self, source_dir, format_dir, target_chunk_size=50 * 1024 * 1024):
        self.source_dir = source_dir
        self.format_dir = format_dir
        self.target_chunk_size = target_chunk_size

    def generate(self):
        txt_path = os.path.join(self.source_dir, 'extras', 'cross_references.txt')
        json_dir = os.path.join(self.source_dir, 'extras')
        if not os.path.exists(json_dir):
            os.makedirs(json_dir)

        cross_references = []

        with open(txt_path, 'r', encoding='utf-8') as file:
            for line in file:
                if line.startswith("From Verse"):
                    continue  # Skip header line
                parts = line.strip().split('\t')
                from_verse = self.parse_reference(parts[0])
                to_verses = self.parse_range(parts[1])
                votes = int(parts[2])

                cross_references.append({
                    "from_verse": from_verse,
                    "to_verse": to_verses,
                    "votes": votes
                })
    
        self.split_into_chunks(cross_references, json_dir, 50000)

    def split_into_chunks(self, cross_references, json_dir, entries_per_chunk):
        chunk_index = 0
        for i in range(0, len(cross_references), entries_per_chunk):
            chunk = cross_references[i:i + entries_per_chunk]
            chunk_file = os.path.join(json_dir, f'cross_references_{chunk_index}.json')
            with open(chunk_file, 'w', encoding='utf-8') as jsonfile:
                json.dump({"cross_references": chunk}, jsonfile, ensure_ascii=False, indent=4)
            print(f"Chunk {chunk_index} generated at {chunk_file}")
            chunk_index += 1

    def parse_reference(self, reference):
        parts = reference.split('.')
        return {
            "book": self.expand_book_name(parts[0]),
            "chapter": int(parts[1]),
            "verse": int(parts[2])
        }

    def parse_range(self, range_str):
        ranges = range_str.split('-')
        to_verses = []
        for r in ranges:
            parts = r.split('.')
            book_name = self.expand_book_name(parts[0])
            chapter = int(parts[1])
            verse_start = int(parts[2])
            if len(parts) > 3:
                verse_end = int(parts[3])
            else:
                verse_end = verse_start
            to_verses.append({
                "book": book_name,
                "chapter": chapter,
                "verse_start": verse_start,
                "verse_end": verse_end
            })
        return to_verses

    def expand_book_name(self, abbreviation):
        book_names = {
            "Gen": "Genesis",
            "Exod": "Exodus",
            "Lev": "Leviticus",
            "Num": "Numbers",
            "Deut": "Deuteronomy",
            "Josh": "Joshua",
            "Judg": "Judges",
            "Ruth": "Ruth",
            "1Sam": "1 Samuel",
            "2Sam": "2 Samuel",
            "1Kgs": "1 Kings",
            "2Kgs": "2 Kings",
            "1Chr": "1 Chronicles",
            "2Chr": "2 Chronicles",
            "Ezra": "Ezra",
            "Neh": "Nehemiah",
            "Esth": "Esther",
            "Job": "Job",
            "Ps": "Psalms",
            "Prov": "Proverbs",
            "Eccl": "Ecclesiastes",
            "Song": "Song of Solomon",
            "Isa": "Isaiah",
            "Jer": "Jeremiah",
            "Lam": "Lamentations",
            "Ezek": "Ezekiel",
            "Dan": "Daniel",
            "Hos": "Hosea",
            "Joel": "Joel",
            "Amos": "Amos",
            "Obad": "Obadiah",
            "Jonah": "Jonah",
            "Mic": "Micah",
            "Nah": "Nahum",
            "Hab": "Habakkuk",
            "Zeph": "Zephaniah",
            "Hag": "Haggai",
            "Zech": "Zechariah",
            "Mal": "Malachi",
            "Matt": "Matthew",
            "Mark": "Mark",
            "Luke": "Luke",
            "John": "John",
            "Acts": "Acts",
            "Rom": "Romans",
            "1Cor": "1 Corinthians",
            "2Cor": "2 Corinthians",
            "Gal": "Galatians",
            "Eph": "Ephesians",
            "Phil": "Philippians",
            "Col": "Colossians",
            "1Thess": "1 Thessalonians",
            "2Thess": "2 Thessalonians",
            "1Tim": "1 Timothy",
            "2Tim": "2 Timothy",
            "Titus": "Titus",
            "Phlm": "Philemon",
            "Heb": "Hebrews",
            "Jas": "James",
            "1Pet": "1 Peter",
            "2Pet": "2 Peter",
            "1John": "1 John",
            "2John": "2 John",
            "3John": "3 John",
            "Jude": "Jude",
            "Rev": "Revelation"
        }
        return book_names.get(abbreviation, abbreviation)
