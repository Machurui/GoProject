CREATE DATABASE go;
USE go;
DROP TABLE IF EXISTS journal;
CREATE TABLE journal (
  id            INT AUTO_INCREMENT NOT NULL,
  dh            datetime NOT NULL,
  mf            VARCHAR(255) NOT NULL,
  argument     VARCHAR(255) NOT NULL,
  statut        VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);
