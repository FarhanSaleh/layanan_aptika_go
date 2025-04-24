package instansi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, instansi *domain.Instansi) error
	Update(ctx context.Context, tx *sql.Tx, instansi *domain.Instansi) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Instansi, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Instansi, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, instansi *domain.Instansi) (err error) {
	SQL := `INSERT INTO instansi (id, nama, alamat, keterangan) VALUES (?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, instansi.Id, instansi.Nama, instansi.Alamat, instansi.Keterangan)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, instansi *domain.Instansi) (err error) {
	SQL := `UPDATE instansi SET nama = ?, alamat = ?, keterangan = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, instansi.Nama, instansi.Alamat, instansi.Keterangan, instansi.Id)
	return
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM instansi WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (instansiResult domain.Instansi, err error) {
	SQL := `SELECT id, nama, alamat, keterangan FROM instansi WHERE id = ?`
	err = tx.QueryRowContext(ctx, SQL, id).Scan(&instansiResult.Id, &instansiResult.Nama, &instansiResult.Alamat, &instansiResult.Keterangan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("instansi tidak ditemukan")
			return
		}
		log.Println("ERROR QUERY: ", err)
		return
	}
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (instansiResult []domain.Instansi, err error) {
	SQL := `SELECT id, nama, alamat, keterangan FROM instansi`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.Instansi
		err = rows.Scan(&i.Id, &i.Nama, &i.Alamat, &i.Keterangan)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		instansiResult = append(instansiResult, i)
	}
	
	return
}