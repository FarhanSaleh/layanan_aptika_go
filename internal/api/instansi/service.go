package instansi

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
	Create(ctx context.Context, request domain.InstansiMutationRequest) (domain.InstansiResponse, error)
	Update(ctx context.Context, request domain.InstansiMutationRequest, id string) (domain.InstansiResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.InstansiResponse, error)
	FindAll(ctx context.Context) ([]domain.InstansiResponse, error)
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.InstansiMutationRequest) (response domain.InstansiResponse, err error) {
	keterangan := helper.StringToNullString(request.Keterangan)
	
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return 
	}

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()
		instansi := domain.Instansi{
			Id: 	 	uuid,
			Nama:    	request.Nama,
			Alamat:  	request.Alamat,
			Keterangan: keterangan,
		}
	
		err = s.Repository.Save(ctx, tx, &instansi)
		if err != nil{
			log.Println("ERROR REPO <save>: ", err)
			return
		}
	
		response = domain.InstansiResponse{
			Id:			instansi.Id,
			Nama:    	instansi.Nama,
			Alamat:  	instansi.Alamat,
			Keterangan: instansi.Keterangan.String,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.InstansiMutationRequest, id string) (response domain.InstansiResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return
	}

	keterangan := helper.StringToNullString(request.Keterangan)

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}

		result = domain.Instansi{
			Id: result.Id,
			Nama: request.Nama,
			Alamat: request.Alamat,
			Keterangan: keterangan,
		}
		
		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}

		response = domain.InstansiResponse{
			Id: id,
			Nama: request.Nama,
			Alamat: request.Alamat,
			Keterangan: request.Keterangan,
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

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.InstansiResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil{
			log.Println("ERROR REPO <findById>:", err)
			return
		}
		response = domain.InstansiResponse{
			Id: result.Id,
			Nama: result.Nama,
			Alamat: result.Nama,
			Keterangan: result.Keterangan.String,
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.InstansiResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:", err)
			return
		}

		for _, instansi := range result {
			response = append(response, domain.InstansiResponse{
				Id: instansi.Id,
				Nama: instansi.Nama,
				Alamat: instansi.Alamat,
				Keterangan: instansi.Keterangan.String,
			})
		}
		return
	})
	return
}