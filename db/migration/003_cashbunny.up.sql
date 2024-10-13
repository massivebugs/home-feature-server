CREATE TABLE
  IF NOT EXISTS `cashbunny_user_preferences` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_user_preferences_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_user_currencies` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `currency_code` CHAR(3) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_user_currencies_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_accounts` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `category` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL DEFAULT "",
    `currency` VARCHAR(255) NOT NULL,
    `order_index` SMALLINT (6) UNSIGNED NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_accounts_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    UNIQUE KEY `uk_cb_accounts_user_id_order_index` (`user_id`, `order_index`)
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_transaction_categories` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_transaction_categories_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_transaction_category_allocations` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `category_id` INT (10) unsigned NOT NULL,
    `amount` DECIMAL(19, 4) NOT NULL DEFAULT 0,
    `currency` VARCHAR(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_transaction_category_allocations_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_cb_transaction_category_allocations_category_id` FOREIGN KEY (`category_id`) REFERENCES `cashbunny_transaction_categories` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_scheduled_transactions` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `category_id` INT (10) unsigned,
    `src_account_id` INT (10) unsigned NOT NULL,
    `dest_account_id` INT (10) unsigned NOT NULL,
    `description` VARCHAR(255) NOT NULL DEFAULT "",
    `amount` DECIMAL(19, 4) NOT NULL DEFAULT 0,
    `currency` VARCHAR(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_scheduled_transactions_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_cb_scheduled_transactions_category_id` FOREIGN KEY (`category_id`) REFERENCES `cashbunny_transaction_categories` (`id`),
    CONSTRAINT `fk_cb_scheduled_transactions_src_account_id` FOREIGN KEY (`src_account_id`) REFERENCES `cashbunny_accounts` (`id`),
    CONSTRAINT `fk_cb_scheduled_transactions_dest_account_id` FOREIGN KEY (`dest_account_id`) REFERENCES `cashbunny_accounts` (`id`)
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_recurrence_rules` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `freq` VARCHAR(10) NOT NULL,
    `dtstart` datetime NOT NULL,
    `count` SMALLINT NOT NULL,
    `interval` SMALLINT NOT NULL,
    `until` datetime NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_scheduled_transactions_recurrence_rules` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `scheduled_transaction_id` INT (10) unsigned NOT NULL,
    `recurrence_rule_id` INT (10) unsigned NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_strr_scheduled_transaction_id` FOREIGN KEY (`scheduled_transaction_id`) REFERENCES `cashbunny_scheduled_transactions` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_cb_strr_recurrence_rule_id` FOREIGN KEY (`recurrence_rule_id`) REFERENCES `cashbunny_recurrence_rules` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_transactions` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `scheduled_transaction_id` INT (10) unsigned,
    `category_id` INT (10) unsigned,
    `src_account_id` INT (10) unsigned NOT NULL,
    `dest_account_id` INT (10) unsigned NOT NULL,
    `description` VARCHAR(255) NOT NULL DEFAULT "",
    `amount` DECIMAL(19, 4) NOT NULL DEFAULT 0,
    `currency` VARCHAR(255) NOT NULL,
    `transacted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_transactions_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_cb_transactions_scheduled_transaction_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT `fk_cb_transactions_category_id` FOREIGN KEY (`category_id`) REFERENCES `cashbunny_transaction_categories` (`id`),
    CONSTRAINT `fk_cb_transactions_src_account_id` FOREIGN KEY (`src_account_id`) REFERENCES `cashbunny_accounts` (`id`),
    CONSTRAINT `fk_cb_transactions_dest_account_id` FOREIGN KEY (`dest_account_id`) REFERENCES `cashbunny_accounts` (`id`)
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_tags` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `description` VARCHAR(255) NOT NULL DEFAULT "",
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_tags_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_accounts_tags` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `account_id` INT (10) unsigned NOT NULL,
    `tag_id` INT (10) unsigned NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_aat_account_id` FOREIGN KEY (`account_id`) REFERENCES `cashbunny_accounts` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_cb_aat_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `cashbunny_tags` (`id`) ON DELETE CASCADE
  );