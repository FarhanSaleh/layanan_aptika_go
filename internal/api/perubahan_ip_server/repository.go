package perubahanipserver

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, perubahanIPServer *domain.PerubahanIPServer) error
	Update(ctx context.Context, tx *sql.Tx, perubahanIPServer *domain.PerubahanIPServer) error
	UpdateStatus(ctx context.Context, tx *sql.Tx, perubahanIPServer *domain.PerubahanIPServer) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.PerubahanIPServer, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PerubahanIPServer, error)
	FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) ([]domain.PerubahanIPServer, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, perubahanIPServer *domain.PerubahanIPServer) (err error) {
	SQL := `INSERT INTO perubahan_ip_server (
			id, 
			nama_lengkap, 
			jabatan, 
			nomor_hp, 
			nama_subdomain, 
			ip_lama, 
			ip_baru, 
			surat_permohonan, 
			instansi_id,
			user_id) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, 
			perubahanIPServer.Id, perubahanIPServer.NamaLengkap, perubahanIPServer.Jabatan, 
			perubahanIPServer.NomorHP, perubahanIPServer.NamaSubdomain, perubahanIPServer.IPLama, 
			perubahanIPServer.IPBaru, perubahanIPServer.SuratPermohonan, perubahanIPServer.InstansiId, perubahanIPServer.UserId)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, perubahanIPServer *domain.PerubahanIPServer) (err error) {
	SQL := `UPDATE perubahan_ip_server 
		SET 
		nama_lengkap = ?, 
		jabatan = ?, 
		nomor_hp = ?, 
		nama_subdomain = ?, 
		ip_lama = ?, 
		ip_baru = ?,
		instansi_id = ?`

	args := []any{
		perubahanIPServer.NamaLengkap, 
		perubahanIPServer.Jabatan, 
		perubahanIPServer.NomorHP, 
		perubahanIPServer.NamaSubdomain, 
		perubahanIPServer.IPLama, 
		perubahanIPServer.IPBaru,
		perubahanIPServer.InstansiId, 
	}

	if perubahanIPServer.SuratPermohonan != "" {
		SQL += `, surat_permohonan = ?`
		args = append(args, perubahanIPServer.SuratPermohonan)
	}

	SQL += ` WHERE id = ?`
	args = append(args, perubahanIPServer.Id)

	_, err = tx.ExecContext(ctx, SQL, args...)
	return
}

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, perubahanIPServer *domain.PerubahanIPServer) (err error) {
	SQL := `UPDATE perubahan_ip_server SET status = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, perubahanIPServer.Status, perubahanIPServer.Id)
	return 
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM perubahan_ip_server WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (result domain.PerubahanIPServer, err error) {
	SQL := `SELECT 
			pis.id, 
			pis.nama_lengkap, 
			pis.jabatan, 
			pis.nomor_hp, 
			pis.nama_subdomain, 
			pis.ip_lama, 
			pis.ip_baru, 
			pis.surat_permohonan, 
			pis.status,
			pis.instansi_id,
			i.nama as nama_instansi,
			pis.created_at, 
			pis.updated_at
			FROM perubahan_ip_server as pis
			LEFT JOIN instansi as i ON pis.instansi_id = i.id 
			WHERE pis.id = ?`
	row := tx.QueryRowContext(ctx, SQL, id)
	err = row.Scan(
			&result.Id, 
			&result.NamaLengkap, 
			&result.Jabatan, 
			&result.NomorHP, 
			&result.NamaSubdomain,
			&result.IPLama,
			&result.IPBaru,
			&result.SuratPermohonan,
			&result.Status,
			&result.InstansiId,
			&result.NamaInstansi,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.PerubahanIPServer, err error) {
	SQL := `SELECT 
			pis.id, 
			pis.nama_lengkap, 
			pis.jabatan, 
			pis.nomor_hp, 
			pis.nama_subdomain, 
			pis.ip_lama, 
			pis.ip_baru, 
			pis.surat_permohonan, 
			pis.status,
			pis.instansi_id,
			i.nama as nama_instansi,
			pis.created_at 
			FROM perubahan_ip_server as pis
			LEFT JOIN instansi as i ON pis.instansi_id = i.id
			ORDER BY pis.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var is domain.PerubahanIPServer
		err = rows.Scan(
			&is.Id,
			&is.NamaLengkap,
			&is.Jabatan,
			&is.NomorHP,
			&is.NamaSubdomain,
			&is.IPLama,
			&is.IPBaru,
			&is.SuratPermohonan,
			&is.Status,
			&is.InstansiId,
			&is.NamaInstansi,
			&is.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, is)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (r *RepositoryImpl) FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) (result []domain.PerubahanIPServer, err error) {
	SQL := `SELECT 
			pis.id, 
			pis.nama_lengkap, 
			pis.jabatan, 
			pis.nomor_hp, 
			pis.nama_subdomain, 
			pis.ip_lama, 
			pis.ip_baru, 
			pis.surat_permohonan, 
			pis.status,
			pis.instansi_id,
			i.nama as nama_instansi,
			pis.created_at 
			FROM perubahan_ip_server as pis
			LEFT JOIN instansi as i ON pis.instansi_id = i.id
			WHERE pis.user_id = ? ORDER BY pis.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var is domain.PerubahanIPServer
		err = rows.Scan(
			&is.Id,
			&is.NamaLengkap,
			&is.Jabatan,
			&is.NomorHP,
			&is.NamaSubdomain,
			&is.IPLama,
			&is.IPBaru,
			&is.SuratPermohonan,
			&is.Status,
			&is.InstansiId,
			&is.NamaInstansi,
			&is.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, is)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}
