import sys
import os
import subprocess

# Add the parent directory to the system path
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from generators.sql.mysql_generator import MySQLGenerator
from generators.sql.cross_references_generator_mysql import CrossReferencesGeneratorMySQL

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
  
    # Step 3: Generate MySQL Dump for Translation
    mysql_generator = MySQLGenerator(source_directory, format_directory)
    mysql_generator.generate(language, translation)

    # Step 4: Optional - Include Cross References
    print("The cross-reference table is quite large. It will take a while to import.")
    include_cross_references = input("Do you want to include cross references? (yes/no): ").strip().lower() == 'yes'
    if include_cross_references:
        cross_references_generator = CrossReferencesGeneratorMySQL(source_directory, format_directory)
        cross_references_generator.generate()

    # Step 5: Build the Database
    db_name = input("Enter the database name: ").strip()
    db_user = input("Enter the MySQL user: ").strip()
    db_password = os.getenv("MYSQL_PWD")

    if not db_password:
        db_password = input("Enter the MySQL password: ").strip()

    # SQL commands to disable and enable foreign key checks
    disable_fk_checks = "SET FOREIGN_KEY_CHECKS=0;"
    enable_fk_checks = "SET FOREIGN_KEY_CHECKS=1;"

    # Run the MySQL import commands
    print("Building the database...")
    sql_files = [f for f in os.listdir(os.path.join(format_directory, 'sql')) if f == "%s.sql"%translation]

    # Disable foreign key checks before importing
    subprocess.run(f"mysql -u {db_user} -p{db_password} {db_name} -e \"{disable_fk_checks}\"", shell=True, check=True)
    
    for sql_file in sql_files:
        if 'extras' in sql_file:
            continue  # Skip cross-reference SQL files for now

        sql_file_path = os.path.join(format_directory, 'sql', sql_file)
        cmd = f"mysql -u {db_user} -p{db_password} {db_name} < {sql_file_path}"
        subprocess.run(cmd, shell=True, check=True)
        print(f"Imported {sql_file}")

    # Process cross-reference SQL files if included
    if include_cross_references:
        extras_dir = os.path.join(format_directory, 'sql', 'extras')
        cross_ref_files = [f for f in os.listdir(extras_dir) if f.endswith('.sql')]
        
        for cross_ref_file in cross_ref_files:
            cross_ref_file_path = os.path.join(extras_dir, cross_ref_file)
            cmd = f"mysql -u {db_user} -p{db_password} {db_name} < {cross_ref_file_path}"
            subprocess.run(cmd, shell=True, check=True)
            print(f"Imported {cross_ref_file}")

    # Enable foreign key checks after importing
    subprocess.run(f"mysql -u {db_user} -p{db_password} {db_name} -e \"{enable_fk_checks}\"", shell=True, check=True)
    
    print("Database build complete!")

if __name__ == "__main__":
    main()
