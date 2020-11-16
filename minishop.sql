-- MySQL dump 10.13  Distrib 5.7.31, for Linux (x86_64)
--
-- Host: localhost    Database: minishop
-- ------------------------------------------------------
-- Server version	5.7.31-0ubuntu0.18.04.1

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
-- Table structure for table `minishop_address`
--

DROP TABLE IF EXISTS `minishop_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_address` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `country_id` smallint(5) NOT NULL DEFAULT '0',
  `province_id` smallint(5) NOT NULL DEFAULT '0',
  `city_id` smallint(5) NOT NULL DEFAULT '0',
  `district_id` smallint(5) NOT NULL DEFAULT '0',
  `address` varchar(120) NOT NULL DEFAULT '',
  `mobile` varchar(60) NOT NULL DEFAULT '',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_admin`
--

DROP TABLE IF EXISTS `minishop_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(10) NOT NULL DEFAULT '''''',
  `password` varchar(255) NOT NULL DEFAULT '''''',
  `password_salt` varchar(255) NOT NULL DEFAULT '''''',
  `last_login_ip` varchar(60) NOT NULL DEFAULT '''''',
  `last_login_time` int(11) NOT NULL DEFAULT '0',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `avatar` varchar(255) NOT NULL DEFAULT '''''',
  `admin_role_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_attribute`
--

DROP TABLE IF EXISTS `minishop_attribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_attribute` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `attribute_category_id` int(11) unsigned NOT NULL DEFAULT '0',
  `name` varchar(60) NOT NULL DEFAULT '',
  `input_type` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `values` text NOT NULL,
  `sort_order` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `cat_id` (`attribute_category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_attribute_category`
--

DROP TABLE IF EXISTS `minishop_attribute_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_attribute_category` (
  `id` int(11) unsigned NOT NULL,
  `name` varchar(60) NOT NULL DEFAULT '',
  `enabled` tinyint(1) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_cart`
--

DROP TABLE IF EXISTS `minishop_cart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_cart` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `session_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_sn` varchar(60) NOT NULL DEFAULT '',
  `product_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_name` varchar(120) NOT NULL DEFAULT '',
  `market_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  `retail_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `number` smallint(5) unsigned NOT NULL DEFAULT '0',
  `goods_specifition_name_value` text NOT NULL COMMENT '规格属性组成的字符串，用来显示用',
  `goods_specifition_ids` varchar(60) NOT NULL DEFAULT '' COMMENT 'product表对应的goods_specifition_ids',
  `checked` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `list_pic_url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `session_id` (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_category`
--

DROP TABLE IF EXISTS `minishop_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(90) NOT NULL DEFAULT '',
  `keywords` varchar(255) NOT NULL DEFAULT '',
  `front_desc` varchar(255) NOT NULL DEFAULT '',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0',
  `sort_order` tinyint(1) unsigned NOT NULL DEFAULT '50',
  `show_index` tinyint(1) NOT NULL DEFAULT '0',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `banner_url` varchar(255) NOT NULL DEFAULT '',
  `icon_url` varchar(255) NOT NULL,
  `img_url` varchar(255) NOT NULL,
  `wap_banner_url` varchar(255) NOT NULL,
  `level` varchar(255) NOT NULL,
  `type` int(11) NOT NULL DEFAULT '0',
  `front_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1036005 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_channel`
--

DROP TABLE IF EXISTS `minishop_channel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_channel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `icon_url` varchar(255) NOT NULL DEFAULT '',
  `sort_order` int(4) unsigned NOT NULL DEFAULT '10',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_collect`
--

DROP TABLE IF EXISTS `minishop_collect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_collect` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `value_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0',
  `is_attention` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是关注',
  `type_id` int(2) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `goods_id` (`value_id`),
  KEY `is_attention` (`is_attention`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_comment`
--

DROP TABLE IF EXISTS `minishop_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `value_id` int(11) unsigned NOT NULL DEFAULT '0',
  `content` varchar(6550) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '储存为base64编码',
  `add_time` bigint(12) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `new_content` varchar(6550) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `id_value` (`value_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1006 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_comment_picture`
--

DROP TABLE IF EXISTS `minishop_comment_picture`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_comment_picture` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `comment_id` int(11) unsigned NOT NULL DEFAULT '0',
  `pic_url` varchar(255) NOT NULL DEFAULT '',
  `sort_order` tinyint(1) unsigned NOT NULL DEFAULT '5',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1121 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_feedback`
--

DROP TABLE IF EXISTS `minishop_feedback`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_feedback` (
  `msg_id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `user_name` varchar(60) NOT NULL DEFAULT '',
  `user_email` varchar(60) NOT NULL DEFAULT '',
  `msg_title` varchar(200) NOT NULL DEFAULT '',
  `msg_type` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `msg_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `msg_content` text NOT NULL,
  `msg_time` int(10) unsigned NOT NULL DEFAULT '0',
  `message_img` varchar(255) NOT NULL DEFAULT '0',
  `order_id` int(11) unsigned NOT NULL DEFAULT '0',
  `msg_area` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`msg_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_footprint`
--

DROP TABLE IF EXISTS `minishop_footprint`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_footprint` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `goods_id` int(11) NOT NULL DEFAULT '0',
  `add_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_goods`
--

DROP TABLE IF EXISTS `minishop_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_goods` (
  `id` int(11) unsigned NOT NULL,
  `category_id` int(11) unsigned NOT NULL DEFAULT '0',
  `goods_sn` varchar(60) NOT NULL DEFAULT '',
  `name` varchar(120) NOT NULL DEFAULT '',
  `brand_id` int(11) unsigned NOT NULL DEFAULT '0',
  `goods_number` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `keywords` varchar(255) NOT NULL DEFAULT '',
  `goods_brief` varchar(255) NOT NULL DEFAULT '',
  `goods_desc` text,
  `is_on_sale` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0',
  `sort_order` smallint(4) unsigned NOT NULL DEFAULT '100',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `attribute_category` int(11) unsigned NOT NULL DEFAULT '0',
  `counter_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '专柜价格',
  `extra_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '附加价格',
  `is_new` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `goods_unit` varchar(45) NOT NULL COMMENT '商品单位',
  `primary_pic_url` varchar(255) NOT NULL COMMENT '商品主图',
  `list_pic_url` varchar(255) NOT NULL COMMENT '商品列表图',
  `retail_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '零售价格',
  `sell_volume` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销售量',
  `primary_product_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主sku　product_id',
  `unit_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '单位价格，单价',
  `promotion_desc` varchar(255) NOT NULL,
  `promotion_tag` varchar(45) NOT NULL,
  `app_exclusive_price` decimal(10,2) unsigned NOT NULL COMMENT 'APP专享价',
  `is_app_exclusive` tinyint(1) unsigned NOT NULL COMMENT '是否是APP专属',
  `is_limited` tinyint(1) unsigned NOT NULL,
  `is_hot` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `goods_sn` (`goods_sn`),
  KEY `cat_id` (`category_id`),
  KEY `brand_id` (`brand_id`),
  KEY `goods_number` (`goods_number`),
  KEY `sort_order` (`sort_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_goods_attribute`
--

DROP TABLE IF EXISTS `minishop_goods_attribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_goods_attribute` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `attribute_id` int(11) unsigned NOT NULL DEFAULT '0',
  `value` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `goods_id` (`goods_id`),
  KEY `attr_id` (`attribute_id`)
) ENGINE=InnoDB AUTO_INCREMENT=872 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_goods_gallery`
--

DROP TABLE IF EXISTS `minishop_goods_gallery`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_goods_gallery` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `img_url` varchar(255) NOT NULL DEFAULT '',
  `img_desc` varchar(255) NOT NULL DEFAULT '',
  `sort_order` int(11) unsigned NOT NULL DEFAULT '5',
  PRIMARY KEY (`id`),
  KEY `goods_id` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=681 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_goods_issue`
--

DROP TABLE IF EXISTS `minishop_goods_issue`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_goods_issue` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `goods_id` text,
  `question` varchar(255) DEFAULT NULL,
  `answer` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_goods_specification`
--

DROP TABLE IF EXISTS `minishop_goods_specification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_goods_specification` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `specification_id` int(11) unsigned NOT NULL DEFAULT '0',
  `value` varchar(50) NOT NULL DEFAULT '',
  `pic_url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `goods_id` (`goods_id`),
  KEY `specification_id` (`specification_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='商品对应规格表值表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_keywords`
--

DROP TABLE IF EXISTS `minishop_keywords`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_keywords` (
  `keyword` varchar(90) NOT NULL DEFAULT '',
  `is_hot` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `sort_order` int(11) unsigned NOT NULL DEFAULT '100',
  `scheme _url` varchar(255) NOT NULL DEFAULT '' COMMENT '关键词的跳转链接',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='热闹关键词表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_order`
--

DROP TABLE IF EXISTS `minishop_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_order` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `order_sn` varchar(20) NOT NULL DEFAULT '',
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `order_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `shipping_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `pay_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `consignee` varchar(60) NOT NULL DEFAULT '',
  `country` smallint(5) unsigned NOT NULL DEFAULT '0',
  `province` smallint(5) unsigned NOT NULL DEFAULT '0',
  `city` smallint(5) unsigned NOT NULL DEFAULT '0',
  `district` smallint(5) unsigned NOT NULL DEFAULT '0',
  `address` varchar(255) NOT NULL DEFAULT '',
  `mobile` varchar(60) NOT NULL DEFAULT '',
  `postscript` varchar(255) NOT NULL DEFAULT '',
  `shipping_fee` decimal(10,2) NOT NULL DEFAULT '0.00',
  `pay_name` varchar(120) NOT NULL DEFAULT '',
  `pay_id` tinyint(3) NOT NULL DEFAULT '0',
  `actual_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '实际需要支付的金额',
  `integral` int(10) unsigned NOT NULL DEFAULT '0',
  `integral_money` decimal(10,2) NOT NULL DEFAULT '0.00',
  `order_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单总价',
  `goods_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品总价',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0',
  `confirm_time` int(11) unsigned NOT NULL DEFAULT '0',
  `pay_time` int(11) unsigned NOT NULL DEFAULT '0',
  `freight_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '配送费用',
  `coupon_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '使用的优惠券id',
  `parent_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `coupon_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `callback_status` enum('true','false') DEFAULT 'true',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_sn` (`order_sn`),
  KEY `user_id` (`user_id`),
  KEY `order_status` (`order_status`),
  KEY `shipping_status` (`shipping_status`),
  KEY `pay_status` (`pay_status`),
  KEY `pay_id` (`pay_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_order_express`
--

DROP TABLE IF EXISTS `minishop_order_express`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_order_express` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `shipper_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `shipper_name` varchar(120) NOT NULL DEFAULT '' COMMENT '物流公司名称',
  `shipper_code` varchar(60) NOT NULL DEFAULT '' COMMENT '物流公司代码',
  `logistic_code` varchar(20) NOT NULL DEFAULT '' COMMENT '快递单号',
  `traces` varchar(2000) NOT NULL DEFAULT '' COMMENT '物流跟踪信息',
  `is_finish` tinyint(1) NOT NULL DEFAULT '0',
  `request_count` int(11) DEFAULT '0' COMMENT '总查询次数',
  `request_time` int(11) DEFAULT '0' COMMENT '最近一次向第三方查询物流信息时间',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单物流信息表，发货时生成';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_order_goods`
--

DROP TABLE IF EXISTS `minishop_order_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_order_goods` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_name` varchar(120) NOT NULL DEFAULT '',
  `goods_sn` varchar(60) NOT NULL DEFAULT '',
  `product_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `number` smallint(5) unsigned NOT NULL DEFAULT '1',
  `market_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `retail_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `goods_specifition_name_value` text NOT NULL,
  `is_real` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `goods_specifition_ids` varchar(255) NOT NULL DEFAULT '',
  `list_pic_url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_product`
--

DROP TABLE IF EXISTS `minishop_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_product` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_specification_ids` varchar(50) NOT NULL DEFAULT '',
  `goods_sn` varchar(60) NOT NULL DEFAULT '',
  `goods_number` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `retail_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=245 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_region`
--

DROP TABLE IF EXISTS `minishop_region`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_region` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  `name` varchar(120) NOT NULL DEFAULT '',
  `type` tinyint(1) NOT NULL DEFAULT '2',
  `agency_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  KEY `region_type` (`type`),
  KEY `agency_id` (`agency_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4044 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_related_goods`
--

DROP TABLE IF EXISTS `minishop_related_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_related_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `related_goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_search_history`
--

DROP TABLE IF EXISTS `minishop_search_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_search_history` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `keyword` char(50) NOT NULL,
  `from` varchar(45) NOT NULL DEFAULT '' COMMENT '搜索来源，如PC、小程序、APP等',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT '搜索时间',
  `user_id` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_shipper`
--

DROP TABLE IF EXISTS `minishop_shipper`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_shipper` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '快递公司名称',
  `code` varchar(10) NOT NULL DEFAULT '' COMMENT '快递公司代码',
  `sort_order` int(11) NOT NULL DEFAULT '10' COMMENT '排序',
  PRIMARY KEY (`id`),
  UNIQUE KEY `minishop_shipper_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='快递公司';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_specification`
--

DROP TABLE IF EXISTS `minishop_specification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_specification` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL DEFAULT '',
  `sort_order` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='规格表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_user`
--

DROP TABLE IF EXISTS `minishop_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_user` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(60) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `gender` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `birthday` int(11) unsigned NOT NULL DEFAULT '0',
  `register_time` int(11) unsigned NOT NULL DEFAULT '0',
  `last_login_time` int(11) unsigned NOT NULL DEFAULT '0',
  `last_login_ip` varchar(255) NOT NULL DEFAULT '',
  `user_level_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `nickname` varchar(60) NOT NULL,
  `mobile` varchar(20) NOT NULL,
  `register_ip` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `weixin_openid` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `minishop_user_level`
--

DROP TABLE IF EXISTS `minishop_user_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `minishop_user_level` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-10-29 20:43:28
