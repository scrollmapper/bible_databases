# Python Scripting / Text Management

Python is the primary scripting language used in this project. To ensure consistency and compatibility, we recommend using Python 3.9. This version has been tested thoroughly with our scripts and provides the necessary features and stability for text management and data processing.


### Scripts Breakdown

#### `build_mysql_database.py`

- **Description**: Builds a MySQL database from Bible translations and optionally includes cross-references. Prompts the user for database credentials and uses SQL dumps to populate the database.
- **Usage**: Run the script and follow the prompts to select translations and include cross-references.

#### `generate_all_versions.py`

- **Description**: Automates the generation of all Bible translations in multiple formats (SQL, CSV, JSON, TXT, YAML, MD). Iterates through all available translations and creates the corresponding files.
- **Usage**: Run the script to generate all formats for each translation.

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

#### `generate_txt.py`

- **Description**: Generates plain text (TXT) files for Bible translations. Each translation is processed and output as a text file.
- **Usage**: Run the script to create text files for each translation.

#### `generate_yaml.py`

- **Description**: Generates YAML files for Bible translations. Each translation is processed and output as a YAML file.
- **Usage**: Run the script to create YAML files for each translation.

#### `generate_cross_references.py`

- **Description**: Generates cross-reference data for Bible translations. Processes raw cross-reference data and formats it for use with Bible translations.
- **Usage**: Run the script to create cross-reference files.

#### `assemble_readme.py`

- **Description**: This script dynamically assembles the README.md file from various parts located in the `docs/main_readme` directory.
- **Purpose**: Ensures the main README.md is up-to-date with the latest information and structure.
- **Usage**: Run the script to generate the complete README.md.