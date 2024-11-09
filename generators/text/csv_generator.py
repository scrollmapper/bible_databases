import os
import json
import csv

class CSVGenerator:
    def __init__(self, source_dir, format_dir):
        self.source_dir = source_dir
        self.format_dir = format_dir

    def generate(self, language, translation):
        data = self.load_json(language, translation)
        csv_path = os.path.join(self.format_dir, 'csv', f'{translation}.csv')

        with open(csv_path, 'w', encoding='utf-8', newline='') as csvfile:
            writer = csv.writer(csvfile)
            writer.writerow(['Book', 'Chapter', 'Verse', 'Text'])  # Write header

            for book in data['books']:
                for chapter in book['chapters']:
                    for verse in chapter['verses']:
                        writer.writerow([book['name'], chapter['chapter'], verse['verse'], verse['text']])

        print(f"CSV for {translation} generated at {csv_path}")

    def load_json(self, language, translation):
        json_path = os.path.join(self.source_dir, language, translation, f"{translation}.json")
        with open(json_path, 'r', encoding='utf-8') as file:
            return json.load(file)
