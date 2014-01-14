Bible Versions and Cross-Reference Databases: mysql, sqlite, xml, csv, json.
===============

Here you will find the following bible versions in sql, sqlite, xml, csv, and json format:

- American Standard-ASV1901 (ASV)
- Bible in Basic English (BBE)
- Darby English Bible (DARBY)
- King James Version (KJV)
- Webster's Bible (WBT)
- World English Bible (WEB)
- Young's Literal Translation (YLT)

Also find a **Master SQL and SQLite file** (top level) containing all tables.


*Please let me know if any verses/text were inaccurately transferred in conversion.*

Each verse is accessed by a unique key, the combination of the BOOK+CHAPTER+VERSE id.

Example: 

**Genesis 1:1 (Genesis chapter 1, verse 1) = 01001001 (01 001 001)**

**Exodus 2:3 (Exodus chapter 2, verse 3) = 02002003 (02 002 003)** 

The verse-id system is used for faster, simplified queries. For instance:

***01001001*** - ***02001005*** would capture all verses between Genesis 1:1 through Exodus 1:5. 

Written simply:

<code>
SELECT *
FROM bible.t_asv
WHERE id
BETWEEN 01001001
AND 02001005
</code>

There is also a number-to-book key (**key_english** table), a cross-reference list (**cross_reference** table), and a bible key containing meta information about the included translations (**bible_version_key** table). ***See below SQL table layout.*** These tables work together providing you a great basis for a bible-reading and cross-referencing app. 

While its expected that your programs would use the verse-id system, *book #, chapter #, and verse #* columns have been included in the bible versions tables.

Special Note on Cross-Reference Table
-------------------
A very special and valuable addition to these databases is the extensive cross-reference table. Its combined from trustworthy sources. Its extremely useful in bible study for discovering related scriptures. For any given verse, you simply query ***vid*** (verse id), and a list of rows will be returned. Each of those rows has a rank (r) for relevance, start-verse (sv), and end verse (ev) if there is one. 

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
      <th>Default
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
      <th>Default
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
      <th>Default
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
</table>
<br />
<h3>t_<version abbreviation></h3>
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

