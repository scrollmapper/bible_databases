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
-- Table structure for table `bible_version_key`
--

DROP TABLE IF EXISTS `bible_version_key`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bible_version_key` (
  `id` int(3) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `table` text NOT NULL COMMENT 'Database Table Name ',
  `abbreviation` text NOT NULL COMMENT 'Version Abbreviation',
  `language` text NOT NULL COMMENT 'Language of bible translation (used for language key tables)',
  `version` text NOT NULL COMMENT 'Version Name',
  `info_text` text NOT NULL COMMENT 'About / Info',
  `info_url` text NOT NULL COMMENT 'Info URL',
  `publisher` text NOT NULL COMMENT 'Publisher',
  `copyright` text NOT NULL COMMENT 'Copyright ',
  `copyright_info` text NOT NULL COMMENT 'Extended Copyright info',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1 COMMENT='This is the general translation information and db links';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bible_version_key`
--
-- WHERE:  1 limit 1000000

LOCK TABLES `bible_version_key` WRITE;
/*!40000 ALTER TABLE `bible_version_key` DISABLE KEYS */;
INSERT INTO `bible_version_key` VALUES (001,'t_asv','ASV','english','American Standard-ASV1901','','http://en.wikipedia.org/wiki/American_Standard_Version','','Public Domain',''),(002,'t_bbe','BBE','english','Bible in Basic English','','http://en.wikipedia.org/wiki/Bible_in_Basic_English','','Public Domain',''),(003,'t_dby','DARBY','english','Darby English Bible','','http://en.wikipedia.org/wiki/Darby_Bible','','Public Domain',''),(004,'t_kjv','KJV','english','King James Version','','http://en.wikipedia.org/wiki/King_James_Version','','Public Domain',''),(005,'t_wbt','WBT','english','Webster\'s Bible','','http://en.wikipedia.org/wiki/Webster%27s_Revision','','Public Domain',''),(006,'t_web','WEB','english','World English Bible','','http://en.wikipedia.org/wiki/World_English_Bible','','Public Domain',''),(007,'t_ylt','YLT','english','Young\'s Literal Translation','','http://en.wikipedia.org/wiki/Young%27s_Literal_Translation','','Public Domain','');
/*!40000 ALTER TABLE `bible_version_key` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-01-13 19:46:49
