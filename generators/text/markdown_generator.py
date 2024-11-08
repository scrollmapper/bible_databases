import os
from generators.base_generator import BaseGenerator

class MDGenerator(BaseGenerator):
    def __init__(self, source_dir, format_dir):
        super().__init__(source_dir, format_dir)

    def generate(self, language, translation):
        data = self.load_json(language, translation)
        translation_name = self.get_readme_title(language, translation)
        prepared_data = self.prepare_data(data)
        md_path = os.path.join(self.format_dir, 'md', f'{translation}.md')

        with open(md_path, 'w', encoding='utf-8') as mdfile:
            mdfile.write(f"# {translation_name}\n\n")
            for book in prepared_data['books']:
                mdfile.write(f"## {book['name']}\n\n")
                for chapter in book['chapters']:
                    mdfile.write(f"### Chapter {chapter['chapter']}\n\n")
                    for verse in chapter['verses']:
                        mdfile.write(f"**[{chapter['chapter']}:{verse['verse']}]** {verse['text']}\n\n")

        print(f"Markdown file for {translation_name} ({translation}) generated at {md_path}")
