package rolepengelola

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, rolePengelola *domain.RolePengelola) error
	Update(ctx context.Context, tx *sql.Tx, rolePengelola *domain.RolePengelola) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.RolePengelola, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.RolePengelola, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, rolePengelola *domain.RolePengelola) (err error) {
	SQL := `INSERT INTO role_pengelola (id, nama) VALUES (?, ?)`
	_, err = tx.ExecContext(ctx, SQL, rolePengelola.Id, rolePengelola.Nama)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, rolePengelola *domain.RolePengelola) (err error) {
	SQL := `UPDATE role_pengelola SET nama = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, rolePengelola.Nama, rolePengelola.Id)
	return
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM role_pengelola WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (rolePengelolaResult domain.RolePengelola, err error) {
	SQL := `SELECT id, nama FROM role_pengelola WHERE id = ?`
	err = tx.QueryRowContext(ctx, SQL, id).Scan(&rolePengelolaResult.Id, &rolePengelolaResult.Nama)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("role pengelola tidak ditemukan")
			return
		}
		log.Println("ERROR QUERY: ", err)
		return
	}
	return
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (rolePengelolaResult []domain.RolePengelola, err error) {
	SQL := `SELECT id, nama FROM role_pengelola`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rp domain.RolePengelola
		err = rows.Scan(&rp.Id, &rp.Nama)
		if err != nil {
			log.Println("ERROR SCANNING: ", err)
			return
		}
		rolePengelolaResult = append(rolePengelolaResult, rp)
	}
	
	return
}