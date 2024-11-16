import os
import sys

# Add the parent directory to the system path
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from generators.sql.mysql_generator import MySQLGenerator
from generators.sql.cross_references_generator_mysql import CrossReferencesGeneratorMySQL
from generators.sqlite.sqlite_generator import SQLiteGenerator
from generators.json.json_generator import JSONGenerator
from generators.text.csv_generator import CSVGenerator
from generators.text.plaintext_generator import TextGenerator
from generators.text.yaml_generator import YAMLGenerator
from generators.text.markdown_generator import MDGenerator

def create_format_directories(format_directory):
    formats = ['sql', 'sqlite', 'csv', 'txt', 'json', 'yaml', 'md']
    for fmt in formats:
        dir_path = os.path.join(format_directory, fmt)
        if not os.path.exists(dir_path):
            os.makedirs(dir_path)

def generate_all_versions():
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_directory = os.path.join(base_dir, 'sources')
    format_directory = os.path.join(base_dir, 'formats')

    # Create necessary directories for formats
    create_format_directories(format_directory)

    # List all languages
    languages = [d for d in os.listdir(source_directory) if os.path.isdir(os.path.join(source_directory, d))]

    # Generate all formats for each language and translation
    for language in languages:
        language_path = os.path.join(source_directory, language)
        
        # List all translations for the current language
        translations = [d for d in os.listdir(language_path) if os.path.isdir(os.path.join(language_path, d))]

        for translation in translations:
            try:
                print(f"Generating formats for {translation} in {language}...")

                # Generate MySQL dump for the translation
                mysql_generator = MySQLGenerator(source_directory, format_directory)
                mysql_generator.generate(language, translation)

                # Generate SQLite database for the translation
                sqlite_generator = SQLiteGenerator(source_directory, format_directory)
                sqlite_generator.generate(language, translation)

                # Generate CSV format
                csv_generator = CSVGenerator(source_directory, format_directory)
                csv_generator.generate(language, translation)

                # Generate TXT format
                txt_generator = TextGenerator(source_directory, format_directory)
                txt_generator.generate(language, translation)

                # Generate JSON format
                json_generator = JSONGenerator(source_directory, format_directory)
                json_generator.generate(language, translation)

                # Generate YAML format
                yaml_generator = YAMLGenerator(source_directory, format_directory)
                yaml_generator.generate(language, translation)

                # Generate MD format
                md_generator = MDGenerator(source_directory, format_directory)
                md_generator.generate(language, translation)

                print(f"Completed generating formats for {translation} in {language}")
            except Exception as e:
                print(f"Error generating formats for {translation} in {language}: {e}")

if __name__ == "__main__":
    generate_all_versions()
