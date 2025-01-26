-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: xyz
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `asset`
--

DROP TABLE IF EXISTS `asset`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `asset` (
  `id` varchar(36) NOT NULL,
  `nama_asset` varchar(100) DEFAULT NULL,
  `otr` bigint unsigned DEFAULT NULL,
  `platform_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `asset`
--

LOCK TABLES `asset` WRITE;
/*!40000 ALTER TABLE `asset` DISABLE KEYS */;
INSERT INTO `asset` VALUES ('17b28329-b593-461c-947d-000943457261','Mobil Dealer',150000000,'ac058b1e-82b4-4521-9ceb-73adb2e30c04'),('8f1d2e76-64de-4f13-8571-d6f314b29d9b','Mobil Ecommerce',12000000,'b0990fce-c0ff-4a17-91ea-b1e6465ef2c8'),('ac058b1e-82b4-4521-9ceb-73adb2e30c04','Mobil XYZ',200000000,'51c4f3c1-9f83-433b-9576-596ef9187611');
/*!40000 ALTER TABLE `asset` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `konsumen`
--

DROP TABLE IF EXISTS `konsumen`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `konsumen` (
  `nik` varchar(16) NOT NULL,
  `full_name` varchar(100) DEFAULT NULL,
  `legal_name` longtext,
  `tempat_lahir` varchar(100) DEFAULT NULL,
  `tgl_lahir` date DEFAULT NULL,
  `map` varchar(100) DEFAULT NULL,
  `nama_pic` varchar(100) DEFAULT NULL,
  `created_by` longtext,
  `updated_by` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `foto_ktp` varchar(100) DEFAULT NULL,
  `foto_selfie` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`nik`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `konsumen`
--

LOCK TABLES `konsumen` WRITE;
/*!40000 ALTER TABLE `konsumen` DISABLE KEYS */;
INSERT INTO `konsumen` VALUES ('3173070203031001','Melly Anastasya','Melly Anastasya','Jakarta','2005-05-05',NULL,NULL,'','','2025-01-26 21:23:46.921','2025-01-26 22:44:15.909','./uploads/WhatsApp Image 2025-01-26 at 21.33.48 (1).jpeg','');
/*!40000 ALTER TABLE `konsumen` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `limit`
--

DROP TABLE IF EXISTS `limit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `limit` (
  `id` varchar(36) NOT NULL,
  `tenor` tinyint unsigned DEFAULT NULL,
  `limit` double DEFAULT NULL,
  `konsumen_nik` varchar(16) DEFAULT NULL,
  `created_by` longtext,
  `updated_by` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `limit`
--

LOCK TABLES `limit` WRITE;
/*!40000 ALTER TABLE `limit` DISABLE KEYS */;
/*!40000 ALTER TABLE `limit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `limits`
--

DROP TABLE IF EXISTS `limits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `limits` (
  `id` varchar(36) NOT NULL,
  `tenor` tinyint unsigned DEFAULT NULL,
  `limit` double DEFAULT NULL,
  `konsumen_nik` varchar(16) DEFAULT NULL,
  `created_by` longtext,
  `updated_by` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `limits`
--

LOCK TABLES `limits` WRITE;
/*!40000 ALTER TABLE `limits` DISABLE KEYS */;
INSERT INTO `limits` VALUES ('22c4a9e0-903c-40ea-908d-e04aa24aafac',1,50000000,'3173070203031001','Rahmat','',NULL,'2025-01-26 22:24:50.392');
/*!40000 ALTER TABLE `limits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `platform`
--

DROP TABLE IF EXISTS `platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `platform` (
  `id` varchar(36) NOT NULL,
  `nama_platform` varchar(100) DEFAULT NULL,
  `jumlah_bunga` tinyint unsigned DEFAULT NULL,
  `admin_fee` bigint unsigned DEFAULT NULL,
  `created_by` longtext,
  `updated_by` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `platform`
--

LOCK TABLES `platform` WRITE;
/*!40000 ALTER TABLE `platform` DISABLE KEYS */;
INSERT INTO `platform` VALUES ('51c4f3c1-9f83-433b-9576-596ef9187611','XYZ',6,45000,NULL,NULL,NULL,NULL),('ac058b1e-82b4-4521-9ceb-73adb2e30c04','Dealer',9,65000,NULL,NULL,NULL,NULL),('b0990fce-c0ff-4a17-91ea-b1e6465ef2c8','Ecommerce',3,25000,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `platform` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaksi`
--

DROP TABLE IF EXISTS `transaksi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaksi` (
  `no_kontrak` varchar(16) NOT NULL,
  `nama_asset` varchar(100) DEFAULT NULL,
  `jumlah_cicilan` tinyint unsigned DEFAULT NULL,
  `jumlah_bunga` tinyint unsigned DEFAULT NULL,
  `otr` double DEFAULT NULL,
  `admin_fee` bigint unsigned DEFAULT NULL,
  `konsumen_nik` varchar(16) DEFAULT NULL,
  `asset_id` varchar(36) DEFAULT NULL,
  `created_by` longtext,
  `updated_by` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`no_kontrak`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaksi`
--

LOCK TABLES `transaksi` WRITE;
/*!40000 ALTER TABLE `transaksi` DISABLE KEYS */;
INSERT INTO `transaksi` VALUES ('TRX-001','Motor',1,4,100000000,0,'3173070203031001','','','','2025-01-26 22:24:50.388','2025-01-26 22:24:50.388'),('TRX-002','Mobil',1,4,100000000,0,'3173070203031001','','','','2025-01-26 22:24:50.392','2025-01-26 22:24:50.392'),('TRX-003','Motor 2',1,4,100000000,0,'3173070203031001','','','','2025-01-26 22:24:50.396','2025-01-26 22:24:50.396');
/*!40000 ALTER TABLE `transaksi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(191) NOT NULL,
  `nik` longtext,
  `name` longtext,
  `password` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('0f2a1c42-a8fe-47af-ba1e-bf90a8c04d82','3173070203031001','Rahmat Afriyadi','$2a$08$SedXCsmbVRDxU82xXMle5.KHW1qzsbRTmP8bcLGdVmXfHDFphPjcu');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-01-26 22:53:33
