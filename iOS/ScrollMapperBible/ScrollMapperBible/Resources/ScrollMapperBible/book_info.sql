CREATE TABLE IF NOT EXISTS "book_info" (
	"order"	INTEGER NOT NULL UNIQUE,
	"title_short"	TEXT NOT NULL UNIQUE,
	"title_full"	TEXT NOT NULL UNIQUE,
	"abbreviation"	TEXT NOT NULL UNIQUE,
	"category"	TEXT NOT NULL,
	"otnt"	TEXT NOT NULL,
    "chapters" 
);

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (1, 'Genesis', 'The First Book of Moses Called Genesis', 'Gen.', 'Law', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (2, 'Exodus', 'The Second Book of Moses Called Exodus', 'Ex.', 'Law', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (3, 'Leviticus', 'The Third Book of Moses Called Leviticus', 'Lev.', 'Law', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (4, 'Numbers', 'The Fourth Book of Moses Called Numbers', 'Num.', 'Law', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (5, 'Deuteronomy', 'The Fifth Book of Moses Called Deuteronomy', 'Deut.', 'Law', 'OT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (6, 'Joshua', 'The Book of Joshua', 'Josh.', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (7, 'Judges', 'The Book of Judges', 'Judg.', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (8, 'Ruth', 'The Book of Ruth', 'Ruth', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (9, '1 Samuel', 'The First Book of Samuel', '1 Sam.', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (10, '2 Samuel', 'The Second Book of Samuel', '2 Sam.', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (11, '1 Kings', 'The First Book of Kings', '1 Kings', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (12, '2 Kings', 'The Second Book of Kings', '2 Kings', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (13, '1 Chronicles', 'The First Book of Chronicles', '1 Chron.', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (14, '2 Chronicles', 'The Second Book of Chronicles', '2 Chron.', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (15, 'Ezra', 'The Book of Ezra', 'Ezra', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (16, 'Nehemiah', 'The Book of Nehemiah', 'Neh.', 'Old Testament Narrative', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (17, 'Esther', 'The Book of Esther', 'Est.', 'Old Testament Narrative', 'OT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (18, 'Job', 'The Book of Job', 'Job', 'Wisdom Literature', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (19, 'Psalms', 'The Book of Psalms', 'Ps.', 'Wisdom Literature', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (20, 'Proverbs', 'The Book of Proverbs', 'Prov.', 'Wisdom Literature', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (21, 'Ecclesiastes', 'The Book of Ecclesiastes', 'Eccles.', 'Wisdom Literature', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (22, 'Song of Solomon', 'Song of Solomon', 'Song', 'Wisdom Literature', 'OT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (23, 'Isaiah', 'The Book of Isaiah', 'Isa.', 'Major Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (24, 'Jeremiah', 'The Book of Jeremiah', 'Jer.', 'Major Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (25, 'Lamentations', 'The Book of Lamentations', 'Lam.', 'Major Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (26, 'Ezekiel', 'The Book of Ezekiel', 'Ezek.', 'Major Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (27, 'Daniel', 'The Book of Daniel', 'Dan.', 'Major Prophets', 'OT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (28, 'Hosea', 'The Book of Hosea', 'Hos.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (29, 'Joel', 'The Book of Joel', 'Joel', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (30, 'Amos', 'The Book of Amos', 'Amos', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (31, 'Obadiah', 'The Book of Obadiah', 'Obad.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (32, 'Jonah', 'The Book of Jonah', 'Jonah', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (33, 'Micah', 'The Book of Micah', 'Mic.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (34, 'Nahum', 'The Book of Nahum', 'Nah.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (35, 'Habakkuk', 'The Book of Habakkuk', 'Hab.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (36, 'Zephaniah', 'The Book of Zephaniah', 'Zeph.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (37, 'Haggai', 'The Book of Haggai', 'Hag.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (38, 'Zechariah', 'The Book of Zechariah', 'Zech.', 'Minor Prophets', 'OT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (39, 'Malachi', 'The Book of Malachi', 'Mal.', 'Minor Prophets', 'OT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (40, 'Matthew', 'The Gospel According to Matthew', 'Matt.', 'New Testament Narrative', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (41, 'Mark', 'The Gospel According to Mark', 'Mark', 'New Testament Narrative', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (42, 'Luke', 'The Gospel According to Luke', 'Luke', 'New Testament Narrative', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (43, 'John', 'The Gospel According to John', 'John', 'New Testament Narrative', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (44, 'Acts', 'The Acts of the Apostles', 'Acts', 'New Testament Narrative', 'NT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (45, 'Romans', 'The Epistle of Paul to the Romans', 'Rom.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (46, '1 Corinthians', 'The First Epistle of Paul to the Corinthians', '1 Cor.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (47, '2 Corinthians', 'The Second Epistle of Paul to the Corinthians', '2 Cor.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (48, 'Galatians', 'The Epistle of Paul to the Galatians', 'Gal.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (49, 'Ephesians', 'The Epistle of Paul to the Ephesians', 'Eph.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (50, 'Philippians', 'The Epistle of Paul to the Philippians', 'Phil.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (51, 'Colossians', 'The Epistle of Paul to the Colossians', 'Col.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (52, '1 Thessalonians', 'The First Epistle of Paul to the Thessalonians', '1 Thess.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (53, '2 Thessalonians', 'The Second Epistle of Paul to the Thessalonians', '2 Thess.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (54, '1 Timothy', 'The First Epistle of Paul to Timothy', '1 Tim.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (55, '2 Timothy', 'The Second Epistle of Paul to Timothy', '2 Tim.', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (56, 'Titus', 'The Epistle of Paul to the Titus', 'Titus', 'Pauline Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (57, 'Philemon', 'The Epistle of Paul to the Philemon', 'Philem.', 'Pauline Epistles', 'NT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (58, 'Hebrews', 'The Epistle to the Hebrews', 'Heb.', 'General Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (59, 'James', 'The General Epistle of James', 'James', 'General Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (60, '1 Peter', 'The First Epistle of Peter', '1 Pet.', 'General Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (61, '2 Peter', 'The Second Epistle of Peter', '2 Pet.', 'General Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (62, '1 John', 'The First Epistle of John', '1 John', 'General Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (63, '2 John', 'The Second Epistle of John', '2 John', 'General Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (64, '3 John', 'The Third Epistle of John', '3 John', 'General Epistles', 'NT');
insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (65, 'Jude', 'The Epistle of Jude', 'Jude', 'General Epistles', 'NT');

insert into books ('order', 'title_short', 'title_full', 'abbreviation', 'category', 'otnt') values (66, 'Revelation', 'The Book of Revelation', 'Rev.', 'Apocalyptic Epistle', 'NT');
