CREATE TABLE `maple_bulletin` (
   `id` int NOT NULL AUTO_INCREMENT,
   `url` varchar(200) DEFAULT NULL,
   `date` date DEFAULT NULL,
   `title` varchar(100) CHARACTER SET utf8mb4,
   `category` varchar(45) CHARACTER SET utf8mb4,
   PRIMARY KEY (`id`),
   UNIQUE KEY `url_UNIQUE` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=152 DEFAULT CHARSET=utf8mb4;
