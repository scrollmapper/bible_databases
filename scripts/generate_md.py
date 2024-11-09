import sys
import os

# Add the parent directory to the system path
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from generators.text.markdown_generator import MDGenerator

def list_options(options, prompt):
    for i, option in enumerate(options, 1):
        print(f"{i}. {option}")
    choice = int(input(prompt)) - 1
    return options[choice]

def main():
    # Set base directories relative to the script location
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_directory = os.path.join(base_dir, 'sources')
    format_directory = os.path.join(base_dir, 'formats')

    # Step 1: Select Language
    languages = [d for d in os.listdir(source_directory) if os.path.isdir(os.path.join(source_directory, d)) and d != "extras"]
    print("Choose your language:")
    language = list_options(languages, "Enter the number corresponding to your language: ")

    # Step 2: Select Translation
    translations = [d for d in os.listdir(os.path.join(source_directory, language)) if os.path.isdir(os.path.join(source_directory, language, d))]
    print(f"Choose your translation for {language}:")
    translation = list_options(translations, "Enter the number corresponding to your translation: ")

    # Step 3: Generate Markdown
    md_generator = MDGenerator(source_directory, format_directory)
    md_generator.generate(language, translation)

if __name__ == "__main__":
    main()
