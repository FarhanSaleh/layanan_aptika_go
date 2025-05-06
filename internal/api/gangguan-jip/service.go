package gangguanjip

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
	Create(ctx context.Context, request domain.GangguanJIPMutationRequest) (domain.GangguanJIPMutationResponse, error)
	Update(ctx context.Context, request domain.GangguanJIPMutationRequest, id string) (domain.GangguanJIPMutationResponse, error)
	UpdateStatus(ctx context.Context, request domain.UpdateStatusLayananRequest, id string) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.GangguanJIPDetailResponse, error)
	FindAll(ctx context.Context) ([]domain.GangguanJIPResponse, error)
	FindAllByUser(ctx context.Context) ([]domain.GangguanJIPResponse, error)
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

func (s *ServiceImpl) Create(ctx context.Context, request domain.GangguanJIPMutationRequest) (response domain.GangguanJIPMutationResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		err = helper.MappingValidationError(err)
		return
	}
	jwtClaims := ctx.Value(contextkey.UserKey).(*domain.JWTClaims)

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		uuid := uuid.NewString()

		gangguanJIP := domain.GangguanJIP{
			Id: uuid,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			LokasiGangguan: request.LokasiGangguan,
			DeskripsiGangguan: request.DeskripsiGangguan,
			Foto: request.Foto,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
			UserId: jwtClaims.UID,
		}

		err = s.Repository.Save(ctx, tx, &gangguanJIP)
		if err != nil {
			log.Println("ERROR REPO <save>:", err)
			return
		}
		response = domain.GangguanJIPMutationResponse{
			Id: gangguanJIP.Id,
			NamaLengkap: gangguanJIP.NamaLengkap,
			Jabatan: gangguanJIP.Jabatan,
			NomorHP: gangguanJIP.NomorHP,
			LokasiGangguan: gangguanJIP.LokasiGangguan,
			DeskripsiGangguan: gangguanJIP.DeskripsiGangguan,
			SuratPermohonan: gangguanJIP.SuratPermohonan,
			Foto: gangguanJIP.Foto,
			InstansiId: gangguanJIP.InstansiId,
		}
		return
	})

	return
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.GangguanJIPMutationRequest, id string) (response domain.GangguanJIPMutationResponse, err error) {
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

		result = domain.GangguanJIP{
			Id: id,
			NamaLengkap: request.NamaLengkap,
			Jabatan: request.Jabatan,
			NomorHP: request.NomorHP,
			LokasiGangguan: request.LokasiGangguan,
			DeskripsiGangguan: request.DeskripsiGangguan,
			Foto: request.Foto,
			SuratPermohonan: request.SuratPermohonan,
			InstansiId: request.InstansiId,
		}

		err = s.Repository.Update(ctx, tx, &result)
		if err != nil {
			log.Println("ERROR REPO <update>:", err)
			return
		}
		response = domain.GangguanJIPMutationResponse{
			Id: id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			LokasiGangguan: result.LokasiGangguan,
			DeskripsiGangguan: result.DeskripsiGangguan,
			SuratPermohonan: result.SuratPermohonan,
			Foto: result.Foto,
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

		result = domain.GangguanJIP{
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

func (s *ServiceImpl) FindById(ctx context.Context, id string) (response domain.GangguanJIPDetailResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindById(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findById>:")
			return
		}

		suratPermohonanUrl := s.Config.StaticDocsOriginUser + result.SuratPermohonan
		fotoUrl := s.Config.StaticImgOriginUser + result.Foto

		response = domain.GangguanJIPDetailResponse{
			Id: result.Id,
			NamaLengkap: result.NamaLengkap,
			Jabatan: result.Jabatan,
			NomorHP: result.NomorHP,
			LokasiGangguan: result.LokasiGangguan,
			DeskripsiGangguan: result.DeskripsiGangguan,
			SuratPermohonan: suratPermohonanUrl,
			Foto: fotoUrl,
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

func (s *ServiceImpl) FindAll(ctx context.Context) (response []domain.GangguanJIPResponse, err error) {
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAll(ctx, tx)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, gangguanJIP := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + gangguanJIP.SuratPermohonan

			response = append(response, domain.GangguanJIPResponse{
				Id: gangguanJIP.Id,
				NamaLengkap: gangguanJIP.NamaLengkap,
				Jabatan: gangguanJIP.Jabatan,
				NomorHP: gangguanJIP.NomorHP,
				LokasiGangguan: gangguanJIP.LokasiGangguan,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: gangguanJIP.InstansiId,
				Status: gangguanJIP.Status,
				NamaInstansi: gangguanJIP.NamaInstansi,
				CreatedAt: gangguanJIP.CreatedAt.String(),
			})
		}
		return
	})
	return
}

func (s *ServiceImpl) FindAllByUser(ctx context.Context) (response []domain.GangguanJIPResponse, err error) {
	id := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.Repository.FindAllByUser(ctx, tx, id)
		if err != nil {
			log.Println("ERROR REPO <findAll>:")
			return
		}

		for _, gangguanJIP := range result {
			suratPermohonanUrl := s.Config.StaticDocsOriginUser + gangguanJIP.SuratPermohonan

			response = append(response, domain.GangguanJIPResponse{
				Id: gangguanJIP.Id,
				NamaLengkap: gangguanJIP.NamaLengkap,
				Jabatan: gangguanJIP.Jabatan,
				NomorHP: gangguanJIP.NomorHP,
				LokasiGangguan: gangguanJIP.LokasiGangguan,
				SuratPermohonan: suratPermohonanUrl,
				InstansiId: gangguanJIP.InstansiId,
				Status: gangguanJIP.Status,
				NamaInstansi: gangguanJIP.NamaInstansi,
				CreatedAt: gangguanJIP.CreatedAt.String(),
			})
		}
		return
	})
	return
}