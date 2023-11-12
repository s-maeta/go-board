CREATE TABLE IF NOT EXISTS `articles` (
  `unique_id` varchar(255) NOT NULL,
  `user_unique_id` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  UNIQUE INDEX `unique_id_key`(`unique_id`),
  PRIMARY KEY (`unique_id`),
  FOREIGN KEY (`user_unique_id`) REFERENCES `users`(`unique_id`)
)