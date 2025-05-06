package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func HandleUploadImage(w http.ResponseWriter, r *http.Request, fieldName string) (fileName string, err error) {
	file, _, err := r.FormFile(fieldName)
	if err != nil {
		err = nil
		return
	}
	defer file.Close()

	uploadDir := filepath.Join(".", "uploads", "img")
	if err = os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return
	}

	mimeType, ext, err := GetMimeTypeFromMultipartForm(file)
	if err != nil {
		return
	}

	if mimeType != "image/png" && mimeType != "image/jpeg" && mimeType != "image/jpg" {
		err = fmt.Errorf("unsupported file type: %s", ext)
		return
	}

	uuid := uuid.NewString()
	fileName = fmt.Sprintf("%s.%s", uuid, ext)
	destPath := filepath.Join(uploadDir, fileName)
	destFile, err := os.Create(destPath)
	if err != nil {
		return
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, file); err != nil {
		return
	}

	return
}

func HandleUploadPdf(w http.ResponseWriter, r *http.Request, fieldName string) (fileName string, err error) {
	file, _, err := r.FormFile(fieldName)
	if err != nil {
		err = nil
		return
	}
	defer file.Close()

	uploadDir := filepath.Join(".", "uploads", "docs")
	if err = os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return
	}

	mimeType, ext, err := GetMimeTypeFromMultipartForm(file)
	if err != nil {
		return
	}

	if mimeType != "application/pdf" {
		err = fmt.Errorf("unsupported file type: %s", ext)
		return
	}

	uuid := uuid.NewString()
	fileName = fmt.Sprintf("%s.%s", uuid, ext)
	destPath := filepath.Join(uploadDir, fileName)
	destFile, err := os.Create(destPath)
	if err != nil {
		return
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, file); err != nil {
		return
	}

	return
}

func GetMimeTypeFromMultipartForm(file multipart.File) (mimeType string, ext string, err error) {
	buffer := make([]byte, 512)
	if _, err = file.Read(buffer); err != nil {
		return
	}

	mimeType = http.DetectContentType(buffer)

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	ext = strings.Split(mimeType, "/")[1]

	return
}

func CheckFormFile(r *http.Request, fieldName string) error {
	file, _, err := r.FormFile(fieldName)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}