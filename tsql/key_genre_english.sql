
/***************************************************************************
    Author: Daniel Bouchard
    Date:   Nov. 18th, 2017

    Notes:
	   This is an adaptation of the mysql script key_genre_english.sql for TSQL
***************************************************************************/

-- Create Table

DROP TABLE IF EXISTS 
     key_genre_english;

CREATE TABLE key_genre_english (
    g INT IDENTITY(1,1) PRIMARY KEY NOT NULL,
    n VARCHAR(255) NOT NULL
);

-- Add Comments

EXEC sp_addextendedproperty 
     @name = N'Description',
     @value = 'Table mapping genre IDs to genre names for book table lookup',
     @level0type = N'Schema',
     @level0name = 'dbo',
     @level1type = N'Table',
     @level1name = 'key_genre_english';

EXEC sp_addextendedproperty 
     @name = N'Description',
     @value = 'Genre ID',
     @level0type = N'Schema',
     @level0name = 'dbo',
     @level1type = N'Table',
     @level1name = 'key_genre_english',
     @level2type = N'Column',
     @level2name = 'g';

EXEC sp_addextendedproperty 
     @name = N'Description',
     @value = 'Name of genre',
     @level0type = N'Schema',
     @level0name = 'dbo',
     @level1type = N'Table',
     @level1name = 'key_genre_english',
     @level2type = N'Column',
     @level2name = 'n';
GO

-- Insert Data

SET IDENTITY_INSERT dbo.key_genre_english ON;

INSERT INTO                   key_genre_english (
       g,
       n
) 
VALUES (
       1,'Law'
 ), (
       2,'History'
 ), (
       3,'Wisdom'
 ), (
       4,'Prophets'
 ), (
       5,'Gospels'
 ), (
       6,'Acts'
 ), (
       7,'Epistles'
 ), (
       8,'Apocalyptic'
 );

SET IDENTITY_INSERT dbo.key_genre_english OFF;
