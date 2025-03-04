package v1

import (
	e "CICD_api_gatawey/api/handlers/regtool"
	m "CICD_api_gatawey/api/models/model_minio"
	"CICD_api_gatawey/internal/pkg/minio"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadFile ...
// @Summary Upload image
// @Description Upload image
// @Tags upload-file
// @Accept image/png
// @Produce json
// @Param file formData file true "file"
// @Param bucketName query string false "bucket" Enums(department, reasons, specialization, doctor, user) "bucket name"
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

	objectURL, err := minio.UploadToMinio(h.cfg, generatedFileName, fileBytes, bucketName)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UploadFile") {
		return
	}

	c.JSON(http.StatusCreated, m.MinioURL{
		URL: objectURL,
	})
}
