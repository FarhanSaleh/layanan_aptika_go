package pusatdatadaerah

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/constants"
	contextkey "github.com/farhansaleh/layanan_aptika_be/internal/context_key"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, request domain.PusatDataDaerahMutationRequest) (domain.PusatDataDaerahMutationResponse, error)
	Update(ctx context.Context, request domain.PusatDataDaerahMutationRequest, id string) (domain.PusatDataDaerahMutationResponse, error)
	UpdateStatus(ctx context.Context, request domain.UpdateStatusLayananRequest, id string) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.PusatDataDaerahDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.PusatDataDaerahResponse, error)
	FindAllByUser(ctx context.Context) ([]domain.PusatDataDaerahResponse, error)
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.PusatDataDaerahMutationRequest) (response domain.PusatDataDaerahMutationResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return
	}
	uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()

		pusatDataDaerah := domain.PusatDataDaerah{
			Id: uuid,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			JenisLayanan: request.JenisLayanan,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
			UserId: uid,
		}

		err = s.Repository.Save(ctx, tx, &pusatDataDaerah)
		if err != nil {
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.PusatDataDaerahMutationResponse{
			Id: pusatDataDaerah.Id,
			NamaLengkap: pusatDataDaerah.NamaLengkap,
			Jabatan: pusatDataDaerah.Jabatan,
			NomorHP: pusatDataDaerah.NomorHP,
			JenisLayanan: pusatDataDaerah.JenisLayanan,
			SuratPermohonan: pusatDataDaerah.SuratPermohonan,
			InstansiId: pusatDataDaerah.InstansiId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.PusatDataDaerahMutationRequest, id string) (response domain.PusatDataDaerahMutationResponse, err error) {
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

		result = domain.PusatDataDaerah{
			Id: id,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			JenisLayanan: request.JenisLayanan,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
		}

		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		response = domain.PusatDataDaerahMutationResponse{
			Id: id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			JenisLayanan: result.JenisLayanan,
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

		result = domain.PusatDataDaerah{
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

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.PusatDataDaerahDetailResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}

		suratPermohonanUrl := s.Config.StaticDocsOriginUser + result.SuratPermohonan

		response = domain.PusatDataDaerahDetailResponse{
			Id: result.Id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			JenisLayanan: result.JenisLayanan,
			SuratPermohonan: suratPermohonanUrl,
			InstansiId: result.InstansiId,
			Status: result.Status,
			NamaInstansi: result.NamaInstansi,
			CreatedAt: result.CreatedAt.Format(constants.TimeLayout),
			UpdatedAt: result.UpdatedAt.Time.Format(constants.TimeLayout),
		}
		return
	})
	
	return
}

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.PusatDataDaerahResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pusatDataDaerah := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + pusatDataDaerah.SuratPermohonan

			response = append(response, domain.PusatDataDaerahResponse{
				Id: pusatDataDaerah.Id,
				NamaLengkap: pusatDataDaerah.NamaLengkap,
				Jabatan: pusatDataDaerah.Jabatan,
				NomorHP: pusatDataDaerah.NomorHP,
				JenisLayanan: pusatDataDaerah.JenisLayanan,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: pusatDataDaerah.InstansiId,
				Status: pusatDataDaerah.Status,
				NamaInstansi: pusatDataDaerah.NamaInstansi,
				CreatedAt: pusatDataDaerah.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAllByUser(ctx context.Context) (response []domain.PusatDataDaerahResponse, err error) {
	id := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAllByUser(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pusatDataDaerah := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + pusatDataDaerah.SuratPermohonan

			response = append(response, domain.PusatDataDaerahResponse{
				Id: pusatDataDaerah.Id,
				NamaLengkap: pusatDataDaerah.NamaLengkap,
				Jabatan: pusatDataDaerah.Jabatan,
				NomorHP: pusatDataDaerah.NomorHP,
				JenisLayanan: pusatDataDaerah.JenisLayanan,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: pusatDataDaerah.InstansiId,
				Status: pusatDataDaerah.Status,
				NamaInstansi: pusatDataDaerah.NamaInstansi,
				CreatedAt: pusatDataDaerah.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}