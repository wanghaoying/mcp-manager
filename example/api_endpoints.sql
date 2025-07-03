-- api_endpoints 表结构
CREATE TABLE IF NOT EXISTS `api_endpoints` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `swagger_id` BIGINT UNSIGNED DEFAULT NULL,
  `path` VARCHAR(255) NOT NULL,
  `method` VARCHAR(16) NOT NULL,
  `summary` VARCHAR(255) DEFAULT '',
  `description` TEXT DEFAULT NULL,
  `tags` VARCHAR(255) DEFAULT '',
  `parameters` JSON DEFAULT NULL,
  `responses` JSON DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_swagger_id` (`swagger_id`),
  KEY `idx_path_method` (`path`, `method`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
