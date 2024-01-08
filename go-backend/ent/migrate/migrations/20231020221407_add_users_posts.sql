-- Create "users" table
CREATE TABLE `users` (`id` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `email_addr` varchar(255) NOT NULL, `favorites` json NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
