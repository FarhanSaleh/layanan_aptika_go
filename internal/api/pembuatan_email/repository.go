package pembuatanemail

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, pembuatanEmail *domain.PembuatanEmail) error
	Update(ctx context.Context, tx *sql.Tx, pembuatanEmail *domain.PembuatanEmail) error
	UpdateStatus(ctx context.Context, tx *sql.Tx, pembuatanEmail *domain.PembuatanEmail) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.PembuatanEmail, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PembuatanEmail, error)
	FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) ([]domain.PembuatanEmail, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pembuatanEmail *domain.PembuatanEmail) (err error) {
	SQL := `INSERT INTO pembuatan_email (
			id, 
			nama_lengkap, 
			nip, 
			jabatan, 
			nomor_hp, 
			berkas_sk, 
			surat_permohonan, 
			instansi_id,
			user_id) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, 
			pembuatanEmail.Id, pembuatanEmail.NamaLengkap, pembuatanEmail.NIP, 
			pembuatanEmail.Jabatan, pembuatanEmail.NomorHP, pembuatanEmail.BerkasSK, 
			pembuatanEmail.SuratPermohonan, pembuatanEmail.InstansiId, pembuatanEmail.UserId)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pembuatanEmail *domain.PembuatanEmail) (err error) {
	SQL := `UPDATE pembuatan_email 
		SET 
		nama_lengkap = ?, 
		nip = ?, 
		jabatan = ?, 
		nomor_hp = ?, 
		instansi_id = ?`

	args := []any{
		pembuatanEmail.NamaLengkap, 
		pembuatanEmail.NIP, 
		pembuatanEmail.Jabatan, 
		pembuatanEmail.NomorHP, 
		pembuatanEmail.InstansiId, 
	}

	if pembuatanEmail.BerkasSK != "" {
		SQL += `, berkas_sk = ?`
		args = append(args, pembuatanEmail.BerkasSK)
	}

	if pembuatanEmail.SuratPermohonan != "" {
		SQL += `, surat_permohonan = ?`
		args = append(args, pembuatanEmail.SuratPermohonan)
	}

	SQL += ` WHERE id = ?`
	args = append(args, pembuatanEmail.Id)

	_, err = tx.ExecContext(ctx, SQL, args...)
	return
}

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, pembuatanEmail *domain.PembuatanEmail) (err error) {
	SQL := `UPDATE pembuatan_email SET status = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, pembuatanEmail.Status, pembuatanEmail.Id)
	return 
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM pembuatan_email WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (result domain.PembuatanEmail, err error) {
	SQL := `SELECT 
			pe.id, 
			pe.nama_lengkap, 
			pe.nip, 
			pe.jabatan, 
			pe.nomor_hp, 
			pe.berkas_sk, 
			pe.surat_permohonan, 
			pe.status,
			pe.instansi_id,
			i.nama as nama_instansi,
			pe.created_at, 
			pe.updated_at
			FROM pembuatan_email as pe
			LEFT JOIN instansi as i ON pe.instansi_id = i.id 
			WHERE pe.id = ?`
	row := tx.QueryRowContext(ctx, SQL, id)
	err = row.Scan(
			&result.Id, 
			&result.NamaLengkap, 
			&result.NIP, 
			&result.Jabatan, 
			&result.NomorHP, 
			&result.BerkasSK,
			&result.SuratPermohonan,
			&result.Status,
			&result.InstansiId,
			&result.NamaInstansi,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.PembuatanEmail, err error) {
	SQL := `SELECT 
			pe.id, 
			pe.nama_lengkap, 
			pe.nip, 
			pe.jabatan, 
			pe.nomor_hp, 
			pe.berkas_sk, 
			pe.surat_permohonan, 
			pe.status,
			pe.instansi_id,
			i.nama as nama_instansi,
			pe.created_at 
			FROM pembuatan_email as pe
			LEFT JOIN instansi as i ON pe.instansi_id = i.id`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pe domain.PembuatanEmail
		err = rows.Scan(
			&pe.Id,
			&pe.NamaLengkap,
			&pe.NIP,
			&pe.Jabatan,
			&pe.NomorHP,
			&pe.BerkasSK,
			&pe.SuratPermohonan,
			&pe.Status,
			&pe.InstansiId,
			&pe.NamaInstansi,
			&pe.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, pe)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (r *RepositoryImpl) FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) (result []domain.PembuatanEmail, err error) {
	SQL := `SELECT 
			pe.id, 
			pe.nama_lengkap, 
			pe.nip, 
			pe.jabatan, 
			pe.nomor_hp, 
			pe.berkas_sk, 
			pe.surat_permohonan, 
			pe.status,
			pe.instansi_id,
			i.nama as nama_instansi,
			pe.created_at 
			FROM pembuatan_email as pe
			LEFT JOIN instansi as i ON pe.instansi_id = i.id
			WHERE pe.user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pe domain.PembuatanEmail
		err = rows.Scan(
			&pe.Id,
			&pe.NamaLengkap,
			&pe.NIP,
			&pe.Jabatan,
			&pe.NomorHP,
			&pe.BerkasSK,
			&pe.SuratPermohonan,
			&pe.Status,
			&pe.InstansiId,
			&pe.NamaInstansi,
			&pe.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, pe)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}
