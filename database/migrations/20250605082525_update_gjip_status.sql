-- +migrate Up
ALTER TABLE pengaduan_gangguan_jip MODIFY COLUMN status ENUM('diajukan', 'disetujui', 'diproses', 'ditolak', 'ditindaklanjuti teknisi', 'selesai') DEFAULT 'open';

-- +migrate Down
ALTER TABLE pengaduan_gangguan_jip MODIFY COLUMN status ENUM('diproses', 'ditolak', 'disetujui') DEFAULT 'diproses';
