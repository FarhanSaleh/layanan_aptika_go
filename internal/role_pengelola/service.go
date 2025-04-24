package rolepengelola

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, request domain.RolePengelolaMutationRequest) (domain.RolePengelolaResponse, error)
	Update(ctx context.Context, request domain.RolePengelolaMutationRequest, id string) (domain.RolePengelolaResponse, error)
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]domain.RolePengelolaResponse, error)
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.RolePengelolaMutationRequest) (response domain.RolePengelolaResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return
	}

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()
		rolePengelola := domain.RolePengelola{
			Id: uuid,
			Nama: request.Nama,
		}

		err = s.Repository.Save(ctx, tx, &rolePengelola)
		if err != nil {
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.RolePengelolaResponse{
			Id: rolePengelola.Id,
			Nama: rolePengelola.Nama,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.RolePengelolaMutationRequest, id string) (response domain.RolePengelolaResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return
	}

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}
		result = domain.RolePengelola{
			Id: result.Id,
			Nama: request.Nama,
		}
		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		response = domain.RolePengelolaResponse{
			Id: id,
			Nama: request.Nama,
		}

		return 
	})
	
	return
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) (err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}

		err = s.Repository.Delete(ctx, tx, result.Id)
		if err != nil {
			log.Println("ERROR REPO <delete>:")
			return
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.RolePengelolaResponse, err error) {
	helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:", err)
			return
		}

		for _, role := range result {
			response = append(response, domain.RolePengelolaResponse{
				Id: role.Id,
				Nama: role.Nama,
			})
		}
		return
	})
	return
}