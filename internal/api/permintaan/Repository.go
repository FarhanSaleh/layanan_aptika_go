package permintaan

import (
	"context"
	"database/sql"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	CountAll(ctx context.Context, tx *sql.Tx) (domain.PermintaanCountResponse, error)
	CountGangguanJIP(ctx context.Context, tx *sql.Tx) (domain.PermintaanCountResponse, error)
	CountPembuatanEmail(ctx context.Context, tx *sql.Tx) (domain.PermintaanCountResponse, error)
	CountPembuatanSubdomain(ctx context.Context, tx *sql.Tx) (domain.PermintaanCountResponse, error)
	CountPembangunanAplikasi(ctx context.Context, tx *sql.Tx) (domain.PermintaanCountResponse, error)
	CountPusatDataDaerah(ctx context.Context, tx *sql.Tx) (domain.PermintaanCountResponse, error)
	CountPerubahanIPServer(ctx context.Context, tx *sql.Tx) (domain.PermintaanCountResponse, error)
	CountAllByUser(ctx context.Context, tx *sql.Tx, uid string) (domain.PermintaanCountResponse, error)
	CountGangguanJIPByUser(ctx context.Context, tx *sql.Tx, uid string) (domain.PermintaanCountResponse, error)
	CountPembuatanEmailByUser(ctx context.Context, tx *sql.Tx, uid string) (domain.PermintaanCountResponse, error)
	CountPembuatanSubdomainByUser(ctx context.Context, tx *sql.Tx, uid string) (domain.PermintaanCountResponse, error)
	CountPembangunanAplikasiByUser(ctx context.Context, tx *sql.Tx, uid string) (domain.PermintaanCountResponse, error)
	CountPusatDataDaerahByUser(ctx context.Context, tx *sql.Tx, uid string) (domain.PermintaanCountResponse, error)
	CountPerubahanIPServerByUser(ctx context.Context, tx *sql.Tx, uid string) (domain.PermintaanCountResponse, error)
}

type RepositoryImpl struct{}

func NewRepository () Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) CountAll(ctx context.Context, tx *sql.Tx) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM (
			SELECT status FROM pengaduan_gangguan_jip
			UNION ALL
			SELECT status FROM pembuatan_email
			UNION ALL
			SELECT status FROM pembuatan_subdomain
			UNION ALL
			SELECT status FROM pembangunan_aplikasi
			UNION ALL
			SELECT status FROM perubahan_ip_server
			UNION ALL
			SELECT status FROM pusat_data_daerah
			) AS gabungan;`
	err = tx.QueryRowContext(ctx, SQL).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountGangguanJIP(ctx context.Context, tx *sql.Tx) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pengaduan_gangguan_jip;`
	err = tx.QueryRowContext(ctx, SQL).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPembuatanEmail(ctx context.Context, tx *sql.Tx) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pembuatan_email;`
	err = tx.QueryRowContext(ctx, SQL).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPembuatanSubdomain(ctx context.Context, tx *sql.Tx) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pembuatan_subdomain;`
	err = tx.QueryRowContext(ctx, SQL).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPembangunanAplikasi(ctx context.Context, tx *sql.Tx) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pembangunan_aplikasi;`
	err = tx.QueryRowContext(ctx, SQL).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPusatDataDaerah(ctx context.Context, tx *sql.Tx) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pusat_data_daerah;`
	err = tx.QueryRowContext(ctx, SQL).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPerubahanIPServer(ctx context.Context, tx *sql.Tx) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM perubahan_ip_server;`
	err = tx.QueryRowContext(ctx, SQL).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountAllByUser(ctx context.Context, tx *sql.Tx, uid string) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM (
			SELECT status, user_id FROM pengaduan_gangguan_jip
			UNION ALL
			SELECT status, user_id FROM pembuatan_email
			UNION ALL
			SELECT status, user_id FROM pembuatan_subdomain
			UNION ALL
			SELECT status, user_id FROM pembangunan_aplikasi
			UNION ALL
			SELECT status, user_id FROM perubahan_ip_server
			UNION ALL
			SELECT status, user_id FROM pusat_data_daerah
			) AS gabungan WHERE user_id = ?;`
	err = tx.QueryRowContext(ctx, SQL, uid).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountGangguanJIPByUser(ctx context.Context, tx *sql.Tx, uid string) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pengaduan_gangguan_jip WHERE user_id = ?;`
	err = tx.QueryRowContext(ctx, SQL, uid).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPembuatanEmailByUser(ctx context.Context, tx *sql.Tx, uid string) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pembuatan_email WHERE user_id = ?;`
	err = tx.QueryRowContext(ctx, SQL, uid).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPembuatanSubdomainByUser(ctx context.Context, tx *sql.Tx, uid string) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pembuatan_subdomain WHERE user_id = ?;`
	err = tx.QueryRowContext(ctx, SQL, uid).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPembangunanAplikasiByUser(ctx context.Context, tx *sql.Tx, uid string) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pembangunan_aplikasi WHERE user_id = ?;`
	err = tx.QueryRowContext(ctx, SQL, uid).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPusatDataDaerahByUser(ctx context.Context, tx *sql.Tx, uid string) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM pusat_data_daerah WHERE user_id = ?;`
	err = tx.QueryRowContext(ctx, SQL, uid).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}

func (r *RepositoryImpl) CountPerubahanIPServerByUser(ctx context.Context, tx *sql.Tx, uid string) (result domain.PermintaanCountResponse, err error) {
	SQL := `SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN status = 'diproses' THEN 1 ELSE 0 END), 0) AS diproses,
			COALESCE(SUM(CASE WHEN status = 'disetujui' THEN 1 ELSE 0 END), 0) AS disetujui,
			COALESCE(SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END), 0) AS ditolak
			FROM perubahan_ip_server WHERE user_id = ?;`
	err = tx.QueryRowContext(ctx, SQL, uid).Scan(&result.Total, &result.Diproses, &result.Disetujui, &result.Ditolak)
	return
}
