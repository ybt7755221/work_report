# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.30)
# Database: work_reports
# Generation Time: 2020-10-23 02:52:16 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table wr_projects
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wr_projects`;

CREATE TABLE `wr_projects` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project_name` varchar(128) NOT NULL DEFAULT '' COMMENT '项目名称',
  `test_time` date NOT NULL DEFAULT '1970-01-01' COMMENT '预计提测时间',
  `publish_time` date NOT NULL DEFAULT '1970-01-01' COMMENT '预计上线时间',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_project_name` (`project_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table wr_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wr_users`;

CREATE TABLE `wr_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) NOT NULL DEFAULT '13000000000' COMMENT '手机号',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table wr_works
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wr_works`;

CREATE TABLE `wr_works` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `project_id` int(11) unsigned NOT NULL COMMENT '项目id',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '工作title',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT 'url地址',
  `progress` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '工作进度',
  `work_type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1-前端;2-后端',
  `backup` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created` date NOT NULL,
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id_A_created` (`user_id`,`created`),
  KEY `idx_created` (`created`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
