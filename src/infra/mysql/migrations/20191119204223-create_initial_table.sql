
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `login_at` timestamp NULL DEFAULT NULL,
  `is_verified` tinyint(1) DEFAULT NULL,
  `verification_token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `verified_at` timestamp NULL DEFAULT NULL,
  `handphone` varchar(15) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=52898 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `agents` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_agents_id` (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `countries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `sortname` varchar(10) NOT NULL DEFAULT '',
  `name` varchar(200) NOT NULL DEFAULT '',
  `phonecode` VARCHAR(10) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `states` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT '',
  `country_id` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_provinces_country_id` (`country_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `cities` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT '',
  `state_id` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_cities_state_id` (`state_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `locations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT '',
  `country_id` int(10) NOT NULL DEFAULT '0',
  `state_id` int(10) NOT NULL DEFAULT '0',
  `cities_id` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_locations_country_id` (`country_id`),
  KEY `idx_locations_state_id` (`state_id`),
  KEY `idx_locations_cities_id` (`cities_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `agents_trips` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `agent_id` int(10) NOT NULL DEFAULT '0',
  `country_id` int(10) NOT NULL DEFAULT '0',
  `state_id` int(10) NOT NULL DEFAULT '0',
  `cities_id` int(10) NOT NULL DEFAULT '0',
  `location_id` int(10) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_trips_agents_id` (`id`),
  KEY `idx_trips_name` (`name`),
  KEY `idx_trips_agent_id` (`agent_id`),
  KEY `idx_trips_country_id` (`country_id`),
  KEY `idx_trips_state_id` (`state_id`),
  KEY `idx_trips_cities_id` (`cities_id`),
  KEY `idx_trips_location_id` (`location_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- +migrate Down
DROP TABLE `users`;
DROP TABLE `agents`;
DROP TABLE `countries`;
DROP TABLE `provinces`;
DROP TABLE `cities`;
DROP TABLE `trips`;
