import os
import json
import mysql.connector
import unicodedata
from mysql.connector import Error

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

def fetch_mysql_data(cursor, translation):
    books_query = f"SELECT id, name FROM {translation}_books;"
    cursor.execute(books_query)
    books = cursor.fetchall()

    translation_data = {'books': []}
    for book_id, book_name in books:
        chapters_query = f"SELECT chapter, verse, text FROM {translation}_verses WHERE book_id = {book_id};"
        cursor.execute(chapters_query)
        verses = cursor.fetchall()

        chapters_dict = {}
        for chapter, verse, text in verses:
            if chapter not in chapters_dict:
                chapters_dict[chapter] = {'chapter': chapter, 'verses': []}
            chapters_dict[chapter]['verses'].append({
                'verse': verse,
                'text': text
            })

        chapters_list = [chapters_dict[chapter] for chapter in sorted(chapters_dict.keys())]
        translation_data['books'].append({
            'name': book_name,
            'chapters': chapters_list
        })

    return translation_data

def verify_text_integrity_mysql(language, translation, db_host, db_name, db_user, db_password):
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_dir = os.path.join(base_dir, 'sources', language)
    json_path = os.path.join(source_dir, translation, f"{translation}.json")

    if not os.path.exists(json_path):
        print(f"JSON file {json_path} does not exist.")
        return

    source_data = load_json(json_path)

    try:
        connection = mysql.connector.connect(
            host=db_host,
            database=db_name,
            user=db_user,
            password=db_password
        )

        if connection.is_connected():
            cursor = connection.cursor()
            mysql_data = fetch_mysql_data(cursor, translation)

            differences = []
            source_books = source_data.get('books', [])
            mysql_books = mysql_data.get('books', [])

            if len(source_books) != len(mysql_books):
                differences.append(f"Number of books mismatch: {len(source_books)} (source) vs {len(mysql_books)} (MySQL)")

            for source_book, mysql_book in zip(source_books, mysql_books):
                if source_book['name'] != mysql_book['name']:
                    differences.append(f"Book name mismatch: {source_book['name']} (source) vs {mysql_book['name']} (MySQL)")

                source_chapters = source_book.get('chapters', [])
                mysql_chapters = mysql_book.get('chapters', [])

                if len(source_chapters) != len(mysql_chapters):
                    differences.append(f"Number of chapters in book '{source_book['name']}' mismatch: {len(source_chapters)} (source) vs {len(mysql_chapters)} (MySQL)")

                for source_chapter, mysql_chapter in zip(source_chapters, mysql_chapters):
                    if source_chapter['chapter'] != mysql_chapter['chapter']:
                        differences.append(f"Chapter number mismatch in book '{source_book['name']}': {source_chapter['chapter']} (source) vs {mysql_chapter['chapter']} (MySQL)")

                    source_verses = source_chapter.get('verses', [])
                    mysql_verses = mysql_chapter.get('verses', [])

                    if len(source_verses) != len(mysql_verses):
                        differences.append(f"Number of verses in chapter '{source_chapter['chapter']}' of book '{source_book['name']}' mismatch: {len(source_verses)} (source) vs {len(mysql_verses)} (MySQL)")

                    for source_verse, mysql_verse in zip(source_verses, mysql_verses):
                        if source_verse['verse'] != mysql_verse['verse']:
                            differences.append(f"Verse number mismatch in chapter '{source_chapter['chapter']}' of book '{source_book['name']}': {source_verse['verse']} (source) vs {mysql_verse['verse']} (MySQL)")

                        # Normalize text for comparison
                        source_text = normalize_text(source_verse['text'])
                        mysql_text = normalize_text(mysql_verse['text'])

                        if source_text != mysql_text:
                            differences.append(f"Verse text mismatch in chapter '{source_chapter['chapter']}' of book '{source_book['name']}':\n{source_text} (source) vs\n{mysql_text} (MySQL)")

            report_path = f"text_integrity_check_mysql.txt"
            with open(report_path, 'w', encoding='utf-8') as report_file:
                if differences:
                    report_file.write("\n".join(differences))
                    print(f"Text integrity check completed. See {report_path} for details.")
                else:
                    report_file.write("Text integrity check completed successfully. No mismatches found.")
                    print("Text integrity check completed successfully. All texts are consistent.")

        else:
            print("Failed to connect to the database")

    except Error as e:
        print(f"Error: {e}")

    finally:
        if 'connection' in locals() and connection.is_connected():
            cursor.close()
            connection.close()
            print("MySQL connection is closed")

def main():
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    sources_dir = os.path.join(base_dir, 'sources')
    languages = [d for d in os.listdir(sources_dir) if os.path.isdir(os.path.join(sources_dir, d)) and d != "extras"]
    print("Choose your language:")
    language = list_options(languages, "Enter the number corresponding to your language: ")

    translations = [d for d in os.listdir(os.path.join(sources_dir, language)) if os.path.isdir(os.path.join(sources_dir, language, d))]
    print(f"Choose your translation for {language}:")
    translation = list_options(translations, "Enter the number corresponding to your translation: ")

    # Prompt for database connection details
    db_host = input("Enter the MySQL host (default: 'localhost'): ").strip() or 'localhost'
    db_name = input("Enter the MySQL database name: ").strip()
    db_user = input("Enter the MySQL username: ").strip()
    db_password = input("Enter the MySQL password: ").strip()

    verify_text_integrity_mysql(language, translation, db_host, db_name, db_user, db_password)

if __name__ == "__main__":
    main()
