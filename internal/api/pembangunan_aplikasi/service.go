package pembangunanaplikasi

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
	Create(ctx context.Context, request domain.PembangunanAplikasiMutationRequest) (domain.PembangunanAplikasiMutationResponse, error)
	Update(ctx context.Context, request domain.PembangunanAplikasiMutationRequest, id string) (domain.PembangunanAplikasiMutationResponse, error)
	UpdateStatus(ctx context.Context, request domain.UpdateStatusLayananRequest, id string) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.PembangunanAplikasiDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.PembangunanAplikasiResponse, error)
	FindAllByUser(ctx context.Context) ([]domain.PembangunanAplikasiResponse, error)
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.PembangunanAplikasiMutationRequest) (response domain.PembangunanAplikasiMutationResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return
	}
	uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()

		pembanguananAplikasi := domain.PembangunanAplikasi{
			Id: uuid,
			NamaPimpinan: request.NamaPimpinan,
			NomorHP: request.NomorHP,
			EmailDinas: request.EmailDinas,
			RiwayatPimpinan: request.RiwayatPimpinan,
			JenisAplikasi: request.JenisAplikasi,
			TujuanAplikasi: request.TujuanAplikasi,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
			UserId: uid,
		}

		err = s.Repository.Save(ctx, tx, &pembanguananAplikasi)
		if err != nil {
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.PembangunanAplikasiMutationResponse{
			Id: pembanguananAplikasi.Id,
			NamaPimpinan: pembanguananAplikasi.NamaPimpinan,
			NomorHP: pembanguananAplikasi.NomorHP,
			EmailDinas: pembanguananAplikasi.EmailDinas,
			RiwayatPimpinan: pembanguananAplikasi.RiwayatPimpinan,
			JenisAplikasi: pembanguananAplikasi.JenisAplikasi,
			TujuanAplikasi: pembanguananAplikasi.TujuanAplikasi,
			SuratPermohonan: pembanguananAplikasi.SuratPermohonan,
			InstansiId: pembanguananAplikasi.InstansiId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.PembangunanAplikasiMutationRequest, id string) (response domain.PembangunanAplikasiMutationResponse, err error) {
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

		result = domain.PembangunanAplikasi{
			Id: id,
			NamaPimpinan: request.NamaPimpinan,
			NomorHP: request.NomorHP,
			EmailDinas: request.EmailDinas,
			RiwayatPimpinan: request.RiwayatPimpinan,
			JenisAplikasi: request.JenisAplikasi,
			TujuanAplikasi: request.TujuanAplikasi,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
		}

		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		response = domain.PembangunanAplikasiMutationResponse{
			Id: id,
			NamaPimpinan: result.NamaPimpinan,
			NomorHP: result.NomorHP,
			EmailDinas: result.EmailDinas,
			RiwayatPimpinan: result.RiwayatPimpinan,
			JenisAplikasi: result.JenisAplikasi,
			TujuanAplikasi: result.TujuanAplikasi,
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
				"Layanan Permohonan Pembangunan Aplikasi", 
				fmt.Sprintf("Permintaan anda atas nama %s, pada tanggal %s, telah %s",
					result.NamaPimpinan, 
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

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.PembangunanAplikasiDetailResponse, err error) {
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

		response = domain.PembangunanAplikasiDetailResponse{
			Id: result.Id,
			NamaPimpinan: result.NamaPimpinan,
			NomorHP: result.NomorHP,
			EmailDinas: result.EmailDinas,
			RiwayatPimpinan: result.RiwayatPimpinan,
			JenisAplikasi: result.JenisAplikasi,
			TujuanAplikasi: result.TujuanAplikasi,
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

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.PembangunanAplikasiResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pembanguananAplikasi := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginPengelola + pembanguananAplikasi.SuratPermohonan

			response = append(response, domain.PembangunanAplikasiResponse{
				Id: pembanguananAplikasi.Id,
				NamaPimpinan: pembanguananAplikasi.NamaPimpinan,
				NomorHP: pembanguananAplikasi.NomorHP,
				EmailDinas: pembanguananAplikasi.EmailDinas,
				JenisAplikasi: pembanguananAplikasi.JenisAplikasi,
				SuratPermohonan: suratPermohonanUrl,
				Status: pembanguananAplikasi.Status,
				InstansiId: pembanguananAplikasi.InstansiId,
				NamaInstansi: pembanguananAplikasi.NamaInstansi,
				CreatedAt: pembanguananAplikasi.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAllByUser(ctx context.Context) (response []domain.PembangunanAplikasiResponse, err error) {
	id := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAllByUser(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, pembanguananAplikasi := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + pembanguananAplikasi.SuratPermohonan

			response = append(response, domain.PembangunanAplikasiResponse{
				Id: pembanguananAplikasi.Id,
				NamaPimpinan: pembanguananAplikasi.NamaPimpinan,
				NomorHP: pembanguananAplikasi.NomorHP,
				EmailDinas: pembanguananAplikasi.EmailDinas,
				JenisAplikasi: pembanguananAplikasi.JenisAplikasi,
				SuratPermohonan: suratPermohonanUrl,
				Status: pembanguananAplikasi.Status,
				InstansiId: pembanguananAplikasi.InstansiId,
				NamaInstansi: pembanguananAplikasi.NamaInstansi,
				CreatedAt: pembanguananAplikasi.CreatedAt.Format(constants.TimeLayout),
			})
		}
		return
	})
	return
}