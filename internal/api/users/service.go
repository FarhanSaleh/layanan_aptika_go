package users

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/constants"
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.UserMutationRequest) (response domain.UserResponse, err error){
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return 
	}

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()
		hashPassword, err := bcrypt.GenerateFromPassword([]byte("112233"), bcrypt.DefaultCost)
		if err != nil {
			log.Println("ERROR HASH PASSWORD:", err)
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
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.UserResponse{
			Id: user.Id,
			Nama: user.Nama,
			Email: user.Email,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.UserMutationRequest, id string) (response domain.UserResponse, err error){
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return 
	}

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}
		
		result = domain.User{
			Id: result.Id,
			Nama: request.Nama,
			Email: request.Email,
		}
	
		err = s.Repository.Update(ctx, tx, &result)
		if err != nil{
			log.Println("ERROR REPO <update>:", err)
			return
		}
	
		response = domain.UserResponse{
			Id: id,
			Nama: result.Nama,
			Email: result.Email,
		}
		return
	})

	return
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) (err error){
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		user, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:", err)
			return
		}
	
		err = s.Repository.Delete(ctx, tx, user.Id)
		if err != nil {
			log.Println("ERROR REPO <delete>:", err)
			return
		}

		return 
	})
	return
}

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.UserDetailResponse, err error){
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		user, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:", err)
			return
		}
		response = domain.UserDetailResponse{
			Id: user.Id,
			Nama: user.Nama,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format(constants.TimeLayout),
		}
		return
	})

	return
}

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.UserResponse, err error){
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		users, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:", err)
			return
		}
	
		for _, user := range users{
			response = append(response, domain.UserResponse{
				Id: user.Id,
				Nama: user.Nama,
				Email: user.Email,
			})
		}
		return
	})

	return
}