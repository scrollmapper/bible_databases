## Database Schema

The following sections describe the general database schema used in this project:

#### Table: `<translation>_books`
This table lists all the books in the given translation of the Bible.

| Column Name | Type          | Nullable | Key         | Default | Extra          | Description                       |
|-------------|---------------|----------|-------------|---------|----------------|-----------------------------------|
| `id`        | int           | NO       | Primary Key | NULL    | auto_increment | Unique identifier for each book.  |
| `name`      | varchar(255)  | YES      |             | NULL    |                | The name of the book.             |

#### Table: `<translation>_verses`
This table contains all the verses in the given translation of the Bible.

| Column Name | Type          | Nullable | Key         | Default | Extra          | Description                       |
|-------------|---------------|----------|-------------|---------|----------------|-----------------------------------|
| `id`        | int           | NO       | Primary Key | NULL    | auto_increment | Unique identifier for each verse. |
| `book_id`   | int           | YES      | Index       | NULL    |                | The ID of the book (foreign key to `<translation>_books`). |
| `chapter`   | int           | YES      |             | NULL    |                | The chapter number.               |
| `verse`     | int           | YES      |             | NULL    |                | The verse number.                 |
| `text`      | text          | YES      |             | NULL    |                | The text of the verse.            |

#### Table: `translations`
This table contains information about the available Bible translations.

| Column Name  | Type          | Nullable | Key         | Default | Extra | Description                    |
|--------------|---------------|----------|-------------|---------|-------|--------------------------------|
| `translation`| varchar(255)  | NO       | Primary Key | NULL    |       | The abbreviation of the translation. |
| `title`      | varchar(255)  | YES      |             | NULL    |       | The full title of the translation. |
| `license`    | text          | YES      |             | NULL    |       | The license information for the translation. |

#### Table: `cross_references`
This table contains cross-reference data between different verses.

| Column Name     | Type          | Nullable | Key         | Default | Extra          | Description                           |
|-----------------|---------------|----------|-------------|---------|----------------|---------------------------------------|
| `id`            | int           | NO       | Primary Key | NULL    | auto_increment | Unique identifier for each cross-reference entry. |
| `from_book`     | varchar(255)  | YES      | Index       | NULL    |                | The book from which the cross-reference starts. |
| `from_chapter`  | int           | YES      |             | NULL    |                | The chapter number in the `from_book`. |
| `from_verse`    | int           | YES      |             | NULL    |                | The verse number in the `from_book`. |
| `to_book`       | varchar(255)  | YES      | Index       | NULL    |                | The book to which the cross-reference points. |
| `to_chapter`    | int           | YES      |             | NULL    |                | The chapter number in the `to_book`. |
| `to_verse_start`| int           | YES      |             | NULL    |                | The starting verse number in the `to_book`. |
| `to_verse_end`  | int           | YES      |             | NULL    |                | The ending verse number in the `to_book`. |
| `votes`         | int           | YES      |             | NULL    |                | The number of votes indicating the relevance of the cross-reference. |
