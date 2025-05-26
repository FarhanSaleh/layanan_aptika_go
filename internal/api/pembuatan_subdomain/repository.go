package pembuatansubdomain

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, pembuatanSubdomain *domain.PembuatanSubdomain) error
	Update(ctx context.Context, tx *sql.Tx, pembuatanSubdomain *domain.PembuatanSubdomain) error
	UpdateStatus(ctx context.Context, tx *sql.Tx, pembuatanSubdomain *domain.PembuatanSubdomain) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.PembuatanSubdomain, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PembuatanSubdomain, error)
	FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) ([]domain.PembuatanSubdomain, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pembuatanSubdomain *domain.PembuatanSubdomain) (err error) {
	SQL := `INSERT INTO pembuatan_subdomain (
			id, 
			nama_lengkap, 
			jabatan, 
			nomor_hp, 
			nama_subdomain, 
			ip_publik,
			deskripsi, 
			surat_permohonan, 
			instansi_id,
			user_id) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, 
			pembuatanSubdomain.Id, pembuatanSubdomain.NamaLengkap, pembuatanSubdomain.Jabatan, pembuatanSubdomain.NomorHP, 
			pembuatanSubdomain.NamaSubdomain, pembuatanSubdomain.IPPublik, pembuatanSubdomain.Deskripsi,
			pembuatanSubdomain.SuratPermohonan, pembuatanSubdomain.InstansiId, pembuatanSubdomain.UserId)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pembuatanSubdomain *domain.PembuatanSubdomain) (err error) {
	SQL := `UPDATE pembuatan_subdomain 
		SET 
		nama_lengkap = ?, 
		jabatan = ?, 
		nomor_hp = ?, 
		nama_subdomain = ?,
		ip_publik = ?,
		deskripsi = ?,
		instansi_id = ?`

	args := []any{
		pembuatanSubdomain.NamaLengkap, 
		pembuatanSubdomain.Jabatan, 
		pembuatanSubdomain.NomorHP, 
		pembuatanSubdomain.NamaSubdomain,
		pembuatanSubdomain.IPPublik, 
		pembuatanSubdomain.Deskripsi, 
		pembuatanSubdomain.InstansiId, 
	}

	if pembuatanSubdomain.SuratPermohonan != "" {
		SQL += `, surat_permohonan = ?`
		args = append(args, pembuatanSubdomain.SuratPermohonan)
	}

	SQL += ` WHERE id = ?`
	args = append(args, pembuatanSubdomain.Id)

	_, err = tx.ExecContext(ctx, SQL, args...)
	return
}

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, pembuatanSubdomain *domain.PembuatanSubdomain) (err error) {
	SQL := `UPDATE pembuatan_subdomain SET status = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, pembuatanSubdomain.Status, pembuatanSubdomain.Id)
	return 
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM pembuatan_subdomain WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (result domain.PembuatanSubdomain, err error) {
	SQL := `SELECT 
			ps.id, 
			ps.nama_lengkap, 
			ps.jabatan, 
			ps.nomor_hp, 
			ps.nama_subdomain, 
			ps.ip_publik, 
			ps.deskripsi, 
			ps.surat_permohonan, 
			ps.status,
			ps.instansi_id,
			i.nama as nama_instansi,
			ps.created_at, 
			ps.updated_at
			FROM pembuatan_subdomain as ps
			LEFT JOIN instansi as i ON ps.instansi_id = i.id 
			WHERE ps.id = ?`
	row := tx.QueryRowContext(ctx, SQL, id)
	err = row.Scan(
			&result.Id, 
			&result.NamaLengkap, 
			&result.Jabatan, 
			&result.NomorHP, 
			&result.NamaSubdomain,
			&result.IPPublik,
			&result.Deskripsi,
			&result.SuratPermohonan,
			&result.Status,
			&result.InstansiId,
			&result.NamaInstansi,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.PembuatanSubdomain, err error) {
	SQL := `SELECT 
			ps.id, 
			ps.nama_lengkap, 
			ps.jabatan, 
			ps.nomor_hp, 
			ps.nama_subdomain, 
			ps.ip_publik, 
			ps.surat_permohonan, 
			ps.status,
			ps.instansi_id,
			i.nama as nama_instansi,
			ps.created_at 
			FROM pembuatan_subdomain as ps
			LEFT JOIN instansi as i ON ps.instansi_id = i.id
			ORDER BY ps.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ps domain.PembuatanSubdomain
		err = rows.Scan(
			&ps.Id,
			&ps.NamaLengkap,
			&ps.Jabatan,
			&ps.NomorHP,
			&ps.NamaSubdomain,
			&ps.IPPublik,
			&ps.SuratPermohonan,
			&ps.Status,
			&ps.InstansiId,
			&ps.NamaInstansi,
			&ps.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, ps)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (r *RepositoryImpl) FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) (result []domain.PembuatanSubdomain, err error) {
	SQL := `SELECT 
			ps.id, 
			ps.nama_lengkap, 
			ps.jabatan, 
			ps.nomor_hp, 
			ps.nama_subdomain, 
			ps.ip_publik, 
			ps.surat_permohonan, 
			ps.status,
			ps.instansi_id,
			i.nama as nama_instansi,
			ps.created_at
			FROM pembuatan_subdomain as ps
			LEFT JOIN instansi as i ON ps.instansi_id = i.id
			WHERE ps.user_id = ? ORDER BY ps.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ps domain.PembuatanSubdomain
		err = rows.Scan(
			&ps.Id,
			&ps.NamaLengkap,
			&ps.Jabatan,
			&ps.NomorHP,
			&ps.NamaSubdomain,
			&ps.IPPublik,
			&ps.SuratPermohonan,
			&ps.Status,
			&ps.InstansiId,
			&ps.NamaInstansi,
			&ps.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, ps)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}
