package pengelola

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, pengelola *domain.Pengelola) error
	Update(ctx context.Context, tx *sql.Tx, pengelola *domain.Pengelola) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Pengelola, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pengelola, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Pengelola, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, pengelola *domain.Pengelola) error
}

type RepositoryImpl struct{}

func NewRepository() Repository{
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pengelola *domain.Pengelola) (err error) {
	SQL := `INSERT INTO pengelola (id, nama, email, password, role_id) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, pengelola.Id, pengelola.Nama, pengelola.Email, pengelola.Password, pengelola.RoleId)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pengelola *domain.Pengelola) (err error) {
	SQL := `UPDATE pengelola SET nama = ?, email = ?, role_id = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, pengelola.Nama, pengelola.Email, pengelola.RoleId, pengelola.Id)
	return
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM pengelola WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (result domain.Pengelola, err error) {
	SQL := `SELECT 
			p.id, 
			p.nama, 
			p.email, 
			p.role_id,
			r.nama as nama_role,
			p.created_at
			FROM 
			pengelola as p
			LEFT JOIN role_pengelola as r ON p.role_id = r.id 
			WHERE 
			p.id = ?`
	err = tx.QueryRowContext(ctx, SQL, id).Scan(&result.Id, &result.Nama, &result.Email, &result.RoleId, &result.NamaRole, &result.CreatedAt)
	return 
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.Pengelola, err error) {
	SQL := `SELECT 
			p.id, 
			p.nama, 
			p.email,
			p.role_id,
			r.nama as nama_role
			FROM pengelola as p
			LEFT JOIN role_pengelola as r ON p.role_id = r.id`

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u domain.Pengelola
		err = rows.Scan(&u.Id, &u.Nama, &u.Email, &u.RoleId, &u.NamaRole)
		if err != nil{
			log.Println("ERROR SCANNING: ", err)
			return 
		}
		result = append(result, u)
	}
	if result == nil {
		err = sql.ErrNoRows
		return
	}

	return	
}

func (r *RepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (result domain.Pengelola, err error) {
	SQL := `SELECT 
			p.id, 
			p.nama, 
			p.email, 
			p.password, 
			p.role_id, 
			r.nama as nama_role 
			FROM pengelola as p 
			LEFT JOIN role_pengelola as r ON p.role_id = r.id
			WHERE email = ?`
	err = tx.QueryRowContext(ctx, SQL, email).Scan(&result.Id, &result.Nama, &result.Email, &result.Password, &result.RoleId, &result.NamaRole)
	return
}

func (r *RepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, pengelola *domain.Pengelola) (err error) {
	SQL := `UPDATE pengelola SET password = ? WHERE email = ?`
	_, err = tx.ExecContext(ctx, SQL, pengelola.Password, pengelola.Email)
	return
}