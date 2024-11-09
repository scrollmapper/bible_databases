import os
import json

def read_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        return file.read()

def write_file(file_path, content):
    with open(file_path, 'w', encoding='utf-8') as file:
        file.write(content)

def get_translation_title(translation_dir):
    readme_path = os.path.join(translation_dir, 'README.md')
    if os.path.exists(readme_path):
        with open(readme_path, 'r', encoding='utf-8') as file:
            for line in file:
                if line.strip().startswith("#"):
                    return line.strip("# ").strip()
    return "Unknown Title"

def generate_translation_list(source_dir, output_file):
    translations = []

    # Loop through all languages
    for language in os.listdir(source_dir):
        language_dir = os.path.join(source_dir, language)
        if not os.path.isdir(language_dir):
            continue
        
        # Loop through all translations within the language
        for translation in os.listdir(language_dir):
            translation_dir = os.path.join(language_dir, translation)
            if os.path.isdir(translation_dir):
                title = get_translation_title(translation_dir)
                translations.append(f"- **{translation} ({language})**: {title}")
    
    translations.sort()
    translation_count = len(translations)

    translation_list_content = f"## Available Translations ({translation_count})\n\n" + "\n".join(translations)
    write_file(output_file, translation_list_content)

def assemble_readme():
    # Directory containing the parts
    docs_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', 'docs', 'main_readme'))
    source_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', 'sources'))

    # Paths to the files
    introduction_path = os.path.join(docs_dir, 'introduction.md')
    translation_list_path = os.path.join(docs_dir, 'translation_list.md')
    schema_path = os.path.join(docs_dir, 'schema.md')
    license_path = os.path.join(docs_dir, 'license.md')

    # Generate dynamic parts
    generate_translation_list(source_dir, translation_list_path)

    # Read the content of each part
    introduction = read_file(introduction_path)
    translation_list = read_file(translation_list_path)
    schema = read_file(schema_path)
    license_content = read_file(license_path)

    # Assemble the README content
    readme_content = f"{introduction}\n\n{translation_list}\n\n{schema}\n\n{license_content}"

    # Write the assembled README
    readme_path = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', 'README.md'))
    write_file(readme_path, readme_content)

    print(f"README.md has been assembled successfully!")

if __name__ == "__main__":
    assemble_readme()
