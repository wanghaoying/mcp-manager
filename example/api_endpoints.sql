-- api_endpoints 表结构
CREATE TABLE IF NOT EXISTS `api_endpoints` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `swagger_id` BIGINT UNSIGNED DEFAULT NULL,
  `path` VARCHAR(255) NOT NULL,
  `method` VARCHAR(16) NOT NULL,
  `summary` VARCHAR(255) DEFAULT '',
  `description` TEXT DEFAULT NULL,
  `operation_id` VARCHAR(64) DEFAULT '', -- 长度调整为64，唯一标识 operationId
  `tags` VARCHAR(255) DEFAULT '',
  `parameters` text DEFAULT NULL,
  `responses` text DEFAULT NULL,
  `headers` text DEFAULT NULL, -- 新增
  `body` text DEFAULT NULL,    -- 新增
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_swagger_id` (`swagger_id`),
  KEY `idx_path_method` (`path`, `method`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API Endpoints Table';
