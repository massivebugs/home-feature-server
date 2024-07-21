CREATE TABLE IF NOT EXISTS `cashbunny_categories` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` INT(10) unsigned NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL DEFAULT "",
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_cashbunny_categories_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `cashbunny_accounts` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` INT(10) unsigned NOT NULL,
  `category_id` INT(10) unsigned NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL DEFAULT "",
  `balance` DECIMAL(19, 4) NOT NULL DEFAULT 0,
  `currency` VARCHAR(255) NOT NULL,
  `type` VARCHAR(255) NOT NULL,
  `order_index` SMALLINT(6) UNSIGNED NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_cashbunny_accounts_user_id_order_index`(`user_id`, `order_index`),
  CONSTRAINT `fk_cashbunny_accounts_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_cashbunny_accounts_category_id_cashbunny_categories_id` FOREIGN KEY (`category_id`) REFERENCES `cashbunny_categories`(`id`)
);

-- CREATE TABLE IF NOT EXISTS `cashbunny_accounts` (
--   `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
--   `user_id` INT(10) unsigned NOT NULL,
--   `account_id` INT(10) unsigned NOT NULL,
--   `description` VARCHAR(255) NOT NULL DEFAULT "",
--   `balance` DECIMAL(19, 4) NOT NULL DEFAULT 0,
--   `currency` VARCHAR(255) NOT NULL,
--   `type` VARCHAR(255) NOT NULL,
--   `order_index` SMALLINT(6) UNSIGNED NOT NULL,
--   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
--   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--   `deleted_at` timestamp NULL DEFAULT NULL,
--   PRIMARY KEY (`id`),
--   UNIQUE KEY (`name`),
--   UNIQUE KEY `uk_user_id_order_index`(`user_id`, `order_index`),
--   CONSTRAINT `fk_cashbunny_accounts_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
-- );