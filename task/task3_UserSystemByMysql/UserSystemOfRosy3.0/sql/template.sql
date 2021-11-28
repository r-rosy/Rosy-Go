drop database if exists `ZeroGravity`;

create database `ZeroGravity`;

use  `ZeroGravity`;

DROP TABLE IF EXISTS `mytest`;

CREATE TABLE `mytest` (
  `name` char(10) DEFAULT NULL,
  `word` char(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

LOCK TABLES `rose` WRITE;

INSERT INTO `rose` VALUES ('rosy','1'),('jack','123'),('bob','123');

UNLOCK TABLES;