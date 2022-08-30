ALTER TABLE `inspirit`.`users` 
ADD COLUMN `mobile_number` CHAR(15) NULL AFTER `password`;

ALTER TABLE `inspirit`.`users` 
ADD UNIQUE INDEX `mobile_number_UNIQUE` (`mobile_number` ASC) VISIBLE;
;