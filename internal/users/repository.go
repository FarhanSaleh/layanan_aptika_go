package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, user *domain.User) error
	Update(ctx context.Context, tx *sql.Tx, user *domain.User) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository{
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user *domain.User) (err error) {
	SQL := `INSERT INTO users (id, nama, email, password) VALUES (?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, user.Id, user.Nama, user.Email, user.Password)
	return
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *domain.User) (err error) {
	SQL := `UPDATE users SET nama = ?, email = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, user.Nama, user.Email, user.Id)
	return
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) (err error) {
	SQL := `DELETE FROM users WHERE id = ?`
	_, err = tx.ExecContext(ctx, SQL, id)
	return
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (userResult domain.User, err error) {
	SQL := `SELECT id, nama, email, created_at FROM users WHERE id = ?`
	err = tx.QueryRowContext(ctx, SQL, id).Scan(&userResult.Id, &userResult.Nama, &userResult.Email, &userResult.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user tidak ditemukan")
			return
		}
		log.Println("ERROR QUERY: ", err)
		return
	}
	return 
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (usersResult []domain.User, err error) {
	SQL := `SELECT id, nama, email FROM users`

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		log.Println("ERROR QUERY: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u domain.User
		err = rows.Scan(&u.Id, &u.Nama, &u.Email)
		if err != nil{
			log.Println("ERROR SCANNING: ", err)
			return 
		}
		usersResult = append(usersResult, u)
	}

	return	
}

func (r *RepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (userResult domain.User, err error) {
	SQL := `SELECT id, nama, email, password FROM users WHERE email = ?`
	err = tx.QueryRowContext(ctx, SQL, email).Scan(&userResult.Id, &userResult.Nama, &userResult.Email, &userResult.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user tidak ditemukan")
			return
		}
		log.Println("ERROR QUERY: ", err)
		return
	}
	return
}