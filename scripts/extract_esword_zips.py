import os
import subprocess

def run_conversion(source_file, bible_version, output_file):
    command = f"python sword_to_json.py --source_file {source_file} --bible_version {bible_version} --output_file {output_file}"
    subprocess.run(command, shell=True)

def main():
    base_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', 'sources'))
    
    for language in os.listdir(base_dir):
        lang_path = os.path.join(base_dir, language)
        if not os.path.isdir(lang_path) or language == 'extras':
            continue
        
        for translation in os.listdir(lang_path):
            trans_path = os.path.join(lang_path, translation)
            if not os.path.isdir(trans_path):
                continue
            
            source_file = os.path.join(trans_path, f"{translation}.zip")
            output_file = os.path.join(trans_path, f"{translation}.json")
            
            if os.path.isfile(source_file) and not os.path.isfile(output_file):
                print(f"Converting {source_file} to {output_file}...")
                run_conversion(source_file, translation, output_file)
            else:
                print(f"Skipping {translation} as {output_file} already exists or {source_file} does not exist.")

if __name__ == '__main__':
    main()
