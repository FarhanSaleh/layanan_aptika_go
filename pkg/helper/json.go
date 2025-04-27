package helper

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/constants"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/go-sql-driver/mysql"
)

func WriteResponseBody(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("ERROR ENCODE: ", err)
		panic(err)
	}
}

func ParseBody(r *http.Request, data any){
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Println("ERROR DECODE: ", err)
		panic(err)
	}
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	// validation error check
	if validationErr, ok := IsValidationError(err); ok {
		WriteResponseBody(w, http.StatusBadRequest, domain.ErrorValidationResponse{
			Message: validationErr.Error(),
			Errors: validationErr.Errors,
		})
		return
	}
	
	// mysql error check
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			WriteResponseBody(w, http.StatusConflict, domain.DefaultResponse{
				Message: mysqlError.Message,
			})
			return
		}
		
		WriteResponseBody(w, http.StatusBadRequest, domain.DefaultResponse{
			Message: mysqlError.Message,
		})
		return
	}

	// auth error check
	if authError, ok := err.(*AuthError); ok {
		WriteResponseBody(w, http.StatusUnauthorized, domain.DefaultResponse{
			Message: authError.Message,
		})
		return
	}

	// error no rows
	if errors.Is(err, sql.ErrNoRows){
		WriteResponseBody(w, http.StatusNotFound, domain.DefaultResponse{
			Message: constants.ErrorNotFound,
		})
		return
	}

	WriteResponseBody(w, http.StatusInternalServerError, domain.DefaultResponse{
		Message: err.Error(),
	})
}