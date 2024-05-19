CREATE TABLE `todos` (
  `id` int AUTO_INCREMENT PRIMARY KEY,
  `user_id` int,
  `title` VARCHAR(255) NOT NULL,
  `body` VARCHAR(255) NOT NULL,
  `priority` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `due_date` TIMESTAMP NULL
);

CREATE TABLE `users` (
  `id` int AUTO_INCREMENT PRIMARY KEY,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `todos` COMMENT = 'Stores user todos';

ALTER TABLE `users` COMMENT = 'Stores users details';

ALTER TABLE `todos` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);