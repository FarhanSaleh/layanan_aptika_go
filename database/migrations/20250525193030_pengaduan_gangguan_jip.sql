-- +migrate Up
CREATE TABLE IF NOT EXISTS `pengaduan_gangguan_jip` (
  `id` char(36) NOT NULL,
  `nama_lengkap` varchar(255) NOT NULL,
  `jabatan` varchar(255) NOT NULL,
  `nomor_hp` char(13) NOT NULL,
  `lokasi_gangguan` text NOT NULL,
  `deskripsi_gangguan` text,
  `foto` text,
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
  CONSTRAINT `pengaduan_gangguan_jip_ibfk_1` FOREIGN KEY (`instansi_id`) REFERENCES `instansi` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `pengaduan_gangguan_jip_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS `pengaduan_gangguan_jip`;