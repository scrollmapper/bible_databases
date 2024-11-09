import os
import json
import sqlite3

class SQLiteGenerator:
    def __init__(self, source_dir, format_dir):
        self.source_dir = source_dir
        self.format_dir = format_dir

    def generate(self, language, translation):
        data = self.load_json(language, translation)
        translation_name = self.get_readme_title(language, translation)
        license_info = self.get_license_info(language, translation)
        prepared_data = self.prepare_data(data)
        sqlite_path = os.path.join(self.format_dir, 'sqlite', f'{translation}.db')

        # Connect to SQLite database
        conn = sqlite3.connect(sqlite_path)
        cursor = conn.cursor()

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
        for book in prepared_data['books']:
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
        for book_index, book in enumerate(prepared_data['books'], start=1):
            for chapter in book['chapters']:
                for verse in chapter['verses']:
                    cursor.execute(f"""
                    INSERT INTO {translation}_verses (book_id, chapter, verse, text)
                    VALUES (?, ?, ?, ?);
                    """, (book_index, chapter['chapter'], verse['verse'], verse['text']))

        conn.commit()
        conn.close()

        print(f"SQLite database for {translation_name} ({translation}) generated at {sqlite_path}")

    def get_license_info(self, language, translation):
        readme_path = os.path.join(self.source_dir, language, translation, "README.md")
        with open(readme_path, 'r', encoding='utf-8') as file:
            for line in file:
                if line.startswith("**License:**"):
                    return line.split("**License:** ")[1].strip()
        return "Unknown"

    def load_json(self, language, translation):
        json_path = os.path.join(self.source_dir, language, translation, f"{translation}.json")
        with open(json_path, 'r', encoding='utf-8') as file:
            return json.load(file)

    def get_readme_title(self, language, translation):
        readme_path = os.path.join(self.source_dir, language, translation, "README.md")
        with open(readme_path, 'r', encoding='utf-8') as file:
            return file.readline().strip()

    def prepare_data(self, data):
        # This method should prepare and return the data in the required format
        return data
