import os
import json
import re
import unicodedata

def normalize_text(text):
    # Replace common characters
    text = text.replace("Æ", "'")
    
    # Unicode normalization
    text = unicodedata.normalize('NFKD', text)
    
    return text

def list_options(options, prompt):
    if len(options) == 1:
        return options[0]
    for i, option in enumerate(options, 1):
        print(f"{i}. {option}")
    choice = input(prompt)
    if choice.isdigit():
        choice = int(choice) - 1
        return options[choice]
    return choice  # Assume the input is the option itself


def load_json(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        return json.load(file)

def load_markdown(file_path):
    data = {}
    book_name = ''
    chapter_number = 0
    
    with open(file_path, 'r', encoding='utf-8') as mdfile:
        for line in mdfile:
            book_match = re.match(r"^## (.+)", line)
            if book_match:
                book_name = book_match.group(1)
                data[book_name] = {}
                continue

            chapter_match = re.match(r"^### Chapter (\d+)", line)
            if chapter_match:
                chapter_number = int(chapter_match.group(1))
                data[book_name][chapter_number] = []
                continue

            verse_match = re.match(r"^\*\*\[(\d+):(\d+)\]\*\* (.+)$", line)
            if verse_match:
                verse_number = int(verse_match.group(2))
                verse_text = verse_match.group(3)
                data[book_name][chapter_number].append({
                    'verse': verse_number,
                    'text': verse_text
                })

    return data

def verify_text_integrity_markdown(language, translation):
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_dir = os.path.join(base_dir, 'sources', language)
    json_path = os.path.join(source_dir, translation, f"{translation}.json")

    if not os.path.exists(json_path):
        print(f"JSON file {json_path} does not exist.")
        return

    source_data = load_json(json_path)
    markdown_path = os.path.join(base_dir, 'formats', 'md', f"{translation}.md")

    if not os.path.exists(markdown_path):
        print(f"Markdown file {markdown_path} does not exist.")
        return

    markdown_data = load_markdown(markdown_path)

    differences = []
    source_books = source_data.get('books', [])
    
    for source_book in source_books:
        book_name = source_book['name']
        if book_name not in markdown_data:
            differences.append(f"Book '{book_name}' not found in Markdown data.")
            continue

        source_chapters = source_book.get('chapters', [])
        markdown_chapters = markdown_data[book_name]

        if len(source_chapters) != len(markdown_chapters):
            differences.append(f"Number of chapters in book '{book_name}' mismatch: {len(source_chapters)} (source) vs {len(markdown_chapters)} (Markdown)")

        for source_chapter in source_chapters:
            chapter_number = source_chapter['chapter']
            if chapter_number not in markdown_chapters:
                differences.append(f"Chapter '{chapter_number}' in book '{book_name}' not found in Markdown data.")
                continue

            source_verses = source_chapter.get('verses', [])
            markdown_verses = markdown_chapters[chapter_number]

            if len(source_verses) != len(markdown_verses):
                differences.append(f"Number of verses in chapter '{chapter_number}' of book '{book_name}' mismatch: {len(source_verses)} (source) vs {len(markdown_verses)} (Markdown)")

            for source_verse, markdown_verse in zip(source_verses, markdown_verses):
                if source_verse['verse'] != markdown_verse['verse']:
                    differences.append(f"Verse number mismatch in chapter '{chapter_number}' of book '{book_name}': {source_verse['verse']} (source) vs {markdown_verse['verse']} (Markdown)")

                # Normalize text for comparison
                source_text = normalize_text(source_verse['text'])
                markdown_text = normalize_text(markdown_verse['text'])

                if source_text != markdown_text:
                    differences.append(f"Verse text mismatch in chapter '{chapter_number}' of book '{book_name}':\n{source_text} (source) vs\n{markdown_text} (Markdown)")

    report_path = f"text_integrity_check_markdown.txt"
    with open(report_path, 'w', encoding='utf-8') as report_file:
        if differences:
            report_file.write("\n".join(differences))
            print(f"Text integrity check completed. See {report_path} for details.")
        else:
            report_file.write("Text integrity check completed successfully. No mismatches found.")
            print("Text integrity check completed successfully. All texts are consistent.")

def main():
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    sources_dir = os.path.join(base_dir, 'sources')
    languages = [d for d in os.listdir(sources_dir) if os.path.isdir(os.path.join(sources_dir, d)) and d != "extras"]
    print("Choose your language:")
    language = list_options(languages, "Enter the number corresponding to your language: ")

    translations = [d for d in os.listdir(os.path.join(sources_dir, language)) if os.path.isdir(os.path.join(sources_dir, language, d))]
    print(f"Choose your translation for {language}:")
    translation = list_options(translations, "Enter the number corresponding to your translation: ")

    verify_text_integrity_markdown(language, translation)

if __name__ == "__main__":
    main()
