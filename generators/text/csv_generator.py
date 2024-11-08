import csv
import os
from generators.base_generator import BaseGenerator

class CSVGenerator(BaseGenerator):
    def __init__(self, source_dir, format_dir):
        super().__init__(source_dir, format_dir)

    def generate(self, language, translation):
        data = self.load_json(language, translation)
        translation_name = self.get_readme_title(language, translation)
        prepared_data = self.prepare_data(data)
        csv_path = os.path.join(self.format_dir, 'csv', f'{translation}.csv')

        with open(csv_path, 'w', newline='') as csvfile:
            writer = csv.writer(csvfile)
            writer.writerow(['Book', 'Chapter', 'Verse', 'Text'])  # Header
            for book in prepared_data['books']:
                for chapter in book['chapters']:
                    for verse in chapter['verses']:
                        writer.writerow([book['name'], chapter['chapter'], verse['verse'], verse['text']])

        print(f"CSV file for {translation_name} ({translation}) generated at {csv_path}")
