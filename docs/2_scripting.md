# Python Scripting / Text Management

Python is the primary scripting language used in this project. To ensure consistency and compatibility, we recommend using Python 3.9. This version has been tested thoroughly with our scripts and provides the necessary features and stability for text management and data processing.

### Dependencies:

To use the scripts in the scripts directory, be sure you have the dependencies installed:

```
pip install mysql.connector
pip install past.builtins
pip install pysword.modules
pip install yaml
```

### Scripts Breakdown

#### `assemble_readme.py`
- **Description**: Dynamically assembles the README.md file from various parts located in the `docs/main_readme` directory.
- **Usage**: Run the script to generate the complete README.md.

- **export_sqlite_database.py**
  - **Description**: Creates an SQLite database for a selected Bible translation and includes cross references. Prompts the user for the path where the new database should be built.
  - **Usage**: Run the script and follow the prompts to create the SQLite database with cross references.

#### `build_sqlite_database.py`
- **Description**: Creates an SQLite database for a selected Bible translation and includes cross references. Prompts the user for the path where the new database should be built.
- **Usage**: Run the script and follow the prompts to create the SQLite database with cross references.

#### `generate_all_versions.py`
- **Description**: Automates the generation of all Bible translations in multiple formats (SQL, SQLite, CSV, JSON, TXT, YAML, MD). Iterates through all available translations and creates the corresponding files.
- **Usage**: Run the script to generate all formats for each translation.

#### `generate_cross_references.py`
- **Description**: Generates cross-reference data for Bible translations. Processes raw cross-reference data and formats it for use with Bible translations.
- **Usage**: Run the script to create cross-reference files.

#### `generate_cross_references_mysql.py`
- **Description**: Specifically generates cross-reference data for MySQL databases. Formats the cross-reference data into SQL insert statements.
- **Usage**: Run the script to create cross-reference SQL files.

#### `generate_csv.py`
- **Description**: Generates CSV files for Bible translations. Each translation is processed and output as a CSV file.
- **Usage**: Run the script to create CSV files for each translation.

#### `generate_json.py`
- **Description**: Generates JSON files for Bible translations. Each translation is processed and output as a JSON file.
- **Usage**: Run the script to create JSON files for each translation.

#### `generate_md.py`
- **Description**: Generates Markdown (MD) files for Bible translations. Each translation is processed and output as a Markdown file.
- **Usage**: Run the script to create Markdown files for each translation.

#### `generate_mysql.py`
- **Description**: Generates MySQL SQL dump files for Bible translations. Each translation is processed and output as an SQL dump file.
- **Usage**: Run the script to create SQL dump files for each translation.

#### `generate_sqlite.py`
- **Description**: Generates SQLite database files for Bible translations. Each translation is processed and output as an SQLite database file.
- **Usage**: Run the script to create SQLite database files for each translation.

#### `generate_txt.py`
- **Description**: Generates plain text (TXT) files for Bible translations. Each translation is processed and output as a text file.
- **Usage**: Run the script to create text files for each translation.

#### `generate_yaml.py`
- **Description**: Generates YAML files for Bible translations. Each translation is processed and output as a YAML file.
- **Usage**: Run the script to create YAML files for each translation.

#### `export_sqlite_database.py`
- **Description**: Exports data from the SQLite database to a specified format.
- **Usage**: Run the script to export data from the SQLite database.

#### `verify_text_integrity_<format>.py`
- **Description**: Checks the integrity of the reformatted text against the source .json files in sources directory.
- **Usage**: Run the script and follow the prompts.

### How to Run the Scripts

1. **Setup**:
   - Ensure you have Python **3.9** or higher installed on your machine.
   - Clone the repository to your local machine.
   - Navigate to the directory containing the scripts.

2. **Run a Script**:
   - Open a terminal or command prompt.
   - Navigate to the script directory.
   - Run the script by typing `python <script_name>.py` and follow any prompts if applicable.

Example:
```bash
python generate_all_versions.py
