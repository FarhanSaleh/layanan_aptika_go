package users

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(ctx context.Context, request domain.UserMutationRequest) (domain.UserResponse, error)
	Update(ctx context.Context, request domain.UserMutationRequest, id string) (domain.UserResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.UserDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.UserResponse, error)
}

type ServiceImpl struct {
	Repository Repository
	DB *sql.DB
	Validate *validator.Validate
}

func NewService(db *sql.DB, repository Repository, validate *validator.Validate) Service{
	return &ServiceImpl{
		Repository: repository,
		DB: db,
		Validate: validate,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, request domain.UserMutationRequest) (userResponse domain.UserResponse, err error){
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return 
	}

	tx, err := s.DB.Begin()
	if err != nil {
		log.Println("Error tx: ", err)
		return
	}
	defer func(){
		log.Println("DEFER ERR: ", err)
		if err != nil {
			if errRoleback := tx.Rollback(); errRoleback != nil{
				log.Println("ERROR ROLLBACK: ", err)
				return
			}
		}	
		if err := tx.Commit(); err != nil {
			return
		}
	}()

	uuid := uuid.NewString()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("112233"), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ERROR HASH PASSWORD: ", err)
		return
	}
	
	user := domain.User{
		Id: uuid,
		Nama: request.Nama,
		Email: request.Email,
		Password: string(hashPassword),
	}

	err = s.Repository.Save(ctx, tx, &user)
	if err != nil{
		log.Println("Error repo: ", err)
		return
	}

	userResponse = domain.UserResponse{
		Id: user.Id,
		Nama: user.Nama,
		Email: user.Email,
	}
	return
}
func (s *ServiceImpl) Update(ctx context.Context, request domain.UserMutationRequest, id string) (userResponse domain.UserResponse, err error){
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return 
	}

	tx, err := s.DB.Begin()
	if err != nil {
		log.Println("Error tx: ", err)
		return
	}
	defer func(){
		log.Println("DEFER ERR: ", err)
		if err != nil {
			if errRoleback := tx.Rollback(); errRoleback != nil{
				log.Println("ERROR ROLLBACK: ", err)
				return
			}
		}	
		if err := tx.Commit(); err != nil {
			return
		}
	}()

	user := domain.User{
		Id: id,
		Nama: request.Nama,
		Email: request.Email,
	}

	err = s.Repository.Update(ctx, tx, &user)
	if err != nil{
		log.Println("Error repo: ", err)
		return
	}

	userResponse = domain.UserResponse{
		Id: id,
		Nama: user.Nama,
		Email: user.Email,
	}

	return
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) (err error){
	tx, err := s.DB.Begin()
	if err != nil {
		log.Println("Error tx: ", err)
		return
	}
	defer func(){
		log.Println("DEFER ERR: ", err)
		if err != nil {
			if errRoleback := tx.Rollback(); errRoleback != nil{
				log.Println("ERROR ROLLBACK: ", err)
				return
			}
		}	
		if err := tx.Commit(); err != nil {
			return
		}
	}()

	user, err := s.Repository.FindById(ctx, tx, id)
	if err != nil {
		log.Println("Error Repository", err)
		return
	}

	err = s.Repository.Delete(ctx, tx, user.Id)
	if err != nil {
		log.Println("Error Repository", err)
		return
	}

	return
}

func (s *ServiceImpl) FindById(ctx context.Context, id string) (userResponse domain.UserDetailResponse, err error){
	tx, err := s.DB.Begin()
	if err != nil {
		log.Println("Error tx: ", err)
		return
	}
	defer func(){
		log.Println("DEFER ERR: ", err)
		if err != nil {
			if errRoleback := tx.Rollback(); errRoleback != nil{
				log.Println("ERROR ROLLBACK: ", err)
				return
			}
		}	
		if err := tx.Commit(); err != nil {
			return
		}
	}()

	user, err := s.Repository.FindById(ctx, tx, id)
	if err != nil {
		log.Println("Error Repository", err)
		return
	}
	userResponse = domain.UserDetailResponse{
		Id: user.Id,
		Nama: user.Nama,
		Email: user.Email,
		CreatedAt: user.CreatedAt.Format("02 January 2006 pukul 15:04"),
	}
	return
}

func (s *ServiceImpl) FindAll(ctx context.Context) (userResponse []domain.UserResponse, err error){
	tx, err := s.DB.Begin()
	if err != nil {
		log.Println("Error tx: ", err)
		return
	}
	defer func(){
		log.Println("DEFER ERR: ", err)
		if err != nil {
			if errRoleback := tx.Rollback(); errRoleback != nil{
				log.Println("ERROR ROLLBACK: ", err)
				return
			}
		}	
		if err := tx.Commit(); err != nil {
			return
		}
	}()

	users, err := s.Repository.FindAll(ctx, tx)
	if err != nil {
		log.Println("Error Repository", err)
		return
	}

	for _, user := range users{
		userResponse = append(userResponse, domain.UserResponse{
			Id: user.Id,
			Nama: user.Nama,
			Email: user.Email,
		})
	}

	return
}