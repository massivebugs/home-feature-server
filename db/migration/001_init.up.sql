CREATE TABLE
  IF NOT EXISTS `users` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_users_name` (`name`)
  );

CREATE TABLE
  IF NOT EXISTS `user_passwords` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_passwords_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `user_refresh_tokens` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `value` CHAR(50) NOT NULL,
    `expires_at` timestamp NULL DEFAULT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_value` (`user_id`, `value`),
    CONSTRAINT `fk_user_refresh_tokens_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );