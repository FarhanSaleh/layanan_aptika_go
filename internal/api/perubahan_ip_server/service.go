package perubahanipserver

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/config"
	contextkey "github.com/farhansaleh/layanan_aptika_be/internal/context_key"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, request domain.PerubahanIPServerMutationRequest) (domain.PerubahanIPServerMutationResponse, error)
	Update(ctx context.Context, request domain.PerubahanIPServerMutationRequest, id string) (domain.PerubahanIPServerMutationResponse, error)
	UpdateStatus(ctx context.Context, request domain.UpdateStatusLayananRequest, id string) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.PerubahanIPServerDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.PerubahanIPServerResponse, error)
	FindAllByUser(ctx context.Context) ([]domain.PerubahanIPServerResponse, error)
}

type ServiceImpl struct {
	Repository Repository
	DB         *sql.DB
	Validate   *validator.Validate
	Config	   *config.Config
}

func NewService(db *sql.DB, repository Repository, validate *validator.Validate, config *config.Config) Service {
	return &ServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
		Config: config,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, request domain.PerubahanIPServerMutationRequest) (response domain.PerubahanIPServerMutationResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return
	}
	uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()

		perubahanIPServer := domain.PerubahanIPServer{
			Id: uuid,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			NamaSubdomain: request.NamaSubdomain,
			IPLama: request.IPLama,
			IPBaru: request.IPBaru,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
			UserId: uid,
		}

		err = s.Repository.Save(ctx, tx, &perubahanIPServer)
		if err != nil {
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.PerubahanIPServerMutationResponse{
			Id: perubahanIPServer.Id,
			NamaLengkap: perubahanIPServer.NamaLengkap,
			Jabatan: perubahanIPServer.Jabatan,
			NomorHP: perubahanIPServer.NomorHP,
			NamaSubdomain: perubahanIPServer.NamaSubdomain,
			IPLama: perubahanIPServer.IPLama,
			IPBaru: perubahanIPServer.IPBaru,
			SuratPermohonan: perubahanIPServer.SuratPermohonan,
			InstansiId: perubahanIPServer.InstansiId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.PerubahanIPServerMutationRequest, id string) (response domain.PerubahanIPServerMutationResponse, err error) {
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

		result = domain.PerubahanIPServer{
			Id: id,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			NamaSubdomain: request.NamaSubdomain,
			IPLama: request.IPLama,
			IPBaru: request.IPBaru,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
		}

		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		response = domain.PerubahanIPServerMutationResponse{
			Id: id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			NamaSubdomain: result.NamaSubdomain,
			IPLama: result.IPLama,
			IPBaru: result.IPBaru,
			SuratPermohonan: result.SuratPermohonan,
			InstansiId: result.InstansiId,
		}
		return
	})

	return
}

func (s *ServiceImpl) UpdateStatus(ctx context.Context, request domain.UpdateStatusLayananRequest, id string) (err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}

		result = domain.PerubahanIPServer{
			Id: id,
			Status: request.Status,
		}

		err = s.Repository.UpdateStatus(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
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

		if result.Status != "diproses" {
			err = helper.NewBadRequestError("tidak dapat dihapus karena statusnya bukan diproses")
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

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.PerubahanIPServerDetailResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}

		suratPermohonanUrl := s.Config.StaticDocsOriginUser + result.SuratPermohonan

		response = domain.PerubahanIPServerDetailResponse{
			Id: result.Id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			NamaSubdomain: result.NamaSubdomain,
			IPLama: result.IPLama,
			IPBaru: result.IPBaru,
			SuratPermohonan: suratPermohonanUrl,
			InstansiId: result.InstansiId,
			Status: result.Status,
			NamaInstansi: result.NamaInstansi,
			CreatedAt: result.CreatedAt.String(),
			UpdatedAt: result.UpdatedAt.Time.String(),
		}
		return
	})
	
	return
}

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.PerubahanIPServerResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, perubahanIPServer := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + perubahanIPServer.SuratPermohonan

			response = append(response, domain.PerubahanIPServerResponse{
				Id: perubahanIPServer.Id,
				NamaLengkap: perubahanIPServer.NamaLengkap,
				Jabatan: perubahanIPServer.Jabatan,
				NomorHP: perubahanIPServer.NomorHP,
				NamaSubdomain: perubahanIPServer.NamaSubdomain,
				IPLama: perubahanIPServer.IPLama,
				IPBaru: perubahanIPServer.IPBaru,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: perubahanIPServer.InstansiId,
				Status: perubahanIPServer.Status,
				NamaInstansi: perubahanIPServer.NamaInstansi,
				CreatedAt: perubahanIPServer.CreatedAt.String(),
			})
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAllByUser(ctx context.Context) (response []domain.PerubahanIPServerResponse, err error) {
	id := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAllByUser(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, perubahanIPServer := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + perubahanIPServer.SuratPermohonan

			response = append(response, domain.PerubahanIPServerResponse{
				Id: perubahanIPServer.Id,
				NamaLengkap: perubahanIPServer.NamaLengkap,
				Jabatan: perubahanIPServer.Jabatan,
				NomorHP: perubahanIPServer.NomorHP,
				NamaSubdomain: perubahanIPServer.NamaSubdomain,
				IPLama: perubahanIPServer.IPLama,
				IPBaru: perubahanIPServer.IPBaru,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: perubahanIPServer.InstansiId,
				Status: perubahanIPServer.Status,
				NamaInstansi: perubahanIPServer.NamaInstansi,
				CreatedAt: perubahanIPServer.CreatedAt.String(),
			})
		}
		return
	})
	return
}