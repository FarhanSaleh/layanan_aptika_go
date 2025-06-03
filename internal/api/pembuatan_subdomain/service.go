package pembuatansubdomain

import (
	"context"
	"database/sql"
	"fmt"
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
	Create(ctx context.Context, request domain.PembuatanSubdomainMutationRequest) (domain.PembuatanSubdomainMutationResponse, error)
	Update(ctx context.Context, request domain.PembuatanSubdomainMutationRequest, id string) (domain.PembuatanSubdomainMutationResponse, error)
	UpdateStatus(ctx context.Context, request domain.UpdateStatusLayananRequest, id string) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.PembuatanSubdomainDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.PembuatanSubdomainResponse, error)
	FindAllByUser(ctx context.Context) ([]domain.PembuatanSubdomainResponse, error)
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.PembuatanSubdomainMutationRequest) (response domain.PembuatanSubdomainMutationResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return
	}
	uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()

		pembuatanSubdomain := domain.PembuatanSubdomain{
			Id: uuid,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			NamaSubdomain: request.NamaSubdomain,
			IPPublik: request.IPPublik,
			Deskripsi: request.Deskripsi,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
			UserId: uid,
		}

		err = s.Repository.Save(ctx, tx, &pembuatanSubdomain)
		if err != nil {
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.PembuatanSubdomainMutationResponse{
			Id: pembuatanSubdomain.Id,
			NamaLengkap: pembuatanSubdomain.NamaLengkap,
			Jabatan: pembuatanSubdomain.Jabatan,
			NomorHP: pembuatanSubdomain.NomorHP,
			NamaSubdomain: pembuatanSubdomain.NamaSubdomain,
			IPPublik: pembuatanSubdomain.IPPublik,
			Deskripsi: pembuatanSubdomain.Deskripsi,
			SuratPermohonan: pembuatanSubdomain.SuratPermohonan,
			InstansiId: pembuatanSubdomain.InstansiId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.PembuatanSubdomainMutationRequest, id string) (response domain.PembuatanSubdomainMutationResponse, err error) {
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

		result = domain.PembuatanSubdomain{
			Id: id,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			NamaSubdomain: result.NamaSubdomain,
			IPPublik: result.IPPublik,
			Deskripsi: result.Deskripsi,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
		}

		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		response = domain.PembuatanSubdomainMutationResponse{
			Id: id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			NamaSubdomain: result.NamaSubdomain,
			IPPublik: result.IPPublik,
			Deskripsi: result.Deskripsi,
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

		result.Status = request.Status

		err = s.Repository.UpdateStatus(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		
		if result.NotificationToken.Valid {
			log.Println("PUSH NOTIFICATION")
			helper.SendPushNotification(
				result.NotificationToken.String,
				"Layanan Pembuatan Subdomain", 
				fmt.Sprintf("Permintaan anda atas nama %s, pada tanggal %s, telah %s",
					result.NamaLengkap,
					result.CreatedAt.Format(constants.TimeLayoutForNotif), 
					request.Status,
				),
			)
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

		if result.SuratPermohonan != "" {
			err = helper.DeleteFile(result.SuratPermohonan, "docs")
			if err != nil {
				log.Println("ERROR DELETING SURAT PERMOHONAN:", err)
			}
		}
		return
	})
	return
}

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.PembuatanSubdomainDetailResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}
		var suratPermohonanUrl string
		if accountType == "pengelola" {
			suratPermohonanUrl = s.Config.StaticDocsOriginPengelola + result.SuratPermohonan
		} else {
			suratPermohonanUrl = s.Config.StaticDocsOriginUser + result.SuratPermohonan
		}

		response = domain.PembuatanSubdomainDetailResponse{
			Id: result.Id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			NamaSubdomain: result.NamaSubdomain,
			IPPublik: result.IPPublik,
			Deskripsi: result.Deskripsi,
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

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.PembuatanSubdomainResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pembuatanSubdomain := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginPengelola + pembuatanSubdomain.SuratPermohonan

			response = append(response, domain.PembuatanSubdomainResponse{
				Id: pembuatanSubdomain.Id,
				NamaLengkap: pembuatanSubdomain.NamaLengkap,
				Jabatan: pembuatanSubdomain.Jabatan,
				NomorHP: pembuatanSubdomain.NomorHP,
				NamaSubdomain: pembuatanSubdomain.NamaSubdomain,
				IPPublik: pembuatanSubdomain.IPPublik,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: pembuatanSubdomain.InstansiId,
				Status: pembuatanSubdomain.Status,
				NamaInstansi: pembuatanSubdomain.NamaInstansi,
				CreatedAt: pembuatanSubdomain.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAllByUser(ctx context.Context) (response []domain.PembuatanSubdomainResponse, err error) {
	id := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAllByUser(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pembuatanSubdomain := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + pembuatanSubdomain.SuratPermohonan

			response = append(response, domain.PembuatanSubdomainResponse{
				Id: pembuatanSubdomain.Id,
				NamaLengkap: pembuatanSubdomain.NamaLengkap,
				Jabatan: pembuatanSubdomain.Jabatan,
				NomorHP: pembuatanSubdomain.NomorHP,
				NamaSubdomain: pembuatanSubdomain.NamaSubdomain,
				IPPublik: pembuatanSubdomain.IPPublik,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: pembuatanSubdomain.InstansiId,
				Status: pembuatanSubdomain.Status,
				NamaInstansi: pembuatanSubdomain.NamaInstansi,
				CreatedAt: pembuatanSubdomain.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}