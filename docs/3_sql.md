# SQL / Database Schema

## MySQL

MySQL remains the primary focus and schema for all SQL implementations in this repository. While other SQL types (e.g., SQLite, PostgreSQL) can be added, MySQL sets the standard.

### Schema Adjustments

In the initial version of this project [(refer to the 2014 - 2024 README)](https://github.com/scrollmapper/bible_databases/blob/2024/README.md), we utilized a verse-ID system that facilitated fast and straightforward queries. However, with the integration of a larger and more diverse library of Bible translations—including versions that contain the Apocrypha—we have revised the database schema. This update allows each Bible version to include its own set of books, independent of the traditional 66-book canon.

## Introduction and Query Examples

### Rationale Behind the Database Schema

The database schema has been carefully designed to accommodate a diverse range of Bible translations, each with unique characteristics and requirements. Here’s a brief explanation of the schema and the reasons behind its structure:

1. **Flexibility Across Translations**:
   - Each Bible translation can have its own set of books, including those that may contain the Apocrypha or other additional texts. This flexibility ensures that the schema can handle any translation, not just those confined to the standard 66-book canon.

2. **Table Structure**:
   - **Books Table (`<translation>_books`)**: This table stores information about the books in each translation. Each book has a unique identifier (`id`) and a name (`name`).
     - **Purpose**: By separating books into their own table, we can manage and reference them efficiently across different translations.
   
   - **Verses Table (`<translation>_verses`)**: This table contains the actual verses from the Bible. Each verse is linked to a specific book through `book_id`, and includes details like chapter, verse number, and the text of the verse.
     - **Purpose**: Organizing verses in a dedicated table allows for efficient querying and retrieval of specific passages.
   
   - **Translations Table (`translations`)**: This table provides metadata about each translation, including its abbreviation, full title, and license information.
     - **Purpose**: Storing translation metadata separately ensures that all translations are consistently documented and easily accessible.
   
   - **Cross References Table (`cross_references`)**: This table holds data about cross-references between verses. Each entry identifies relationships between verses in different books and chapters.
     - **Purpose**: Cross references enhance the study of scripture by highlighting connections between different passages, making it easier to explore related themes and concepts.

### Query Examples

Let's dive into some practical examples of how to query this database:

#### 1. Retrieve All Books in a Translation

```sql
SELECT * FROM kjv_books;
```

Replace kjv_books with <translation>_books for other translations.

#### 2. Retrieve the Passage "Turn You at My Reproof" in Proverbs

```sql
SELECT * FROM kjv_verses
WHERE book_id = (SELECT id FROM kjv_books WHERE name = 'Proverbs')
AND text LIKE '%turn you at my reproof%';
```

#### 3. Get the First Chapter of John

```sql
SELECT * FROM kjv_verses
WHERE book_id = (SELECT id FROM kjv_books WHERE name = 'John')
AND chapter = 1;
```

#### 4. Find Cross References for a Specific Verse in Proverbs

```sql
SELECT * FROM cross_references
WHERE from_book = 'Proverbs'
AND from_chapter = 1
AND from_verse = 23;
```

#### 5. Retrieve Cross-Referenced Scriptures for Proverbs 1:23

```sql
SELECT 
    cr.from_book, cr.from_chapter, cr.from_verse, cr.to_book, cr.to_chapter, cr.to_verse_start, cr.to_verse_end,
    v.text
FROM 
    cross_references cr
JOIN 
    kjv_verses v ON cr.to_book = (SELECT name FROM kjv_books WHERE id = v.book_id)
                AND cr.to_chapter = v.chapter
                AND v.verse BETWEEN cr.to_verse_start AND cr.to_verse_end
WHERE 
    cr.from_book = 'Proverbs'
    AND cr.from_chapter = 1
    AND cr.from_verse = 23;
```

That example query should return this result:

#### 9. Extracted Verses for Proverbs 1:23 Cross References

```sql
+-----------+--------------+------------+-----------+------------+----------------+--------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Proverbs  |            1 |         23 | Psalms    |        145 |              1 |            1 | I will extol thee, my God, O king; and I will bless thy name for ever and ever.
|
| Proverbs  |            1 |         23 | Proverbs  |          1 |             25 |           25 | But ye have set at nought all my counsel, and would none of my reproof:
|
| Proverbs  |            1 |         23 | Proverbs  |          1 |             30 |           30 | They would none of my counsel: they despised all my reproof.
|
| Proverbs  |            1 |         23 | Proverbs  |          6 |             23 |           23 | For the commandment is a lamp; and the law is light; and reproofs of instruction are the way of life:
|
| Proverbs  |            1 |         23 | Proverbs  |         10 |             17 |           17 | He is in the way of life that keepeth instruction: but he that refuseth reproof erreth.
|
| Proverbs  |            1 |         23 | Proverbs  |         12 |              1 |            1 | Whoso loveth instruction loveth knowledge: but he that hateth reproof is brutish.
|
| Proverbs  |            1 |         23 | Proverbs  |         29 |              1 |            1 | He, that being often reproved hardeneth his neck, shall suddenly be destroyed, and that without remedy.
|
| Proverbs  |            1 |         23 | Isaiah    |         32 |             15 |           15 | Until the spirit be poured upon us from on high, and the wilderness be a fruitful field, and the fruitful field be counted for a forest.     
|
| Proverbs  |            1 |         23 | Isaiah    |         45 |              8 |            8 | Drop down, ye heavens, from above, and let the skies pour down righteousness: let the earth open, and let them bring forth salvation, and let righteousness spring up together; I the Lord have created it.
|
| Proverbs  |            1 |         23 | Isaiah    |         55 |              1 |            1 | Ho, every one that thirsteth, come ye to the waters, and he that hath no money; come ye, buy, and eat; yea, come, buy wine and milk without money and without price.
|
| Proverbs  |            1 |         23 | Isaiah    |         55 |              3 |            3 | Incline your ear, and come unto me: hear, and your soul shall live; and I will make an everlasting covenant with you, even the sure mercies of David.
|
| Proverbs  |            1 |         23 | Isaiah    |         55 |              6 |            6 | Seek ye the Lord while he may be found, call ye upon him while he is near:
|
| Proverbs  |            1 |         23 | Isaiah    |         55 |              7 |            7 | Let the wicked forsake his way, and the unrighteous man his thoughts: and let him return unto the Lord, and he will have mercy upon him; and to our God, for he will abundantly pardon.
|
| Proverbs  |            1 |         23 | Jeremiah  |          3 |             14 |           14 | Turn, O backsliding children, saith the Lord; for I am married unto you: and I will take you one of a city, and two of a family, and I will bring you to Zion:
|
| Proverbs  |            1 |         23 | Ezekiel   |         18 |             27 |           27 | Again, when the wicked man turneth away from his wickedness that he hath committed, and doeth that which is lawful and right, he shall save his soul alive.
|
| Proverbs  |            1 |         23 | Ezekiel   |         18 |             30 |           30 | Therefore I will judge you, O house of Israel, every one according to his ways, saith the Lord God. Repent, and turn yourselves from all your transgressions; so iniquity shall not be your ruin.
|
| Proverbs  |            1 |         23 | Ezekiel   |         33 |             11 |           11 | Say unto them, As I live, saith the Lord God, I have no pleasure in the death of the wicked; but that the wicked turn from his way and live: turn ye, turn ye from your evil ways; for why will ye die, O house of Israel?
|
| Proverbs  |            1 |         23 | Hosea     |         14 |              1 |            1 | O Israel, return unto the Lord thy God; for thou hast fallen by thine iniquity.
|
| Proverbs  |            1 |         23 | Joel      |          2 |             28 |           28 | And it shall come to pass afterward, that I will pour out my spirit upon all flesh; and your sons and your daughters shall prophesy, your old men shall dream dreams, your young men shall see visions:
|
| Proverbs  |            1 |         23 | Zechariah |         12 |             10 |           10 | And I will pour upon the house of David, and upon the inhabitants of Jerusalem, the spirit of grace and of supplications: and they shall look upon me whom they have pierced, and they shall mourn for him, as one mourneth for his only son, and shall be in bitterness for him, as one that is in bitterness for his firstborn.
|
| Proverbs  |            1 |         23 | Luke      |         11 |             13 |           13 | If ye then, being evil, know how to give good gifts unto your children: how much more shall your heavenly Father give the Holy Spirit to them that ask him?
|
| Proverbs  |            1 |         23 | John      |          7 |             36 |           36 | What manner of saying is this that he said, Ye shall seek me, and shall not find me: and where I am, thither ye cannot come?
|
| Proverbs  |            1 |         23 | John      |          7 |             37 |           37 | In the last day, that great day of the feast, Jesus stood and cried, saying, If any man thirst, let him come unto me, and drink.
|
| Proverbs  |            1 |         23 | Acts      |          2 |             36 |           36 | Therefore let all the house of Israel know assuredly, that God hath made that same Jesus, whom ye have crucified, both Lord and Christ.
|
| Proverbs  |            1 |         23 | Acts      |          2 |             38 |           38 | Then Peter said unto them, Repent, and be baptized every one of you in the name of Jesus Christ for the remission of sins, and ye shall receive the gift of the Holy Ghost.
|
| Proverbs  |            1 |         23 | Acts      |          3 |             19 |           19 | Repent ye therefore, and be converted, that your sins may be blotted out, when the times of refreshing shall come from the presence of the Lord;
|
| Proverbs  |            1 |         23 | Acts      |         26 |             20 |           20 | But shewed first unto them of Damascus, and at Jerusalem, and throughout all the coasts of Judea, and then to the Gentiles, that they should repent and turn to God, and do works meet for repentance.
|
+-----------+--------------+------------+-----------+------------+----------------+--------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
```