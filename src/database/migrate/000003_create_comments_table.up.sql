CREATE TABLE IF NOT EXISTS `comments` (
  `unique_id` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `article_unique_id` varchar(255) NOT NULL,
  `user_unique_id` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  PRIMARY KEY (`unique_id`),
  FOREIGN KEY (`article_unique_id`)  REFERENCES `articles`(`unique_id`),
  FOREIGN KEY (`user_unique_id`)  REFERENCES `users`(`unique_id`)
)