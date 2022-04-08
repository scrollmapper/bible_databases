Bible Versions and Cross-Reference Databases: mysql, sqlite, xml, csv, json, txt, md.
===============

https://scrollmapper.github.io/

Here you will find the following bible versions in sql, sqlite, xml, csv, json, txt, and md format:

- **King James Version (KJV)**
- American Standard-ASV1901 (ASV)
- Bible in Basic English (BBE)
- World English Bible (WEB)
- Young's Literal Translation (YLT)

All included Bible translations are in the public domain.

*Please let me know if any verses/text were inaccurately transferred in conversion.*

Looking for deuterocanonical / apocryphal books? Find them here: https://github.com/scrollmapper/bible_databases_deuterocanonical

Summary:
-------------------

- *bible-mysql.sql* (MySQL) is the main database and most feature-oriented due to contributions from developers. It is suggested you use that for most things, or at least convert the information from it.
- *cross_references-mysql.sql* (MySQL) is the cross-reference table. It has been separated to become an optional feature. This is converted from the project at http://www.openbible.info/labs/cross-references/.
- *bible-sqlite.db* (SQLite) is converted from bible-mysql.sql using mysql2sqlite. (includes cross-references too).
- *cross_references.txt* is the source cross-reference file obtained from http://www.openbible.info/labs/cross-references/


Verse ID System:
-------------------

Each verse is accessed by a unique key, the combination of the BOOK+CHAPTER+VERSE id.

Example:

**Genesis 1:1 (Genesis chapter 1, verse 1) = 01001001 (01 001 001)**

**Exodus 2:3 (Exodus chapter 2, verse 3) = 02002003 (02 002 003)**

The verse-id system is used for faster, simplified queries. For instance:

***01001001*** - ***02001005*** would capture all verses between **Genesis 1:1** through **Exodus 1:5**.

Written simply:

<code>
SELECT *
FROM bible.t_asv
WHERE id
BETWEEN 01001001
AND 02001005
</code>

Coordinating Tables
-------------------

There is also a number-to-book key (**key_english** table), a cross-reference list (**cross_reference** table), and a bible key containing meta information about the included translations (**bible_version_key** table). ***See below SQL table layout.*** These tables work together providing you a great basis for a bible-reading and cross-referencing app. In addition, each book is marked with a particular genre, mapping in the number-to-genre key (**key_genre_english** table) and common abbreviations for each book can be looked up in the abbreviations list (**key_abbreviations_english** table).

While its expected that your programs would use the verse-id system, *book #, chapter #, and verse #* columns have been included in the bible versions tables.

A Valuable Cross-Reference Table
-------------------
A very special and valuable addition to these databases is the extensive cross-reference table. It was created from the project at http://www.openbible.info/labs/cross-references/. See .txt version included from http://www.openbible.info website. Its extremely useful in bible study for discovering related scriptures. For any given verse, you simply query ***vid*** (verse id), and a list of rows will be returned. Each of those rows has a rank (r) for relevance, start-verse (sv), and end verse (ev) if there is one.

Basic Web Interaction
-------------------
The web folder contains two php files. Edit the first few lines of index.php to match your server's settings. Place these in a folder on your webserver.

The references search box can be multiple comma separated values. (i.e. John 3:16, Rom 3:23, 1 Jn 1:9, Romans 10:9-10) You can also directly link to a verse by altering the URI: [http://localhost/index.php?b=John 3:16, Rom 3:23, 1 Jn 1:9, Romans 10:9-10](http://localhost/index.php?b=John 3:16, Rom 3:23, 1 Jn 1:9, Romans 10:9-10)

-----------------------------

SQL Database Layout:
-------------------

<h3>bible_version_key</h3>
<table cellpadding="0" cellspacing="0" class="db-table">
   <tr>
      <th>Field</th>
      <th>Type</th>
      <th>Null</th>
      <th>Key</th>
      <th>Default</th>
      <th>Extra</th>
   </tr>
   <tr>
      <td>id</td>
      <td>int(3) unsigned zerofill</td>
      <td>NO</td>
      <td>PRI</td>
      <td></td>
      <td>auto_increment</td>
   </tr>
   <tr>
      <td>table</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Corresponding Bible version table name.</td>
   </tr>
   <tr>
      <td>abbreviation</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Bible version abbreviation.</td>
   </tr>
   <tr>
      <td>language</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Language translation.</td>
   </tr>
   <tr>
      <td>version</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Version name.</td>
   </tr>
   <tr>
      <td>info_text</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Extra info.</td>
   </tr>
   <tr>
      <td>info_url</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Wiki info or similar URL</td>
   </tr>
   <tr>
      <td>publisher</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Who published.</td>
   </tr>
   <tr>
      <td>copyright</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Copyright</td>
   </tr>
   <tr>
      <td>copyright_info</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>More info on copyright.</td>
   </tr>
</table>
<br />
<h3>cross_reference</h3>
<table cellpadding="0" cellspacing="0" class="db-table">
   <tr>
      <th>Field</th>
      <th>Type</th>
      <th>Null</th>
      <th>Key</th>
      <th>Default</th>
      <th>Extra</th>
   </tr>
   <tr>
      <td>vid</td>
      <td>int(8) unsigned zerofill</td>
      <td>NO</td>
      <td>MUL</td>
      <td></td>
      <td>VERSE ID</td>
   </tr>
   <tr>
      <td>r</td>
      <td>int(11)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>RANK</td>
   </tr>
   <tr>
      <td>sv</td>
      <td>int(8) unsigned zerofill</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>START VERSE</td>
   </tr>
   <tr>
      <td>ev</td>
      <td>int(8) unsigned zerofill</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>END VERSE</td>
   </tr>
</table>
<br />
<h3>key_english</h3>
<table cellpadding="0" cellspacing="0" class="db-table">
   <tr>
      <th>Field</th>
      <th>Type</th>
      <th>Null</th>
      <th>Key</th>
      <th>Default</th>
      <th>Extra</th>
   </tr>
   <tr>
      <td>b</td>
      <td>int(11)</td>
      <td>NO</td>
      <td>PRI</td>
      <td></td>
      <td>Book #</td>
   </tr>
   <tr>
      <td>n</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Book Name</td>
   </tr>
   <tr>
      <td>t</td>
      <td>VARCHAR(2)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Testament (OT or NT)</td>
   </tr>
   <tr>
      <td>g</td>
      <td>tinyint(3)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Genre ID</td>
   </tr>
</table>
<br />
<h3>key_abbreviations_english</h3>
<table cellpadding="0" cellspacing="0" class="db-table">
   <tr>
      <th>Field</th>
      <th>Type</th>
      <th>Null</th>
      <th>Key</th>
      <th>Default</th>
      <th>Extra</th>
   </tr>
   <tr>
      <td>id</td>
      <td>smallint(5)</td>
      <td>NO</td>
      <td>PRI</td>
      <td></td>
      <td>Abbreviation ID</td>
   </tr>
   <tr>
      <td>a</td>
      <td>VARCHAR(255)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Abbreviation</td>
   </tr>
   <tr>
      <td>b</td>
      <td>smallint(5)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Book ID this is refers to</td>
   </tr>
   <tr>
      <td>p</td>
      <td>tinyint(1)</td>
      <td>NO</td>
      <td></td>
      <td>0</td>
      <td>If this is the desired abbreviation</td>
   </tr>
</table>
<br />
<h3>key_genre_english</h3>
<table cellpadding="0" cellspacing="0" class="db-table">
   <tr>
      <th>Field</th>
      <th>Type</th>
      <th>Null</th>
      <th>Key</th>
      <th>Default</th>
      <th>Extra</th>
   </tr>
   <tr>
      <td>g</td>
      <td>tinyint(3)</td>
      <td>NO</td>
      <td>PRI</td>
      <td></td>
      <td>Genre ID</td>
   </tr>
   <tr>
      <td>n</td>
      <td>VARCHAR(255)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Genre name</td>
   </tr>
</table>
<br />
<h3>t_(version abbreviation)</h3>
<table cellpadding="0" cellspacing="0" class="db-table">
   <tr>
      <th>Field</th>
      <th>Type</th>
      <th>Null</th>
      <th>Key</th>
      <th>Default
      <th>Extra</th>
   </tr>
   <tr>
      <td>id</td>
      <td>int(8) unsigned zerofill</td>
      <td>NO</td>
      <td>PRI</td>
      <td></td>
      <td>Verse ID</td>
   </tr>
   <tr>
      <td>b</td>
      <td>int(11)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Book</td>
   </tr>
   <tr>
      <td>c</td>
      <td>int(11)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Chapter</td>
   </tr>
   <tr>
      <td>v</td>
      <td>int(11)</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Verse</td>
   </tr>
   <tr>
      <td>t</td>
      <td>text</td>
      <td>NO</td>
      <td></td>
      <td></td>
      <td>Text</td>
   </tr>
</table>
<br />



LICENSE:
-------------------
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.
