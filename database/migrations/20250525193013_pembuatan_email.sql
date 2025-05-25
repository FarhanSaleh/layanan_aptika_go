-- +migrate Up
CREATE TABLE `pembuatan_email` (
  `id` char(36) NOT NULL,
  `nama_lengkap` varchar(255) NOT NULL,
  `nip` char(18) NOT NULL,
  `jabatan` varchar(255) NOT NULL,
  `nomor_hp` char(13) NOT NULL,
  `berkas_sk` text NOT NULL,
  `surat_permohonan` text,
  `status` enum('diproses','disetujui','ditolak') DEFAULT 'diproses',
  `is_deleted` tinyint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `instansi_id` char(36) DEFAULT NULL,
  `user_id` char(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `instansi_id` (`instansi_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `pembuatan_email_ibfk_1` FOREIGN KEY (`instansi_id`) REFERENCES `instansi` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `pembuatan_email_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS `pembuatan_email`;