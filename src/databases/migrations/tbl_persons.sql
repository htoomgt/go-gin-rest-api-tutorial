use db_test;

DROP TABLE IF EXISTS `tbl_persons`;
CREATE TABLE IF NOT EXISTS `tbl_persons`(
	`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `first_name` VARCHAR(128) NULL,
    `last_name` VARCHAR(128) NULL,
    `email` VARCHAR(512) NULL,
    `gender` VARCHAR(6) NULL,
    `age` INT NULL,
    `password` VARCHAR(128) NULL,
    `confirmed` BOOLEAN DEFAULT FALSE,
    `created_at` DATETIME NULL,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP
    
) ENGINE=INNODB;

explain tbl_persons;