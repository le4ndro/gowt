CREATE DATABASE gowtdb CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE gowtdb;

CREATE TABLE tools (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  category varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  url varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  rating int(11) DEFAULT NULL,
  notes text COLLATE utf8_unicode_ci,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;