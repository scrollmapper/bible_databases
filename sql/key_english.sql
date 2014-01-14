-- MySQL dump 10.13  Distrib 5.5.34, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: bible
-- ------------------------------------------------------
-- Server version	5.5.34-0ubuntu0.13.10.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `key_english`
--

DROP TABLE IF EXISTS `key_english`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `key_english` (
  `b` int(11) NOT NULL COMMENT 'Book #',
  `n` text NOT NULL COMMENT 'Name',
  PRIMARY KEY (`b`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `key_english`
--
-- WHERE:  1 limit 1000000

LOCK TABLES `key_english` WRITE;
/*!40000 ALTER TABLE `key_english` DISABLE KEYS */;
INSERT INTO `key_english` VALUES (1,'Genesis'),(2,'Exodus'),(3,'Leviticus'),(4,'Numbers'),(5,'Deuteronomy'),(6,'Joshua'),(7,'Judges'),(8,'Ruth'),(9,'1 Samuel'),(10,'2 Samuel'),(11,'1 Kings'),(12,'2 Kings'),(13,'1 Chronicles'),(14,'2 Chronicles'),(15,'Ezra'),(16,'Nehemiah'),(17,'Esther'),(18,'Job'),(19,'Psalms'),(20,'Proverbs'),(21,'Ecclesiastes'),(22,'Song of Solomon'),(23,'Isaiah'),(24,'Jeremiah'),(25,'Lamentations'),(26,'Ezekiel'),(27,'Daniel'),(28,'Hosea'),(29,'Joel'),(30,'Amos'),(31,'Obadiah'),(32,'Jonah'),(33,'Micah'),(34,'Nahum'),(35,'Habakkuk'),(36,'Zephaniah'),(37,'Haggai'),(38,'Zechariah'),(39,'Malachi'),(40,'Matthew'),(41,'Mark'),(42,'Luke'),(43,'John'),(44,'Acts'),(45,'Romans'),(46,'1 Corinthians'),(47,'2 Corinthians'),(48,'Galatians'),(49,'Ephesians'),(50,'Philippians'),(51,'Colossians'),(52,'1 Thessalonians'),(53,'2 Thessalonians'),(54,'1 Timothy'),(55,'2 Timothy'),(56,'Titus'),(57,'Philemon'),(58,'Hebrews'),(59,'James'),(60,'1 Peter'),(61,'2 Peter'),(62,'1 John'),(63,'2 John'),(64,'3 John'),(65,'Jude'),(66,'Revelation');
/*!40000 ALTER TABLE `key_english` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-01-13 19:46:51
