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