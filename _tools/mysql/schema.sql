CREATE TABLE `user`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) NOT NULL,
    `password` VARCHAR(80) NOT NULL,
    `role` VARCHAR(80) NOT NULL,
    `created` DATETIME(6) NOT NULL,
    `modified` DATETIME(6) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `task`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `title`    VARCHAR(128) NOT NULL,
    `status`   VARCHAR(20)  NOT NULL,
    `created`  DATETIME(6) NOT NULL,
    `modified` DATETIME(6) NOT NULL,
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;