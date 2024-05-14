package v1

import (
	e "dennic_api_gateway/api/handlers/regtool"
	m "dennic_api_gateway/api/models/model_minio"
	"dennic_api_gateway/internal/pkg/minio"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"path/filepath"
)

// UploadFile ...
// @Summary Upload image
// @Description Upload image
// @Tags upload-file
// @Accept image/png
// @Produce json
// @Param file formData file true "file"
// @Param bucketName query string true "bucket" Enums(department, reasons, specialization, doctor, user) "bucket name"
// @Success 200 {object} model_minio.MinioURL
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/file-upload [post]
func (h *HandlerV1) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UploadFile") {
		return
	}

	bucketName := c.Query("bucketName")

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UploadFile") {
		return
	}

	generatedFileName := uuid.New().String() + filepath.Ext(header.Filename)

<<<<<<< HEAD
	objectURL, err := minio.UploadToMinio(h.cfg, generatedFileName, fileBytes, int64(len(fileBytes)))
=======
	objectURL, err := minio.UploadToMinio(h.cfg, id, fileBytes, bucketName)
>>>>>>> eca93a4806164b242209e3dd00f4c94991ddf34d

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UploadFile") {
		return
	}

	c.JSON(http.StatusCreated, m.MinioURL{
		URL: objectURL,
	})
}
