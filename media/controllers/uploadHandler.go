package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	Config "github.com/nileshnk/media_upload_poc/media/config"
	Services "github.com/nileshnk/media_upload_poc/media/services"
	Utils "github.com/nileshnk/media_upload_poc/media/utils"
)

func UploadHandlerfunc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()

	r.ParseForm()
	file, fileHeader, err := r.FormFile("file")
	Utils.CheckErrorWithCommonResponse(err)
	fileBytes := make([]byte, fileHeader.Size)
	n, errRead := file.Read(fileBytes)
	Utils.CheckErrorWithCommonResponse(errRead)
	if n != int(fileHeader.Size) {
		fmt.Println(n)
	}

	cwd, errCwd := os.Getwd()
	log.Println(errCwd)
	Utils.CheckErrorWithCommonResponse(errors.New(Config.HttpInternalServerErrorMessage))

	newFile, errCreate := os.Create(cwd + Config.GetConfig.Media.Directory + fileHeader.Filename)
	log.Println(errCreate)
	Utils.CheckErrorWithCommonResponse(errors.New(Config.HttpInternalServerErrorMessage))
	defer newFile.Close()

	errUpload := Services.S3Upload(ctx, file, fileHeader.Filename)
	log.Println(errUpload)
	Utils.CheckErrorWithCommonResponse(errors.New(Config.HttpInternalServerErrorMessage))

	_, errWrite := newFile.Write(fileBytes)
	log.Println(errWrite)
	Utils.CheckErrorWithCommonResponse(errors.New(Config.HttpInternalServerErrorMessage))

	w.Write([]byte("File Uploaded"))
}
