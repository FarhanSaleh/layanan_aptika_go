package pengelola

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
	Create(ctx context.Context, request domain.PengelolaMutationRequest) (domain.PengelolaMutateResponse, error)
	Update(ctx context.Context, request domain.PengelolaMutationRequest, id string) (domain.PengelolaMutateResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.PengelolaDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.PengelolaResponse, error)
}

type ServiceImpl struct {
	Repository Repository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewService(db *sql.DB, repository Repository, validate *validator.Validate) Service {
	return &ServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, request domain.PengelolaMutationRequest) (response domain.PengelolaMutateResponse, err error) {
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
		
		pengelola := domain.Pengelola{
			Id: uuid,
			Nama: request.Nama,
			Email: request.Email,
			Password: string(hashPassword),
			RoleId: request.RoleId,
		}
	
		err = s.Repository.Save(ctx, tx, &pengelola)
		if err != nil{
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.PengelolaMutateResponse{
			Id: pengelola.Id,
			Nama: pengelola.Nama,
			Email: pengelola.Email,
			RoleId: pengelola.RoleId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.PengelolaMutationRequest, id string) (response domain.PengelolaMutateResponse, err error){
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
		
		result = domain.Pengelola{
			Id: result.Id,
			Nama: request.Nama,
			Email: request.Email,
			RoleId: request.RoleId,
		}
	
		err = s.Repository.Update(ctx, tx, &result)
		if err != nil{
			log.Println("ERROR REPO <update>:", err)
			return
		}
	
		response = domain.PengelolaMutateResponse{
			Id: id,
			Nama: result.Nama,
			Email: result.Email,
			RoleId: result.RoleId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) (err error){
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		pengelola, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:", err)
			return
		}
	
		err = s.Repository.Delete(ctx, tx, pengelola.Id)
		if err != nil {
			log.Println("ERROR REPO <delete>:", err)
			return
		}

		return 
	})
	return
}

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.PengelolaDetailResponse, err error){
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		pengelola, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:", err)
			return
		}
		response = domain.PengelolaDetailResponse{
			Id: pengelola.Id,
			Nama: pengelola.Nama,
			Email: pengelola.Email,
			RoleId: pengelola.RoleId,
			NamaRole: pengelola.NamaRole,
			CreatedAt: pengelola.CreatedAt.GoString(),
		}
		return
	})

	return
}

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.PengelolaResponse, err error){
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:", err)
			return
		}
	
		for _, pengelola := range result {
			response = append(response, domain.PengelolaResponse{
				Id: pengelola.Id,
				Nama: pengelola.Nama,
				Email: pengelola.Email,
				RoleId: pengelola.RoleId,
				NamaRole: pengelola.NamaRole,
			})
		}
		return
	})

	return
}