# Bible Versions and Cross-Reference Databases: MySQL, CSV, JSON, YAML, TXT, MD.

This is a collection of bible versions in different formats. Here are some quick introductions:

- Find the bible versions you are looking for in the [Formats Folder](./formats).
- Find the source texts/resources in the [Sources Folder](./sources).
- Utilize the production / conversion scripts for this project in the [Scripts Folder](./scripts).

> **⚠️ Important: The legacy version of this project is available on the [2024](https://github.com/scrollmapper/bible_databases/tree/2024) branch. Please note that significant changes to the database schema have been implemented in the 2025 branch and subsequent versions.**


## Available Translations (29)

- [ACV](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/ACV)
- [ASV](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/ASV)
- [BBE](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/BBE)
- [BSB](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/BSB)
- [CPDV](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/CPDV)
- [DRC](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/DRC)
- [Darby](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Darby)
- [Geneva1599](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Geneva1599)
- [JPS](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/JPS)
- [Jubilee2000](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Jubilee2000)
- [KJVA](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/KJVA)
- [KJVPCE](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/KJVPCE)
- [KJV](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/KJV)
- [LEB](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/LEB)
- [LITV](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/LITV)
- [MKJV](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/MKJV)
- [NHEBJE](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/NHEBJE)
- [NHEBME](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/NHEBME)
- [NHEB](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/NHEB)
- [Noyes](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Noyes)
- [OEB](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/OEB)
- [RLT](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/RLT)
- [RWebster](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/RWebster)
- [Rotherham](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Rotherham)
- [Twenty](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Twenty)
- [Tyndale](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Tyndale)
- [UKJV](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/UKJV)
- [Webster](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/Webster)
- [YLT](https://github.com/scrollmapper/bible_databases/tree/2025/sources/en/YLT)

## Database Schema

The following sections describe the general database schema used in this project:

#### Table: `kjv_books`
This table lists all the books in the King James Version (KJV) of the Bible.

| Column Name | Type          | Nullable | Key         | Default | Extra          | Description                       |
|-------------|---------------|----------|-------------|---------|----------------|-----------------------------------|
| `id`        | int           | NO       | Primary Key | NULL    | auto_increment | Unique identifier for each book.  |
| `name`      | varchar(255)  | YES      |             | NULL    |                | The name of the book.             |

#### Table: `kjv_verses`
This table contains all the verses in the King James Version (KJV) of the Bible.

| Column Name | Type          | Nullable | Key         | Default | Extra          | Description                       |
|-------------|---------------|----------|-------------|---------|----------------|-----------------------------------|
| `id`        | int           | NO       | Primary Key | NULL    | auto_increment | Unique identifier for each verse. |
| `book_id`   | int           | YES      | Index       | NULL    |                | The ID of the book (foreign key to `kjv_books`). |
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

# MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
