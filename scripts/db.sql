CREATE USER 'cow_R'@'%' IDENTIFIED BY 'Admin123';
CREATE USER 'cow_W'@'%' IDENTIFIED BY 'Admin123';
GRANT ALL PRIVILEGES ON * . * TO 'cow_R'@'%';
GRANT ALL PRIVILEGES ON * . * TO 'cow_W'@'%';
CREATE TABLE IF NOT EXISTS `c_group` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(100) NOT NULL,
  `debt` INT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `group_uk` (`code` ASC) VISIBLE)
ENGINE = InnoDB;
CREATE TABLE IF NOT EXISTS `c_team` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `group_id` INT NOT NULL,
  `user_id` VARCHAR(45) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_team_group1_idx` (`group_id` ASC) VISIBLE,
  INDEX `fk_team_user1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_team_group1`
    FOREIGN KEY (`group_id`)
    REFERENCES `c_group` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
CREATE TABLE IF NOT EXISTS `c_pay` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `team_id` INT NOT NULL,
  `description` VARCHAR(200) NOT NULL,
  `value` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_pay_team1_idx` (`team_id` ASC) VISIBLE,
  CONSTRAINT `fk_pay_team1`
    FOREIGN KEY (`team_id`)
    REFERENCES `c_team` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
