import sys
import os
import json
import sqlite3

def list_options(options, prompt):
    for i, option in enumerate(options, 1):
        print(f"{i}. {option}")
    choice = int(input(prompt)) - 1
    return options[choice]

def create_sqlite_db(db_path):
    # Create the database file if it doesn't exist
    os.makedirs(os.path.dirname(db_path), exist_ok=True)
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    return conn, cursor

def generate_translation_tables(language, translation, source_directory, cursor):
    json_path = os.path.join(source_directory, language, translation, f"{translation}.json")
    with open(json_path, 'r', encoding='utf-8') as file:
        data = json.load(file)

    readme_path = os.path.join(source_directory, language, translation, "README.md")
    with open(readme_path, 'r', encoding='utf-8') as file:
        translation_name = file.readline().strip()
        license_info = "Unknown"
        for line in file:
            if line.startswith("**License:**"):
                license_info = line.split("**License:** ")[1].strip()

    # Create translations table if it doesn't exist
    cursor.execute("""
    CREATE TABLE IF NOT EXISTS translations (
        translation TEXT PRIMARY KEY,
        title TEXT,
        license TEXT
    );
    """)
    cursor.execute("""
    INSERT OR IGNORE INTO translations (translation, title, license)
    VALUES (?, ?, ?);
    """, (translation, translation_name, license_info))

    # Create books table
    cursor.execute(f"""
    CREATE TABLE IF NOT EXISTS {translation}_books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );
    """)

    # Insert books
    for book in data['books']:
        cursor.execute(f"INSERT INTO {translation}_books (name) VALUES (?);", (book['name'],))

    # Create verses table
    cursor.execute(f"""
    CREATE TABLE IF NOT EXISTS {translation}_verses (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        book_id INTEGER,
        chapter INTEGER,
        verse INTEGER,
        text TEXT,
        FOREIGN KEY (book_id) REFERENCES {translation}_books(id)
    );
    """)

    # Insert verses
    for book_index, book in enumerate(data['books'], start=1):
        for chapter in book['chapters']:
            for verse in chapter['verses']:
                cursor.execute(f"""
                INSERT INTO {translation}_verses (book_id, chapter, verse, text)
                VALUES (?, ?, ?, ?);
                """, (book_index, chapter['chapter'], verse['verse'], verse['text']))

def generate_cross_references(source_directory, cursor):
    json_dir = os.path.join(source_directory, 'extras')

    cursor.execute("""
    CREATE TABLE IF NOT EXISTS cross_references (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        from_book TEXT,
        from_chapter INTEGER,
        from_verse INTEGER,
        to_book TEXT,
        to_chapter INTEGER,
        to_verse_start INTEGER,
        to_verse_end INTEGER,
        votes INTEGER
    );
    """)

    cross_reference_files = [f for f in os.listdir(json_dir) if f.startswith('cross_references') and f.endswith('.json')]
    for file in cross_reference_files:
        with open(os.path.join(json_dir, file), 'r', encoding='utf-8') as jsonfile:
            data = json.load(jsonfile)
        
        for ref in data['cross_references']:
            from_verse = ref['from_verse']
            for to_verse in ref['to_verse']:
                cursor.execute("""
                INSERT INTO cross_references (from_book, from_chapter, from_verse, to_book, to_chapter, to_verse_start, to_verse_end, votes)
                VALUES (?, ?, ?, ?, ?, ?, ?, ?);
                """, (from_verse['book'], from_verse['chapter'], from_verse['verse'], to_verse['book'], to_verse['chapter'], to_verse['verse_start'], to_verse['verse_end'], ref['votes']))

def main():
    # Set base directories relative to the script location
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_directory = os.path.join(base_dir, 'sources')

    # Ask user for the path where the new database should be built
    target_db_path = input("Enter the path where the new SQLite database should be built: ").strip()

    # Step 1: Select Language
    languages = [d for d in os.listdir(source_directory) if os.path.isdir(os.path.join(source_directory, d)) and d != "extras"]
    print("Choose your language:")
    language = list_options(languages, "Enter the number corresponding to your language: ")

    # Step 2: Select Translation
    translations = [d for d in os.listdir(os.path.join(source_directory, language)) if os.path.isdir(os.path.join(source_directory, language, d))]
    print(f"Choose your translation for {language}:")
    translation = list_options(translations, "Enter the number corresponding to your translation: ")

    # Create SQLite database
    conn, cursor = create_sqlite_db(target_db_path)
    
    # Generate translation tables
    generate_translation_tables(language, translation, source_directory, cursor)
    
    # Generate cross references
    generate_cross_references(source_directory, cursor)
    
    # Commit changes and close connection
    conn.commit()
    conn.close()

    print("SQLite database with cross references built successfully!")

if __name__ == "__main__":
    main()
