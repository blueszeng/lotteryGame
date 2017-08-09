-- MySQL dump 10.13  Distrib 5.7.17, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: lottery
-- ------------------------------------------------------
-- Server version	5.7.18-log

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
-- Table structure for table `drawlotterys`
--

DROP TABLE IF EXISTS `drawlotterys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `drawlotterys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `periodNo` varchar(256) NOT NULL COMMENT '期号',
  `drawFirstNumber` varchar(3) NOT NULL COMMENT '开奖',
  `drawSecondNumber` varchar(3) NOT NULL COMMENT '开奖',
  `drawThirdNumber` varchar(3) NOT NULL COMMENT '开奖',
  `result` varchar(20) NOT NULL COMMENT '结果',
  PRIMARY KEY (`id`,`periodNo`)
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8 COMMENT='开奖';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `drawlotterys`
--

LOCK TABLES `drawlotterys` WRITE;
/*!40000 ALTER TABLE `drawlotterys` DISABLE KEYS */;
INSERT INTO `drawlotterys` VALUES (83,'2017080950','856','214','529','大'),(84,'2017080951','611','689','710','小'),(85,'2017080952','957','787','084','小'),(86,'2017080953','548','420','505','大'),(87,'2017080954','088','250','159','大'),(88,'2017080955','156','719','019','大'),(89,'2017080956','115','267','997','大'),(90,'2017080957','014','439','698','大'),(91,'2017080958','982','592','668','大'),(92,'2017080959','695','586','883','小'),(93,'2017080960','116','771','154','小'),(94,'2017080961','393','623','232','小'),(95,'2017080962','945','109','318','大'),(96,'2017080963','369','045','161','小'),(97,'2017080964','452','022','976','大'),(98,'2017080965','932','611','313','小'),(99,'2017080966','243','984','579','大'),(100,'2017080967','463','243','476','大'),(101,'2017080968','780','167','856','大'),(103,'2017080968','189','437','498','大'),(104,'2017080969','433','444','074','小');
/*!40000 ALTER TABLE `drawlotterys` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-08-09 19:04:55
