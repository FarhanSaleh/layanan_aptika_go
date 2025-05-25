-- +migrate Up
CREATE TABLE `instansi` (
  `id` char(36) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `alamat` text,
  `keterangan` text,
  `is_deleted` tinyint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);
-- +migrate Down
DROP TABLE IF EXISTS `instansi`;