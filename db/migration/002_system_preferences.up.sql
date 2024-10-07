CREATE TABLE
  IF NOT EXISTS `user_system_preferences` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `language` VARCHAR(35),
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_system_preferences_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );