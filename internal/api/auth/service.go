package auth

import (
	"context"
	"database/sql"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/internal/api/pengelola"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/users"
	contextkey "github.com/farhansaleh/layanan_aptika_be/internal/context_key"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, request domain.LoginRequest) (domain.LoginResponse, error)
	PengelolaLogin(ctx context.Context, request domain.LoginRequest) (domain.LoginResponse, error)
	Logout(ctx context.Context) error
	PengelolaChangePassword(ctx context.Context, request domain.ChangePasswordRequest) error
	UserChangePassword(ctx context.Context, request domain.ChangePasswordRequest) error
}

type ServiceImpl struct {
	UserRepository users.Repository
	PengelolaRepository pengelola.Repository
	DB *sql.DB
	Validate *validator.Validate
}

func NewService(db *sql.DB, userRepository users.Repository, pengelolaRepository pengelola.Repository, validate *validator.Validate) Service {
	return &ServiceImpl{
		UserRepository: userRepository,
		PengelolaRepository: pengelolaRepository,
		DB: db,
		Validate: validate,
	}
}

func (s *ServiceImpl) Login(ctx context.Context, request domain.LoginRequest) (response domain.LoginResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return
	}
	
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		user, err := s.UserRepository.FindByEmail(ctx, tx, request.Email)
		if err != nil {
			log.Println("ERROR REPO <findByEmail>:", err)
			err = helper.NewAuthError("email atau password salah")
			return
		}
		
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			log.Println("ERROR COMPARE PASSWORD:", err)
			err = helper.NewAuthError("email atau password salah")
			return
		}
	
		token, err := helper.GenerateJWT(user)
		if err != nil {
			log.Println("ERROR GENERATE TOKEN:", err)
			return
		}
		response.AccessToken = token
		return
	})
	
	return 
}

func (s *ServiceImpl) PengelolaLogin(ctx context.Context, request domain.LoginRequest) (response domain.LoginResponse, err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return
	}

	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		pengelola, err := s.PengelolaRepository.FindByEmail(ctx, tx, request.Email)
		if err != nil {
			log.Println("ERROR REPO <findByEmail>:", err)
			err = helper.NewAuthError("email atau password salah")
			return
		}
		
		err = bcrypt.CompareHashAndPassword([]byte(pengelola.Password), []byte(request.Password))
		if err != nil {
			log.Println("ERROR COMPARE PASSWORD: ", err)
			err = helper.NewAuthError("email atau password salah")
			return
		}
	
		
		token, err := helper.GeneratePengelolaJWT(pengelola)
		if err != nil {
			log.Println("ERROR GENERATE TOKEN: ", err)
			return
		}
		response.AccessToken = token
		return
	})
	
	return 
}

func (s *ServiceImpl) Logout(ctx context.Context) (err error) {
	return
}

func (s *ServiceImpl) UserChangePassword(ctx context.Context, request domain.ChangePasswordRequest) (err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return
	}
	email := ctx.Value(contextkey.UserKey).(string)
	
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.UserRepository.FindByEmail(ctx, tx, email)
		if err != nil {
			log.Println("ERROR REPO <findByEmail>:", err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.OldPassword))
		if err != nil {
			log.Println("ERROR COMPARE PASSWORD: ", err)
			err = helper.NewBadRequestError("password salah")
			return
		}

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Println("ERROR HASH PASSWORD:", err)
			return
		}
		
		result.Password = string(hashPassword)

		err = s.UserRepository.UpdatePassword(ctx, tx, &result)
		return
	})

	return
}

func (s *ServiceImpl) PengelolaChangePassword(ctx context.Context, request domain.ChangePasswordRequest) (err error) {
	err = s.Validate.Struct(request)
	if err != nil {
		log.Println("ERROR VALIDATE:", err)
		err = helper.MappingValidationError(err)
		return
	}
	email := ctx.Value(contextkey.PengelolaKey).(string)
	
	err = helper.WithTransaction(s.DB, func(tx *sql.Tx) (err error) {
		result, err := s.PengelolaRepository.FindByEmail(ctx, tx, email)
		if err != nil {
			log.Println("ERROR REPO <findByEmail>:", err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.OldPassword))
		if err != nil {
			log.Println("ERROR COMPARE PASSWORD: ", err)
			err = helper.NewBadRequestError("password salah")
			return
		}

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Println("ERROR HASH PASSWORD:", err)
			return
		}
		
		result.Password = string(hashPassword)

		err = s.PengelolaRepository.UpdatePassword(ctx, tx, &result)
		return
	})

	return
}