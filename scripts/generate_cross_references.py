import sys
import os

# Add the parent directory to the system path
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from generators.text.cross_references_generator import CrossReferencesGenerator

def main():
    # Set base directories relative to the script location
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    source_directory = os.path.join(base_dir, 'sources')
    format_directory = os.path.join(base_dir, 'formats')

    # Generate Cross References JSON
    cross_references_generator = CrossReferencesGenerator(source_directory, format_directory)
    cross_references_generator.generate()

if __name__ == "__main__":
    main()
