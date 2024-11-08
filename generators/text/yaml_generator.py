import os
import yaml
from generators.base_generator import BaseGenerator

class YAMLGenerator(BaseGenerator):
    def __init__(self, source_dir, format_dir):
        super().__init__(source_dir, format_dir)

    def generate(self, language, translation):
        data = self.load_json(language, translation)
        translation_name = self.get_readme_title(language, translation)
        prepared_data = self.prepare_data(data)
        yaml_path = os.path.join(self.format_dir, 'yaml', f'{translation}.yaml')

        # Create structured YAML data
        structured_data = {
            "translation": translation_name,
            "books": []
        }

        for book in prepared_data['books']:
            book_data = {
                "name": book['name'],
                "chapters": []
            }
            for chapter in book['chapters']:
                chapter_data = {
                    "chapter": chapter['chapter'],
                    "verses": []
                }
                for verse in chapter['verses']:
                    verse_data = {
                        "verse": verse['verse'],
                        "text": verse['text']
                    }
                    chapter_data["verses"].append(verse_data)
                book_data["chapters"].append(chapter_data)
            structured_data["books"].append(book_data)

        # Write YAML data to file
        with open(yaml_path, 'w', encoding='utf-8') as yamlfile:
            yaml.dump(structured_data, yamlfile, allow_unicode=True, default_flow_style=False)

        print(f"YAML file for {translation_name} ({translation}) generated at {yaml_path}")
