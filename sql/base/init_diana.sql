-- 楓谷幣值
CREATE TABLE `currency` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server` varchar(10) DEFAULT NULL,
  `value` decimal(18,2) DEFAULT NULL,
  `added_time` date DEFAULT NULL,
  `title` text,
  `url` text,
  `abnormal` int(4) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `v` (`value`,`server`,`added_time`)
) ENGINE=InnoDB AUTO_INCREMENT=23265 DEFAULT CHARSET=utf8mb4;

-- 待看文章
CREATE TABLE `web_site` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(200) DEFAULT NULL,
  `tag` varchar(50) DEFAULT NULL,
  `added_time` time DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Line使用者
CREATE TABLE `line_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(100) DEFAULT NULL,
  `added_time` time DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 我的line id
INSERT INTO `line_user` (`user_id`,`added_time`) VALUES ('Ud446ab1ac3d5ea31f973aa8967f15a78',NOW());

-- kktix活動資訊
CREATE TABLE `kktix_activity` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` text CHARSET utf8mb4,
  `url` text CHARSET utf8mb4,
  `introduction` text CHARSET utf8mb4,
  `category` text CHARSET utf8mb4,
  `create_time` text CHARSET utf8mb4,
  `ticket_status` text CHARSET utf8mb4,
  `participate_number` text CHARSET utf8mb4,
  `activity_time` text CHARSET utf8mb4,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tourist_attraction_list` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `url` varchar(200) NOT NULL,
   `place` varchar(45) DEFAULT NULL,
   `country` varchar(45) DEFAULT NULL,
   `location` varchar(45) DEFAULT NULL,
   `address` varchar(45) DEFAULT NULL,
   `activity_time` varchar(200) DEFAULT NULL,
   PRIMARY KEY (`id`),
   UNIQUE KEY `url_UNIQUE` (`url`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;