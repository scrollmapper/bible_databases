Bible Versions and Cross-Reference Databases: mysql, xml, csv, json.
===============

Here you will find the following bible versions in sql, xml, csv, and json format:

- American Standard-ASV1901 (ASV)
- Bible in Basic English (BBE)
- Darby English Bible (DARBY)
- King James Version (KJV)
- Webster's Bible (WBT)
- World English Bible (WEB)
- Young's Literal Translation (YLT)

Also find a **Master SQL file** (top level) containing all tables.


*Please let me know if any verses/text were inaccurately transferred in conversion.*

Each verse is accessed by a unique key, the combination of the BOOK+CHAPTER+VERSE id.

Example: 

**Genesis 1:1 (Genesis chapter 1, verse 1) = 01001001 (01 001 001)**

**Exodus 2:3 (Exodus chapter 2, verse 3) = 02002003 (02 002 003)** 


There is also a number-to-book key (**key_english** table), a cross-reference list (**cross_reference** table), and a bible key containing meta information about the included translations (**bible_version_key** table).

While its expected that your programs would use the verse-id system, *book #, chapter #, and verse #* columns have been included in the bible versions tables.

