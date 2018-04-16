/*
SQLyog Ultimate v12.4.1 (64 bit)
MySQL - 5.6.36 : Database - youzan
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`youzan` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `youzan`;

/*Table structure for table `city` */

DROP TABLE IF EXISTS `city`;

CREATE TABLE `city` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '城市编号',
  `seq` int(10) unsigned NOT NULL COMMENT '省份编号',
  `city_name` varchar(25) DEFAULT NULL COMMENT '城市名称',
  `description` varchar(25) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_auth_token` */

DROP TABLE IF EXISTS `yz_auth_token`;

CREATE TABLE `yz_auth_token` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `openid` varchar(30) DEFAULT NULL,
  `access_token` varchar(155) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `expires_in` int(11) NOT NULL DEFAULT '7200',
  `refresh_token` varchar(155) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `scope` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'snsapi_base',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='微信认证表';

/*Table structure for table `yz_category` */

DROP TABLE IF EXISTS `yz_category`;

CREATE TABLE `yz_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `pid` int(11) NOT NULL,
  `seq` int(4) NOT NULL,
  `url` varchar(45) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url_UNIQUE` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_email_verify` */

DROP TABLE IF EXISTS `yz_email_verify`;

CREATE TABLE `yz_email_verify` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(45) NOT NULL,
  `code` varchar(45) NOT NULL,
  `mail_type` varchar(20) NOT NULL,
  `duration` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_level_config` */

DROP TABLE IF EXISTS `yz_level_config`;

CREATE TABLE `yz_level_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL COMMENT '类型 1:图文 2:段子 3:吐槽',
  `level` int(11) NOT NULL COMMENT '级别',
  `title` varchar(45) NOT NULL COMMENT '级别名称',
  `min_score` int(11) NOT NULL COMMENT '最低分数',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='级别配置';

/*Table structure for table `yz_oauth` */

DROP TABLE IF EXISTS `yz_oauth`;

CREATE TABLE `yz_oauth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `access_token` varchar(45) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_operation_history` */

DROP TABLE IF EXISTS `yz_operation_history`;

CREATE TABLE `yz_operation_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `type` varchar(1) NOT NULL,
  `comment` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `yz_relation_ship` */

DROP TABLE IF EXISTS `yz_relation_ship`;

CREATE TABLE `yz_relation_ship` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `focus_id` int(11) NOT NULL COMMENT '关注者id',
  `target_id` int(11) NOT NULL COMMENT '被关注者id',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '关注状态0:已取消关注 2：已关注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='关注和被关注表';

/*Table structure for table `yz_sign_history` */

DROP TABLE IF EXISTS `yz_sign_history`;

CREATE TABLE `yz_sign_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_tech` */

DROP TABLE IF EXISTS `yz_tech`;

CREATE TABLE `yz_tech` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(4) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `creator_id` int(11) DEFAULT NULL,
  `content` text NOT NULL,
  `prise_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
  `diss_num` int(11) NOT NULL DEFAULT '0' COMMENT '鄙视数',
  `reply_num` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
  `view_times` int(11) NOT NULL DEFAULT '0' COMMENT '查看数',
  `last_reply_user_id` int(11) DEFAULT NULL,
  `last_reply_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` varchar(1) DEFAULT 'C' COMMENT 'C: 创建，P：发布，B:打回，S:保存',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_tech_reply` */

DROP TABLE IF EXISTS `yz_tech_reply`;

CREATE TABLE `yz_tech_reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tech_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `content` text NOT NULL,
  `status` char(1) NOT NULL DEFAULT 'A' COMMENT 'A：评论有效 D：评论被删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_user` */

DROP TABLE IF EXISTS `yz_user`;

CREATE TABLE `yz_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `roles` varchar(20) NOT NULL,
  `client_id` varchar(20) NOT NULL,
  `client_secret` varchar(20) NOT NULL,
  `scope` varchar(45) NOT NULL,
  `open_id` varchar(100) DEFAULT NULL,
  `grant_type` varchar(45) DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  `token_type` varchar(45) DEFAULT NULL,
  `expires_in` int(11) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `telephone` varchar(11) DEFAULT NULL,
  `head_image` varchar(255) DEFAULT NULL,
  `verified` char(1) NOT NULL DEFAULT 'N',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

/*Table structure for table `yz_user_bak` */

DROP TABLE IF EXISTS `yz_user_bak`;

CREATE TABLE `yz_user_bak` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL COMMENT '用户昵称',
  `password` varchar(45) DEFAULT NULL COMMENT '密码',
  `qq` varchar(45) DEFAULT NULL COMMENT 'QQ号',
  `open_id` varchar(45) DEFAULT NULL COMMENT '微信id',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '用户活跃状态 0：不正常 1:正常 ',
  `verified` int(11) DEFAULT '0' COMMENT '实名认证0：未认证 1：已认证',
  `tw_score` int(11) NOT NULL DEFAULT '0' COMMENT '图文积分',
  `piece_score` int(11) NOT NULL DEFAULT '0' COMMENT '段子积分',
  `mock_score` int(11) NOT NULL DEFAULT '0' COMMENT '吐槽积分',
  `exchanged_score` int(11) NOT NULL DEFAULT '0' COMMENT '已兑换积分',
  `exchanged_coin` int(11) NOT NULL DEFAULT '0' COMMENT '已兑换金额',
  `last_login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后登录时间',
  `email` varchar(45) DEFAULT NULL,
  `art_level` int(11) NOT NULL DEFAULT '0' COMMENT '图文级别',
  `piece_level` int(11) NOT NULL DEFAULT '0' COMMENT '段子手段位',
  `mock_level` int(11) NOT NULL DEFAULT '0' COMMENT '吐槽级别',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='用户表';

/*Table structure for table `yz_wechat` */

DROP TABLE IF EXISTS `yz_wechat`;

CREATE TABLE `yz_wechat` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(100) NOT NULL,
  `access_token` varchar(255) NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  `scope` varchar(25) NOT NULL,
  `expires_in` int(11) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
