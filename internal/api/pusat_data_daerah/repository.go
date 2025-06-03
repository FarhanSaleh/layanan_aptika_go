package pusatdatadaerah

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, pusatDataDaerah *domain.PusatDataDaerah) error
	Update(ctx context.Context, tx *sql.Tx, pusatDataDaerah *domain.PusatDataDaerah) error
	UpdateStatus(ctx context.Context, tx *sql.Tx, pusatDataDaerah *domain.PusatDataDaerah) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.PusatDataDaerah, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PusatDataDaerah, error)
	FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) ([]domain.PusatDataDaerah, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pusatDataDaerah *domain.PusatDataDaerah) (err error) {
	SQL := `INSERT INTO pusat_data_daerah (
			id, 
			nama_lengkap, 
			jabatan, 
			nomor_hp, 
			jenis_layanan, 
			surat_permohonan, 
			instansi_id,
			user_id) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, 
			pusatDataDaerah.Id, pusatDataDaerah.NamaLengkap, pusatDataDaerah.Jabatan, pusatDataDaerah.NomorHP, 
			pusatDataDaerah.JenisLayanan, pusatDataDaerah.SuratPermohonan, pusatDataDaerah.InstansiId, pusatDataDaerah.UserId)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pusatDataDaerah *domain.PusatDataDaerah) (err error) {
	SQL := `UPDATE pusat_data_daerah 
		SET 
		nama_lengkap = ?, 
		jabatan = ?, 
		nomor_hp = ?, 
		jenis_layanan = ?,
		instansi_id = ?`

	args := []any{
		pusatDataDaerah.NamaLengkap, 
		pusatDataDaerah.Jabatan, 
		pusatDataDaerah.NomorHP, 
		pusatDataDaerah.JenisLayanan,
		pusatDataDaerah.InstansiId, 
	}

	if pusatDataDaerah.SuratPermohonan != "" {
		SQL += `, surat_permohonan = ?`
		args = append(args, pusatDataDaerah.SuratPermohonan)
	}

	SQL += ` WHERE id = ?`
	args = append(args, pusatDataDaerah.Id)

	_, err = tx.ExecContext(ctx, SQL, args...)
	return
}

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, pusatDataDaerah *domain.PusatDataDaerah) (err error) {
	SQL := `UPDATE pusat_data_daerah SET status = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, pusatDataDaerah.Status, pusatDataDaerah.Id)
	return 
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM pusat_data_daerah WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (result domain.PusatDataDaerah, err error) {
	SQL := `SELECT 
			pdd.id, 
			pdd.nama_lengkap, 
			pdd.jabatan, 
			pdd.nomor_hp, 
			pdd.jenis_layanan, 
			pdd.surat_permohonan, 
			pdd.status,
			pdd.instansi_id,
			i.nama as nama_instansi,
			u.notification_token,
			pdd.created_at, 
			pdd.updated_at
			FROM pusat_data_daerah as pdd
			LEFT JOIN instansi as i ON pdd.instansi_id = i.id 
			LEFT JOIN users as u ON pdd.user_id = u.id 
			WHERE pdd.id = ?`
	row := tx.QueryRowContext(ctx, SQL, id)
	err = row.Scan(
			&result.Id, 
			&result.NamaLengkap, 
			&result.Jabatan, 
			&result.NomorHP, 
			&result.JenisLayanan,
			&result.SuratPermohonan,
			&result.Status,
			&result.InstansiId,
			&result.NamaInstansi,
			&result.NotificationToken,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.PusatDataDaerah, err error) {
	SQL := `SELECT 
			pdd.id, 
			pdd.nama_lengkap, 
			pdd.jabatan, 
			pdd.nomor_hp, 
			pdd.jenis_layanan, 
			pdd.surat_permohonan, 
			pdd.status,
			pdd.instansi_id,
			i.nama as nama_instansi,
			pdd.created_at 
			FROM pusat_data_daerah as pdd
			LEFT JOIN instansi as i ON pdd.instansi_id = i.id
			ORDER BY pdd.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pdd domain.PusatDataDaerah
		err = rows.Scan(
			&pdd.Id,
			&pdd.NamaLengkap,
			&pdd.Jabatan,
			&pdd.NomorHP,
			&pdd.JenisLayanan,
			&pdd.SuratPermohonan,
			&pdd.Status,
			&pdd.InstansiId,
			&pdd.NamaInstansi,
			&pdd.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, pdd)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (r *RepositoryImpl) FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) (result []domain.PusatDataDaerah, err error) {
	SQL := `SELECT 
			pdd.id, 
			pdd.nama_lengkap, 
			pdd.jabatan, 
			pdd.nomor_hp, 
			pdd.jenis_layanan, 
			pdd.surat_permohonan, 
			pdd.status,
			pdd.instansi_id,
			i.nama as nama_instansi,
			pdd.created_at 
			FROM pusat_data_daerah as pdd
			LEFT JOIN instansi as i ON pdd.instansi_id = i.id
			WHERE pdd.user_id = ? ORDER BY pdd.created_at DESC`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pdd domain.PusatDataDaerah
		err = rows.Scan(
			&pdd.Id,
			&pdd.NamaLengkap,
			&pdd.Jabatan,
			&pdd.NomorHP,
			&pdd.JenisLayanan,
			&pdd.SuratPermohonan,
			&pdd.Status,
			&pdd.InstansiId,
			&pdd.NamaInstansi,
			&pdd.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, pdd)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}
