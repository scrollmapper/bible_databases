import os
import json

class TextGenerator:
    def __init__(self, source_dir, format_dir):
        self.source_dir = source_dir
        self.format_dir = format_dir

    def generate(self, language, translation):
        data = self.load_json(language, translation)
        txt_path = os.path.join(self.format_dir, 'txt', f'{translation}.txt')

        with open(txt_path, 'w', encoding='utf-8') as txtfile:
            for book in data['books']:
                txtfile.write(f"### {book['name']}\n\n")
                for chapter in book['chapters']:
                    for verse in chapter['verses']:
                        txtfile.write(f"[{chapter['chapter']}:{verse['verse']}] {verse['text']}\n")
                txtfile.write("\n")

        print(f"TXT for {translation} generated at {txt_path}")

    def load_json(self, language, translation):
        json_path = os.path.join(self.source_dir, language, translation, f"{translation}.json")
        with open(json_path, 'r', encoding='utf-8') as file:
            return json.load(file)
