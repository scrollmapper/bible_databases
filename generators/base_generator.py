import json
import os

class BaseGenerator:
    def __init__(self, source_dir, format_dir):
        self.source_dir = source_dir
        self.format_dir = format_dir

    def load_json(self, language, translation):
        json_path = os.path.join(self.source_dir, language, translation, f"{translation}.json")
        with open(json_path, 'r', encoding='utf-8') as file:
            return json.load(file)
    
    def get_readme_title(self, language, translation):
        readme_path = os.path.join(self.source_dir, language, translation, "README.md")
        with open(readme_path, 'r', encoding='utf-8') as file:
            for line in file:
                if line.startswith("# "):
                    return line[2:].strip()
    
    def prepare_data(self, data):
        # Common data preparation logic (if needed)
        return data
