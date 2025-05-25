-- +migrate Up
CREATE TABLE `role_pengelola` (
  `id` char(36) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `is_deleted` tinyint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `role_pengelola`;