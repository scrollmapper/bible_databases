/*

    Author: Daniel Bouchard
    Date:   Nov. 18th, 2017

    Notes:
	   This is an adaptation of the mysql script bible_version_key.sql for TSQL

*/

-- Create Table

DROP TABLE IF EXISTS bible_version_key;

CREATE TABLE bible_version_key
(
  id             INT IDENTITY( 1, 1 ) PRIMARY KEY,
  [table]        TEXT NOT NULL,
  abbreviation   TEXT NOT NULL,
  [language]     TEXT NOT NULL,
  [version]      TEXT NOT NULL,
  info_text      TEXT NOT NULL,
  info_url       TEXT NOT NULL,
  publisher      TEXT NOT NULL,
  copyright      TEXT NOT NULL,
  copyright_info TEXT NOT NULL);
GO

-- Add Comments

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Database Table Name',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='table';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Version Abbreviation',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='abbreviation';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Language of bible translation (used for language key tables)',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='language';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Version Name',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='version';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='About / Info',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='info_text';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Info URL',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='info_url';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Publisher',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='publisher';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Copyright',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='copyright';

EXEC sp_addextendedproperty
     @name=N'Description',
     @value='Extended Copyright info',
     @level0type=N'Schema',
     @level0name='dbo',
     @level1type=N'Table',
     @level1name='bible_version_key',
     @level2type=N'Column',
     @level2name='copyright_info';
GO

SET IDENTITY_INSERT dbo.bible_version_key ON;

INSERT INTO dbo.bible_version_key
(
    id,
    [table],
    abbreviation,
    language,
    version,
    info_text,
    info_url,
    publisher,
    copyright,
    copyright_info
)
VALUES
(
  001, 't_asv', 'ASV', 'english', 'American Standard-ASV1901', '', 'http://en.wikipedia.org/wiki/American_Standard_Version', '', 'Public Domain', ''),
(
  002, 't_bbe', 'BBE', 'english', 'Bible in Basic English', '', 'http://en.wikipedia.org/wiki/Bible_in_Basic_English', '', 'Public Domain', ''),
(
  003, 't_dby', 'DARBY', 'english', 'Darby English Bible', '', 'http://en.wikipedia.org/wiki/Darby_Bible', '', 'Public Domain', ''),
(
  004, 't_kjv', 'KJV', 'english', 'King James Version', '', 'http://en.wikipedia.org/wiki/King_James_Version', '', 'Public Domain', ''),
(
  005, 't_wbt', 'WBT', 'english', 'Webster''s Bible', '', 'http://en.wikipedia.org/wiki/Webster%27s_Revision', '', 'Public Domain', ''),
(
  006, 't_web', 'WEB', 'english', 'World English Bible', '', 'http://en.wikipedia.org/wiki/World_English_Bible', '', 'Public Domain', ''),
(
  007, 't_ylt', 'YLT', 'english', 'Young''s Literal Translation', '', 'http://en.wikipedia.org/wiki/Young%27s_Literal_Translation', '', 'Public Domain', '');

  SET IDENTITY_INSERT dbo.bible_version_key OFF;