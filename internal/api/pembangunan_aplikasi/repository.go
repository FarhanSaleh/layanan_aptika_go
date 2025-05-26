package pembangunanaplikasi

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, pembangunanAplikasi *domain.PembangunanAplikasi) error
	Update(ctx context.Context, tx *sql.Tx, pembangunanAplikasi *domain.PembangunanAplikasi) error
	UpdateStatus(ctx context.Context, tx *sql.Tx, pembangunanAplikasi *domain.PembangunanAplikasi) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.PembangunanAplikasi, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PembangunanAplikasi, error)
	FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) ([]domain.PembangunanAplikasi, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pembangunanAplikasi *domain.PembangunanAplikasi) (err error) {
	SQL := `INSERT INTO pembangunan_aplikasi (
			id, 
			nama_pimpinan, 
			nomor_hp, 
			email_dinas, 
			riwayat_pimpinan, 
			jenis_aplikasi, 
			tujuan_aplikasi, 
			surat_permohonan, 
			instansi_id,
			user_id) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, 
			pembangunanAplikasi.Id, pembangunanAplikasi.NamaPimpinan, pembangunanAplikasi.NomorHP, 
			pembangunanAplikasi.EmailDinas,	pembangunanAplikasi.RiwayatPimpinan, pembangunanAplikasi.JenisAplikasi,
			pembangunanAplikasi.TujuanAplikasi, pembangunanAplikasi.SuratPermohonan, pembangunanAplikasi.InstansiId,
			pembangunanAplikasi.UserId)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pembangunanAplikasi *domain.PembangunanAplikasi) (err error) {
	SQL := `UPDATE pembangunan_aplikasi 
		SET 
		nama_pimpinan = ?, 
		nomor_hp = ?, 
		email_dinas = ?, 
		riwayat_pimpinan = ?,
		jenis_aplikasi = ?,
		tujuan_aplikasi = ?,
		instansi_id = ?`

	args := []any{
		pembangunanAplikasi.NamaPimpinan, 
		pembangunanAplikasi.NomorHP, 
		pembangunanAplikasi.EmailDinas, 
		pembangunanAplikasi.RiwayatPimpinan, 
		pembangunanAplikasi.JenisAplikasi, 
		pembangunanAplikasi.TujuanAplikasi, 
		pembangunanAplikasi.InstansiId, 
	}

	if pembangunanAplikasi.SuratPermohonan != "" {
		SQL += `, surat_permohonan = ?`
		args = append(args, pembangunanAplikasi.SuratPermohonan)
	}

	SQL += ` WHERE id = ?`
	args = append(args, pembangunanAplikasi.Id)

	_, err = tx.ExecContext(ctx, SQL, args...)
	return
}

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, pembangunanAplikasi *domain.PembangunanAplikasi) (err error) {
	SQL := `UPDATE pembangunan_aplikasi SET status = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, pembangunanAplikasi.Status, pembangunanAplikasi.Id)
	return 
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM pembangunan_aplikasi WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (result domain.PembangunanAplikasi, err error) {
	SQL := `SELECT 
			pa.id, 
			pa.nama_pimpinan, 
			pa.nomor_hp, 
			pa.email_dinas, 
			pa.riwayat_pimpinan, 
			pa.jenis_aplikasi, 
			pa.tujuan_aplikasi, 
			pa.surat_permohonan, 
			pa.status,
			pa.instansi_id,
			i.nama as nama_instansi,
			pa.created_at, 
			pa.updated_at
			FROM pembangunan_aplikasi as pa
			LEFT JOIN instansi as i ON pa.instansi_id = i.id 
			WHERE pa.id = ?`
	row := tx.QueryRowContext(ctx, SQL, id)
	err = row.Scan(
			&result.Id, 
			&result.NamaPimpinan, 
			&result.NomorHP, 
			&result.EmailDinas, 
			&result.RiwayatPimpinan,
			&result.JenisAplikasi,
			&result.TujuanAplikasi,
			&result.SuratPermohonan,
			&result.Status,
			&result.InstansiId,
			&result.NamaInstansi,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.PembangunanAplikasi, err error) {
	SQL := `SELECT 
			pa.id, 
			pa.nama_pimpinan, 
			pa.nomor_hp, 
			pa.email_dinas, 
			pa.jenis_aplikasi, 
			pa.surat_permohonan, 
			pa.status,
			pa.instansi_id,
			i.nama as nama_instansi,
			pa.created_at 
			FROM pembangunan_aplikasi as pa
			LEFT JOIN instansi as i ON pa.instansi_id = i.id
			ORDER BY pa.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pa domain.PembangunanAplikasi
		err = rows.Scan(
			&pa.Id,
			&pa.NamaPimpinan,
			&pa.NomorHP,
			&pa.EmailDinas,
			&pa.JenisAplikasi,
			&pa.SuratPermohonan,
			&pa.Status,
			&pa.InstansiId,
			&pa.NamaInstansi,
			&pa.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, pa)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (r *RepositoryImpl) FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) (result []domain.PembangunanAplikasi, err error) {
	SQL := `SELECT 
			pa.id, 
			pa.nama_pimpinan, 
			pa.nomor_hp, 
			pa.email_dinas, 
			pa.jenis_aplikasi, 
			pa.surat_permohonan, 
			pa.status,
			pa.instansi_id,
			i.nama as nama_instansi,
			pa.created_at 
			FROM pembangunan_aplikasi as pa
			LEFT JOIN instansi as i ON pa.instansi_id = i.id
			WHERE pa.user_id = ? ORDER BY pa.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pa domain.PembangunanAplikasi
		err = rows.Scan(
			&pa.Id,
			&pa.NamaPimpinan,
			&pa.NomorHP,
			&pa.EmailDinas,
			&pa.JenisAplikasi,
			&pa.SuratPermohonan,
			&pa.Status,
			&pa.InstansiId,
			&pa.NamaInstansi,
			&pa.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, pa)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}
