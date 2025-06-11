-- +migrate Up
ALTER TABLE pembangunan_aplikasi MODIFY COLUMN status ENUM('diajukan', 'diterima', 'diproses', 'ditolak', 'disetujui') DEFAULT 'open';

-- +migrate Down
ALTER TABLE pembangunan_aplikasi MODIFY COLUMN status ENUM('diproses', 'ditolak', 'disetujui') DEFAULT 'diproses';
