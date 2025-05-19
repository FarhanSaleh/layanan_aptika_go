package permintaan

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/config"
	contextkey "github.com/farhansaleh/layanan_aptika_be/internal/context_key"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
)

type Service interface {
	CountAll(ctx context.Context) (domain.PermintaanCountResponse, error)
	CountGangguanJIP(ctx context.Context) (domain.PermintaanCountResponse, error)
	CountPembuatanEmail(ctx context.Context) (domain.PermintaanCountResponse, error)
	CountPembuatanSubdomain(ctx context.Context) (domain.PermintaanCountResponse, error)
	CountPembangunanAplikasi(ctx context.Context) (domain.PermintaanCountResponse, error)
	CountPusatDataDaerah(ctx context.Context) (domain.PermintaanCountResponse, error)
	CountPerubahanIPServer(ctx context.Context) (domain.PermintaanCountResponse, error)
}

type ServiceImpl struct {
	Repository Repository
	DB *sql.DB
	Config *config.Config
}

func NewService(db *sql.DB, config *config.Config, repository Repository ) Service {
	return &ServiceImpl{
		DB:         db,
		Config: config,
		Repository: repository,
	}
}

func (s *ServiceImpl) CountAll(ctx context.Context) (response domain.PermintaanCountResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		if accountType == "user" {
			uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
			response, err = s.Repository.CountAllByUser(ctx, tx, uid)
			if err != nil {
				log.Println("ERROR REPO <countAllByUser>:")
				return
			}
			return
		}
		response, err = s.Repository.CountAll(ctx, tx)

		if err != nil {
			log.Println("ERROR REPO <countAll>:")
			return
		}
		return
	})
	return
}

func (s *ServiceImpl) CountGangguanJIP(ctx context.Context) (response domain.PermintaanCountResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		if accountType == "user" {
			uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
			response, err = s.Repository.CountGangguanJIPByUser(ctx, tx, uid)
			if err != nil {
				log.Println("ERROR REPO <countGangguanJIPByUser>:")
				return
			}
			return
		}
		response, err = s.Repository.CountGangguanJIP(ctx, tx)

		if err != nil {
			log.Println("ERROR REPO <countGangguanJIP>:")
			return
		}
		return
	})
	return
}

func (s *ServiceImpl) CountPembuatanEmail(ctx context.Context) (response domain.PermintaanCountResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		if accountType == "user" {
			uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
			response, err = s.Repository.CountPembuatanEmailByUser(ctx, tx, uid)
			if err != nil {
				log.Println("ERROR REPO <countPembuatanEmailByUser>:")
				return
			}
			return
		}
		response, err = s.Repository.CountPembuatanEmail(ctx, tx)

		if err != nil {
			log.Println("ERROR REPO <countPembuatanEmail>:")
			return
		}
		return
	})
	return
}

func (s *ServiceImpl) CountPembuatanSubdomain(ctx context.Context) (response domain.PermintaanCountResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		if accountType == "user" {
			uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
			response, err = s.Repository.CountPembuatanSubdomainByUser(ctx, tx, uid)
			if err != nil {
				log.Println("ERROR REPO <countPembuatanSubdomainByUser>:")
				return
			}
			return
		}
		response, err = s.Repository.CountPembuatanSubdomain(ctx, tx)

		if err != nil {
			log.Println("ERROR REPO <countPembuatanSubdomain>:")
			return
		}
		return
	})
	return
}

func (s *ServiceImpl) CountPembangunanAplikasi(ctx context.Context) (response domain.PermintaanCountResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		if accountType == "user" {
			uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
			response, err = s.Repository.CountPembangunanAplikasiByUser(ctx, tx, uid)
			if err != nil {
				log.Println("ERROR REPO <countPembangunanAplikasiByUser>:")
				return
			}
			return
		}
		response, err = s.Repository.CountPembangunanAplikasi(ctx, tx)

		if err != nil {
			log.Println("ERROR REPO <countPembangunanAplikasi>:")
			return
		}
		return
	})
	return
}

func (s *ServiceImpl) CountPusatDataDaerah(ctx context.Context) (response domain.PermintaanCountResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		if accountType == "user" {
			uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
			response, err = s.Repository.CountPusatDataDaerahByUser(ctx, tx, uid)
			if err != nil {
				log.Println("ERROR REPO <countPusatDataDaerahByUser>:")
				return
			}
			return
		}
		response, err = s.Repository.CountPusatDataDaerah(ctx, tx)

		if err != nil {
			log.Println("ERROR REPO <countPusatDataDaerah>:")
			return
		}
		return
	})
	return
}

func (s *ServiceImpl) CountPerubahanIPServer(ctx context.Context) (response domain.PermintaanCountResponse, err error) {
	accountType := ctx.Value(contextkey.TypeAccountKey).(string)
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		if accountType == "user" {
			uid := ctx.Value(contextkey.UserKey).(*domain.JWTClaims).UID
			response, err = s.Repository.CountPerubahanIPServerByUser(ctx, tx, uid)
			if err != nil {
				log.Println("ERROR REPO <countPerubahanIPServerByUser>:")
				return
			}
			return
		}
		response, err = s.Repository.CountPerubahanIPServer(ctx, tx)

		if err != nil {
			log.Println("ERROR REPO <countPerubahanIPServer>:")
			return
		}
		return
	})
	return
}