import os
import json

class CrossReferencesGeneratorMySQL:
    def __init__(self, source_dir, format_dir):
        self.source_dir = source_dir
        self.format_dir = format_dir

    def generate(self):
        json_dir = os.path.join(self.source_dir, 'extras')
        sql_dir = os.path.join(self.format_dir, 'sql', 'extras')
        if not os.path.exists(sql_dir):
            os.makedirs(sql_dir)

        cross_reference_files = [f for f in os.listdir(json_dir) if f.startswith('cross_references') and f.endswith('.json')]
        for file in cross_reference_files:
            self.process_file(os.path.join(json_dir, file), sql_dir)

    def process_file(self, file_path, sql_dir):
        with open(file_path, 'r', encoding='utf-8') as jsonfile:
            data = json.load(jsonfile)

        sql_path = os.path.join(sql_dir, os.path.basename(file_path).replace('.json', '.sql'))
        with open(sql_path, 'w', encoding='utf-8') as sqlfile:
            sqlfile.write("-- SQL Dump for Cross References\n\n")

            # Create table
            sqlfile.write("""
            CREATE TABLE IF NOT EXISTS `cross_references` (
                `id` INT AUTO_INCREMENT PRIMARY KEY,
                `from_book` VARCHAR(255),
                `from_chapter` INT,
                `from_verse` INT,
                `to_book` VARCHAR(255),
                `to_chapter` INT,
                `to_verse_start` INT,
                `to_verse_end` INT,
                `votes` INT,
                INDEX (`from_book`),
                INDEX (`to_book`)
            );
            \n""")

            # Insert data
            for ref in data['cross_references']:
                from_verse = ref['from_verse']
                for to_verse in ref['to_verse']:
                    sqlfile.write(f"""
                    INSERT INTO `cross_references` (`from_book`, `from_chapter`, `from_verse`, `to_book`, `to_chapter`, `to_verse_start`, `to_verse_end`, `votes`)
                    VALUES ('{from_verse['book']}', {from_verse['chapter']}, {from_verse['verse']}, '{to_verse['book']}', {to_verse['chapter']}, {to_verse['verse_start']}, {to_verse['verse_end']}, {ref['votes']});
                    """)

        print(f"SQL dump for {file_path} generated at {sql_path}")
