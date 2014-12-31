-- phpMyAdmin SQL Dump
-- version 4.0.6deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Aug 11, 2014 at 12:15 AM
-- Server version: 5.5.37-0ubuntu0.13.10.1
-- PHP Version: 5.5.3-1ubuntu2.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `bible`
--

-- --------------------------------------------------------

--
-- Table structure for table `key_genre_english`
--

DROP TABLE IF EXISTS `key_genre_english`;
CREATE TABLE `key_genre_english` (
  `g` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Genre ID',
  `n` varchar(255) NOT NULL COMMENT 'Name of genre',
  PRIMARY KEY (`g`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 COMMENT='Table mapping genre IDs to genre names for book table lookup' AUTO_INCREMENT=9 ;

--
-- Dumping data for table `key_genre_english`
--

INSERT INTO `key_genre_english` (`g`, `n`) VALUES
(1, 'Law'),
(2, 'History'),
(3, 'Wisdom'),
(4, 'Prophets'),
(5, 'Gospels'),
(6, 'Acts'),
(7, 'Epistles'),
(8, 'Apocalyptic');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
