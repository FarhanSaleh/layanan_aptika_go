package gangguanjip

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, gangguanJIP *domain.GangguanJIP) error
	Update(ctx context.Context, tx *sql.Tx, gangguanJIP *domain.GangguanJIP) error
	UpdateStatus(ctx context.Context, tx *sql.Tx, gangguanJIP *domain.GangguanJIP) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.GangguanJIP, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.GangguanJIP, error)
	FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) ([]domain.GangguanJIP, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, gangguanJIP *domain.GangguanJIP) (err error) {
	SQL := `INSERT INTO pengaduan_gangguan_jip (
			id, 
			nama_lengkap, 
			jabatan, 
			nomor_hp, 
			lokasi_gangguan, 
			deskripsi_gangguan, 
			surat_permohonan, 
			foto, 
			instansi_id,
			user_id) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, 
			gangguanJIP.Id, gangguanJIP.NamaLengkap, gangguanJIP.Jabatan, 
			gangguanJIP.NomorHP, gangguanJIP.LokasiGangguan, gangguanJIP.DeskripsiGangguan, 
			gangguanJIP.SuratPermohonan, gangguanJIP.Foto, gangguanJIP.InstansiId, gangguanJIP.UserId)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, gangguanJIP *domain.GangguanJIP) (err error) {
	SQL := `UPDATE pengaduan_gangguan_jip 
		SET 
		nama_lengkap = ?, 
		jabatan = ?, 
		nomor_hp = ?, 
		lokasi_gangguan = ?, 
		deskripsi_gangguan = ?, 
		instansi_id = ?`

	args := []any{
		gangguanJIP.NamaLengkap, 
		gangguanJIP.Jabatan, 
		gangguanJIP.NomorHP, 
		gangguanJIP.LokasiGangguan, 
		gangguanJIP.DeskripsiGangguan, 
		gangguanJIP.InstansiId, 
	}

	if gangguanJIP.SuratPermohonan != "" {
		SQL += `, surat_permohonan = ?`
		args = append(args, gangguanJIP.SuratPermohonan)
	}

	if gangguanJIP.Foto != "" {
		SQL += `, foto = ?`
		args = append(args, gangguanJIP.Foto)
	}

	SQL += ` WHERE id = ?`
	args = append(args, gangguanJIP.Id)

	_, err = tx.ExecContext(ctx, SQL, args...)
	return
}

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, gangguanJIP *domain.GangguanJIP) (err error) {
	SQL := `UPDATE pengaduan_gangguan_jip SET status = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, gangguanJIP.Status, gangguanJIP.Id)
	return 
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM pengaduan_gangguan_jip WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (result domain.GangguanJIP, err error) {
	SQL := `SELECT 
			gj.id, 
			gj.nama_lengkap, 
			gj.jabatan, 
			gj.nomor_hp, 
			gj.lokasi_gangguan, 
			gj.deskripsi_gangguan, 
			gj.surat_permohonan, 
			gj.foto, 
			gj.status,
			gj.instansi_id,
			i.nama as nama_instansi,
			gj.created_at, 
			gj.updated_at
			FROM pengaduan_gangguan_jip as gj
			LEFT JOIN instansi as i ON gj.instansi_id = i.id 
			WHERE gj.id = ?`
	row := tx.QueryRowContext(ctx, SQL, id)
	err = row.Scan(
			&result.Id, 
			&result.NamaLengkap, 
			&result.Jabatan, 
			&result.NomorHP, 
			&result.LokasiGangguan,
			&result.DeskripsiGangguan,
			&result.SuratPermohonan,
			&result.Foto,
			&result.Status,
			&result.InstansiId,
			&result.NamaInstansi,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.GangguanJIP, err error) {
	SQL := `SELECT 
			gj.id, 
			gj.nama_lengkap, 
			gj.jabatan, 
			gj.nomor_hp, 
			gj.lokasi_gangguan, 
			gj.surat_permohonan, 
			gj.status,
			gj.instansi_id,
			i.nama as nama_instansi,
			gj.created_at 
			FROM pengaduan_gangguan_jip as gj
			LEFT JOIN instansi as i ON gj.instansi_id = i.id`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var gj domain.GangguanJIP
		err = rows.Scan(
			&gj.Id,
			&gj.NamaLengkap,
			&gj.Jabatan,
			&gj.NomorHP,
			&gj.LokasiGangguan,
			&gj.SuratPermohonan,
			&gj.Status,
			&gj.InstansiId,
			&gj.NamaInstansi,
			&gj.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, gj)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (r *RepositoryImpl) FindAllByUser(ctx context.Context, tx *sql.Tx, userId string) (result []domain.GangguanJIP, err error) {
	SQL := `SELECT 
			gj.id, 
			gj.nama_lengkap, 
			gj.jabatan, 
			gj.nomor_hp, 
			gj.lokasi_gangguan, 
			gj.surat_permohonan, 
			gj.status,
			gj.instansi_id,
			i.nama as nama_instansi,
			gj.created_at 
			FROM pengaduan_gangguan_jip as gj
			LEFT JOIN instansi as i ON gj.instansi_id = i.id
			WHERE gj.user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var gj domain.GangguanJIP
		err = rows.Scan(
			&gj.Id,
			&gj.NamaLengkap,
			&gj.Jabatan,
			&gj.NomorHP,
			&gj.LokasiGangguan,
			&gj.SuratPermohonan,
			&gj.Status,
			&gj.InstansiId,
			&gj.NamaInstansi,
			&gj.CreatedAt,
		)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		result = append(result, gj)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}
	
	return
}
