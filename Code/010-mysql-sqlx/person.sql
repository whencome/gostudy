CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(260) NOT NULL DEFAULT '',
  `sex` varchar(260) NOT NULL DEFAULT 'Male',
  `email` varchar(260) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;