CREATE TABLE IF NOT EXISTS `cashbunny_accounts` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` INT(10) unsigned NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255),
  `balance` DECIMAL(19, 4) NOT NULL DEFAULT 0,
  `currency` VARCHAR(255) NOT NULL,
  `type` VARCHAR(255) NOT NULL,
  `order_index` SMALLINT(6) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`name`),
  UNIQUE KEY `uk_user_id_order_index`(`user_id`, `order_index`),
  CONSTRAINT `fk_cashbunny_accounts_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);