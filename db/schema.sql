CREATE DATABASE memberships;

USE memberships;

CREATE TABLE `memberships`.`members` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `email` VARCHAR(45) NULL,
  `phone` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `memberships`.`cards` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `member_id` INT NOT NULL,
  `display_name` VARCHAR(45) NULL,
  `last_four` CHAR(4) NULL,
  `expiration_date` DATE NULL,
  `blocked_at` DATETIME NULL,
  `created_at` DATETIME NULL,
  `update_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_cards_member_id`
    FOREIGN KEY (`id`)
    REFERENCES `memberships`.`members` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

CREATE TABLE `memberships`.`payments` (
  `id` INT NOT NULL,
  `member_id` INT NOT NULL,
  `card_id` INT NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_payments_member_id`
    FOREIGN KEY (`id`)
    REFERENCES `memberships`.`members` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_payments_card_id`
    FOREIGN KEY (`id`)
    REFERENCES `memberships`.`cards` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);