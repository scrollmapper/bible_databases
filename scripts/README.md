# Bible Databases Scripts

This repository contains a collection of scripts designed to generate various formats of Bible translations and cross-reference data. Below is a description of each script and its function.

### Scripts

- **extract_esword_zips.py**
  - **Description**: Extracts all ESword zip files in the sources directory that do not yet have an accompanying `<translation>.json` file.
  - **Usage**: Run the script ... it knows what to do.


- **build_mysql_database.py**
  - **Description**: Builds a MySQL database from Bible translations and optionally includes cross-references. Prompts the user for database credentials and uses SQL dumps to populate the database. You can use this for every version you wish to add. Cross-reference tables only need
  to be added once.
  - **Usage**: Run the script and follow the prompts to select translations and include cross-references.

- **export_sqlite_database.py**
  - **Description**: Creates an SQLite database for a selected Bible translation and includes cross references. Prompts the user for the path where the new database should be built.
  - **Usage**: Run the script and follow the prompts to create the SQLite database with cross references.

- **generate_all_versions.py**
  - **Description**: Automates the generation of all Bible translations in multiple formats (SQL, SQLite, CSV, JSON, TXT, YAML, MD). Iterates through all available translations and creates the corresponding files.
  - **Usage**: Run the script to generate all formats for each translation.

- **generate_cross_references.py**
  - **Description**: Generates cross-reference data for Bible translations. Processes raw cross-reference data and formats it for use with Bible translations.
  - **Usage**: Run the script to create cross-reference files.

- **generate_cross_references_mysql.py**
  - **Description**: Specifically generates cross-reference data for MySQL databases. Formats the cross-reference data into SQL insert statements.
  - **Usage**: Run the script to create cross-reference SQL files.

- **generate_csv.py**
  - **Description**: Generates CSV files for Bible translations. Each translation is processed and output as a CSV file.
  - **Usage**: Run the script to create CSV files for each translation.

- **generate_json.py**
  - **Description**: Generates JSON files for Bible translations. Each translation is processed and output as a JSON file.
  - **Usage**: Run the script to create JSON files for each translation.

- **generate_md.py**
  - **Description**: Generates Markdown (MD) files for Bible translations. Each translation is processed and output as a Markdown file.
  - **Usage**: Run the script to create Markdown files for each translation.

- **generate_mysql.py**
  - **Description**: Generates MySQL SQL dump files for Bible translations. Each translation is processed and output as an SQL dump file.
  - **Usage**: Run the script to create SQL dump files for each translation.

- **generate_sqlite.py**
  - **Description**: Generates SQLite database files for Bible translations. Each translation is processed and output as an SQLite database file.
  - **Usage**: Run the script to create SQLite database files for each translation.

- **generate_txt.py**
  - **Description**: Generates plain text (TXT) files for Bible translations. Each translation is processed and output as a text file.
  - **Usage**: Run the script to create text files for each translation.

- **generate_yaml.py**
  - **Description**: Generates YAML files for Bible translations. Each translation is processed and output as a YAML file.
  - **Usage**: Run the script to create YAML files for each translation.

#### `verify_text_integrity_<format>.py`
- **Description**: Checks the integrity of the reformatted text against the source .json files in sources directory. It will output the verification in this directory. Relocate it or delete it after check.
- **Usage**: Run the script and follow the prompts.

### How to Run the Scripts

1. **Setup**:
   - Ensure you have Python **3.8** or higher installed on your machine.
   - Clone the repository to your local machine.
   - Navigate to the directory containing the scripts.

2. **Run a Script**:
   - Open a terminal or command prompt.
   - Navigate to the script directory.
   - Run the script by typing `python <script_name>.py` and follow any prompts if applicable.

Example:
```bash
python generate_all_versions.py
