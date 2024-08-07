CREATE TABLE
  IF NOT EXISTS `cashbunny_accounts` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `category` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL DEFAULT "",
    `balance` DECIMAL(19, 4) NOT NULL DEFAULT 0,
    `currency` VARCHAR(255) NOT NULL,
    `type` VARCHAR(255) NOT NULL,
    `order_index` SMALLINT (6) UNSIGNED NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_accounts_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_transactions` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `description` VARCHAR(255) NOT NULL DEFAULT "",
    `transacted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_transactions_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS `cashbunny_entries` (
    `id` INT (10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` INT (10) unsigned NOT NULL,
    `account_id` INT (10) unsigned NOT NULL,
    `transaction_id` INT (10) unsigned NOT NULL,
    `amount` DECIMAL(19, 4) NOT NULL DEFAULT 0,
    `currency` VARCHAR(255) NOT NULL,
    `transacted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_cb_entries_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_cb_entries_account_id_cashbunny_accounts_id` FOREIGN KEY (`account_id`) REFERENCES `cashbunny_accounts` (`id`),
    CONSTRAINT `fk_cb_entries_transaction_id_cashbunny_transactions_id` FOREIGN KEY (`transaction_id`) REFERENCES `cashbunny_transactions` (`id`) ON DELETE CASCADE
  );