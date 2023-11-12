CREATE TABLE IF NOT EXISTS `users` (
  `unique_id` varchar(255) NOT NULL,
  `name` varchar(30) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  UNIQUE INDEX `users_email_key`(`email`),
  UNIQUE INDEX `users_unique_id_key`(`unique_id`),
  PRIMARY KEY (`unique_id`)
);