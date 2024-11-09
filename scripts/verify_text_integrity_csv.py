import os
import json
import csv
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

def load_csv(file_path):
    data = {}
    with open(file_path, 'r', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            book = row['Book']
            chapter = int(row['Chapter'])
            verse = int(row['Verse'])
            text = row['Text']

            if book not in data:
                data[book] = {}
            if chapter not in data[book]:
                data[book][chapter] = []
            data[book][chapter].append({
                'verse': verse,
                'text': text
            })
    return data

def verify_text_integrity_csv(language, translation):
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_dir = os.path.join(base_dir, 'sources', language)
    json_path = os.path.join(source_dir, translation, f"{translation}.json")

    if not os.path.exists(json_path):
        print(f"JSON file {json_path} does not exist.")
        return

    source_data = load_json(json_path)
    csv_path = os.path.join(base_dir, 'formats', 'csv', f"{translation}.csv")

    if not os.path.exists(csv_path):
        print(f"CSV file {csv_path} does not exist.")
        return

    csv_data = load_csv(csv_path)

    differences = []
    source_books = source_data.get('books', [])
    
    for source_book in source_books:
        book_name = source_book['name']
        if book_name not in csv_data:
            differences.append(f"Book '{book_name}' not found in CSV data.")
            continue

        source_chapters = source_book.get('chapters', [])
        csv_chapters = csv_data[book_name]

        if len(source_chapters) != len(csv_chapters):
            differences.append(f"Number of chapters in book '{book_name}' mismatch: {len(source_chapters)} (source) vs {len(csv_chapters)} (CSV)")

        for source_chapter in source_chapters:
            chapter_number = source_chapter['chapter']
            if chapter_number not in csv_chapters:
                differences.append(f"Chapter '{chapter_number}' in book '{book_name}' not found in CSV data.")
                continue

            source_verses = source_chapter.get('verses', [])
            csv_verses = csv_chapters[chapter_number]

            if len(source_verses) != len(csv_verses):
                differences.append(f"Number of verses in chapter '{chapter_number}' of book '{book_name}' mismatch: {len(source_verses)} (source) vs {len(csv_verses)} (CSV)")

            for source_verse, csv_verse in zip(source_verses, csv_verses):
                if source_verse['verse'] != csv_verse['verse']:
                    differences.append(f"Verse number mismatch in chapter '{chapter_number}' of book '{book_name}': {source_verse['verse']} (source) vs {csv_verse['verse']} (CSV)")

                # Normalize text for comparison
                source_text = normalize_text(source_verse['text'])
                csv_text = normalize_text(csv_verse['text'])

                if source_text != csv_text:
                    differences.append(f"Verse text mismatch in chapter '{chapter_number}' of book '{book_name}':\n{source_text} (source) vs\n{csv_text} (CSV)")

    report_path = f"text_integrity_check_csv.txt"
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

    verify_text_integrity_csv(language, translation)

if __name__ == "__main__":
    main()
