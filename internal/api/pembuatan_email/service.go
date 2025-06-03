package pembuatanemail

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
	Create(ctx context.Context, request domain.PembuatanEmailMutationRequest) (domain.PembuatanEmailMutationResponse, error)
	Update(ctx context.Context, request domain.PembuatanEmailMutationRequest, id string) (domain.PembuatanEmailMutationResponse, error)
	UpdateStatus(ctx context.Context, request domain.UpdateStatusLayananRequest, id string) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.PembuatanEmailDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.PembuatanEmailResponse, error)
	FindAllByUser(ctx context.Context) ([]domain.PembuatanEmailResponse, error)
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.PembuatanEmailMutationRequest) (response domain.PembuatanEmailMutationResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return
	}
	uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()

		pembuatanEmail := domain.PembuatanEmail{
			Id: uuid,
			NamaLengkap: request.NamaLengkap,
			NIP: request.NIP,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			BerkasSK: request.BerkasSK,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
			UserId: uid,
		}

		err = s.Repository.Save(ctx, tx, &pembuatanEmail)
		if err != nil {
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.PembuatanEmailMutationResponse{
			Id: pembuatanEmail.Id,
			NamaLengkap: pembuatanEmail.NamaLengkap,
			NIP: pembuatanEmail.NIP,
			Jabatan: pembuatanEmail.Jabatan,
			NomorHP: pembuatanEmail.NomorHP,
			BerkasSK: pembuatanEmail.BerkasSK,
			SuratPermohonan: pembuatanEmail.SuratPermohonan,
			InstansiId: pembuatanEmail.InstansiId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.PembuatanEmailMutationRequest, id string) (response domain.PembuatanEmailMutationResponse, err error) {
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

		result = domain.PembuatanEmail{
			Id: id,
			NamaLengkap: request.NamaLengkap,
			NIP: request.NIP,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			BerkasSK: request.BerkasSK,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
		}

		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		response = domain.PembuatanEmailMutationResponse{
			Id: id,
			NamaLengkap: result.NamaLengkap,
			NIP: result.NIP,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			BerkasSK: result.BerkasSK,
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
				"Layanan Pembuatan Email", 
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
		
		if result.BerkasSK != "" {
			err = helper.DeleteFile(result.BerkasSK, "docs")
			if err != nil {
				log.Println("ERROR DELETING BERKAS SK:", err)
			}
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

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.PembuatanEmailDetailResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}
		var suratPermohonanUrl string
		var berkasSKUrl string
		if accountType == "pengelola" {
			suratPermohonanUrl = s.Config.StaticDocsOriginPengelola + result.SuratPermohonan
			berkasSKUrl = s.Config.StaticDocsOriginPengelola + result.BerkasSK
		} else {
			suratPermohonanUrl = s.Config.StaticDocsOriginUser + result.SuratPermohonan
			berkasSKUrl = s.Config.StaticDocsOriginUser + result.BerkasSK
		}

		response = domain.PembuatanEmailDetailResponse{
			Id: result.Id,
			NamaLengkap: result.NamaLengkap,
			NIP: result.NIP,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			BerkasSK: berkasSKUrl,
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

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.PembuatanEmailResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pembuatanEmail := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginPengelola + pembuatanEmail.SuratPermohonan
			berkasSKUrl := s.Config.StaticDocsOriginPengelola + pembuatanEmail.BerkasSK

			response = append(response, domain.PembuatanEmailResponse{
				Id: pembuatanEmail.Id,
				NamaLengkap: pembuatanEmail.NamaLengkap,
				NIP: pembuatanEmail.NIP,
				Jabatan: pembuatanEmail.Jabatan,
				NomorHP: pembuatanEmail.NomorHP,
				BerkasSK: berkasSKUrl,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: pembuatanEmail.InstansiId,
				Status: pembuatanEmail.Status,
				NamaInstansi: pembuatanEmail.NamaInstansi,
				CreatedAt: pembuatanEmail.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAllByUser(ctx context.Context) (response []domain.PembuatanEmailResponse, err error) {
	id := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAllByUser(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pembuatanEmail := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + pembuatanEmail.SuratPermohonan
			berkasSKUrl := s.Config.StaticDocsOriginUser + pembuatanEmail.BerkasSK

			response = append(response, domain.PembuatanEmailResponse{
				Id: pembuatanEmail.Id,
				NamaLengkap: pembuatanEmail.NamaLengkap,
				NIP: pembuatanEmail.NIP,
				Jabatan: pembuatanEmail.Jabatan,
				NomorHP: pembuatanEmail.NomorHP,
				BerkasSK: berkasSKUrl,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: pembuatanEmail.InstansiId,
				Status: pembuatanEmail.Status,
				NamaInstansi: pembuatanEmail.NamaInstansi,
				CreatedAt: pembuatanEmail.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}