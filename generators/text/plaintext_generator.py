import os
from generators.base_generator import BaseGenerator

class TextGenerator(BaseGenerator):
    def __init__(self, source_dir, format_dir):
        super().__init__(source_dir, format_dir)

    def generate(self, language, translation):
        data = self.load_json(language, translation)
        translation_name = self.get_readme_title(language, translation)
        prepared_data = self.prepare_data(data)
        txt_path = os.path.join(self.format_dir, 'txt', f'{translation}.txt')

        with open(txt_path, 'w') as txtfile:
            # Write translation title
            txtfile.write(f"{translation_name}\n\n")
            for book in prepared_data['books']:
                # Write book title
                txtfile.write(f"\nBook: {book['name']}\n")
                for chapter in book['chapters']:
                    # Write chapter title
                    txtfile.write(f"\n\nChapter {chapter['chapter']}\n\n")
                    for verse in chapter['verses']:
                        # Write verse
                        txtfile.write(f"[{chapter['chapter']}:{verse['verse']}] {verse['text']}\n")

        print(f"Text file for {translation_name} ({translation}) generated at {txt_path}")
