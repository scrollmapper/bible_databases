import os
import json
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

def verify_text_integrity_json(language, translation):
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_dir = os.path.join(base_dir, 'sources', language)
    json_path = os.path.join(source_dir, translation, f"{translation}.json")

    if not os.path.exists(json_path):
        print(f"JSON file {json_path} does not exist.")
        return

    source_data = load_json(json_path)
    target_json_path = os.path.join(base_dir, 'formats', 'json', f"{translation}.json")

    if not os.path.exists(target_json_path):
        print(f"JSON file {target_json_path} does not exist.")
        return

    target_data = load_json(target_json_path)

    differences = []
    source_books = source_data.get('books', [])
    target_books = target_data.get('books', [])

    if len(source_books) != len(target_books):
        differences.append(f"Number of books mismatch: {len(source_books)} (source) vs {len(target_books)} (target)")

    for source_book, target_book in zip(source_books, target_books):
        if source_book['name'] != target_book['name']:
            differences.append(f"Book name mismatch: {source_book['name']} (source) vs {target_book['name']} (target)")

        source_chapters = source_book.get('chapters', [])
        target_chapters = target_book.get('chapters', [])

        if len(source_chapters) != len(target_chapters):
            differences.append(f"Number of chapters in book '{source_book['name']}' mismatch: {len(source_chapters)} (source) vs {len(target_chapters)} (target)")

        for source_chapter, target_chapter in zip(source_chapters, target_chapters):
            if source_chapter['chapter'] != target_chapter['chapter']:
                differences.append(f"Chapter number mismatch in book '{source_book['name']}': {source_chapter['chapter']} (source) vs {target_chapter['chapter']} (target)")

            source_verses = source_chapter.get('verses', [])
            target_verses = target_chapter.get('verses', [])

            if len(source_verses) != len(target_verses):
                differences.append(f"Number of verses in chapter '{source_chapter['chapter']}' of book '{source_book['name']}' mismatch: {len(source_verses)} (source) vs {len(target_verses)} (target)")

            for source_verse, target_verse in zip(source_verses, target_verses):
                if source_verse['verse'] != target_verse['verse']:
                    differences.append(f"Verse number mismatch in chapter '{source_chapter['chapter']}' of book '{source_book['name']}': {source_verse['verse']} (source) vs {target_verse['verse']} (target)")

                # Normalize text for comparison
                source_text = normalize_text(source_verse['text'])
                target_text = normalize_text(target_verse['text'])

                if source_text != target_text:
                    differences.append(f"Verse text mismatch in chapter '{source_chapter['chapter']}' of book '{source_book['name']}':\n{source_text} (source) vs\n{target_text} (target)")

    report_path = f"text_integrity_check_json.txt"
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

    verify_text_integrity_json(language, translation)

if __name__ == "__main__":
    main()
