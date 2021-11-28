drop database if exists `user`;
create database `user`;
use database `user`;
CREATE TABLE `info` (
  `name` char(10) DEFAULT NULL,
  `pass` char(20) DEFAULT NULL,
  `niname` char(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
CREATE TABLE `cookie` (
  `name` char(10) DEFAULT NULL,
  `value` char(20) DEFAULT NULL,
  `expire` char(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;