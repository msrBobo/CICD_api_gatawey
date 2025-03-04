package v1

import (
	e "CICD_api_gatawey/api/handlers/regtool"
	"CICD_api_gatawey/api/models"
	"CICD_api_gatawey/api/models/model_booking_service"
	pb "CICD_api_gatawey/genproto/booking_service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateArchive ...
// @Summary CreateArchive
// @Description CreateArchive - Api for crete archive
// @Tags Archive
// @Accept json
// @Produce json
// @Param CreateArchiveReq body model_booking_service.CreateArchiveReq true "CreateArchiveReq"
// @Success 200 {object} model_booking_service.Archive
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/archive [post]
func (h *HandlerV1) CreateArchive(c *gin.Context) {
	var (
		body        model_booking_service.CreateArchiveReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateArchive") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	archive, err := h.serviceManager.BookingService().ArchiveService().CreateArchive(ctx, &pb.CreateArchiveReq{
		DoctorAvailabilityId: body.DoctorAvailabilityId,
		StartTime:            body.StartTime,
		EndTime:              body.EndTime,
		PatientProblem:       body.PatientProblem,
		Status:               body.Status,
		PaymentType:          body.PaymentType,
		PaymentAmount:        float32(body.PaymentAmount),
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateArchive") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Archive{
		Id:                   archive.Id,
		DoctorAvailabilityId: archive.DoctorAvailabilityId,
		StartTime:            archive.StartTime,
		EndTime:              archive.EndTime,
		PatientProblem:       archive.PatientProblem,
		Status:               archive.Status,
		PaymentType:          archive.PaymentType,
		PaymentAmount:        float64(archive.PaymentAmount),
		CreatedAt:            archive.CreatedAt,
		UpdatedAt:            e.UpdateTimeFilter(archive.UpdatedAt),
	})
}

// GetArchive ...
// @Summary GetArchive
// @Description GetArchive - Api for get archive
// @Tags Archive
// @Accept json
// @Produce json
// @Param id query integer true "id"
// @Success 200 {object} model_booking_service.Archive
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/archive/get [get]
func (h *HandlerV1) GetArchive(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	archive, err := h.serviceManager.BookingService().ArchiveService().GetArchive(ctx, &pb.ArchiveFieldValueReq{
		Field:    "id",
		Value:    id,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetArchive") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Archive{
		Id:                   archive.Id,
		DoctorAvailabilityId: archive.DoctorAvailabilityId,
		StartTime:            archive.StartTime,
		EndTime:              archive.EndTime,
		PatientProblem:       archive.PatientProblem,
		Status:               archive.Status,
		PaymentType:          archive.PaymentType,
		PaymentAmount:        float64(archive.PaymentAmount),
		CreatedAt:            archive.CreatedAt,
		UpdatedAt:            e.UpdateTimeFilter(archive.UpdatedAt),
	})
}

// ListArchive ...
// @Summary ListArchive
// @Description ListArchive - Api for list archive
// @Tags Archive
// @Accept json
// @Produce json
// @Param searchField query string false "searchField" Enums(status)
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_booking_service.ArchivesType
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/archive [get]
func (h *HandlerV1) ListArchive(c *gin.Context) {
	field := c.Query("searchField")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListArchive") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	archives, err := h.serviceManager.BookingService().ArchiveService().GetAllArchives(ctx, &pb.GetAllArchivesReq{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     pageInt,
		Limit:    limitInt,
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListArchive") {
		return
	}

	var archivesRes model_booking_service.ArchivesType
	for _, archiveRes := range archives.Archives {
		var archive model_booking_service.Archive
		archive.Id = archiveRes.Id
		archive.DoctorAvailabilityId = archiveRes.DoctorAvailabilityId
		archive.StartTime = archiveRes.StartTime
		archive.EndTime = archiveRes.EndTime
		archive.PatientProblem = archiveRes.PatientProblem
		archive.Status = archiveRes.Status
		archive.PaymentType = archiveRes.PaymentType
		archive.PaymentAmount = float64(archiveRes.PaymentAmount)
		archive.CreatedAt = archiveRes.CreatedAt
		archive.UpdatedAt = e.UpdateTimeFilter(archiveRes.UpdatedAt)
		archivesRes.Archives = append(archivesRes.Archives, &archive)
	}

	c.JSON(http.StatusOK, model_booking_service.ArchivesType{
		Count:    archives.Count,
		Archives: archivesRes.Archives,
	})
}

// UpdateArchive ...
// @Summary UpdateArchive
// @Description UpdateArchive - Api for update archive
// @Tags Archive
// @Accept json
// @Produce json
// @Param UpdateArchiveReq body model_booking_service.UpdateArchiveReq true "UpdateArchiveReq"
// @Success 200 {object} model_booking_service.Archive
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/archive [put]
func (h *HandlerV1) UpdateArchive(c *gin.Context) {
	var (
		body        model_booking_service.UpdateArchiveReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateArchive") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	archive, err := h.serviceManager.BookingService().ArchiveService().UpdateArchive(ctx, &pb.UpdateArchiveReq{
		Field:                "id",
		Value:                body.ArchiveId,
		DoctorAvailabilityId: body.DoctorAvailabilityId,
		StartTime:            body.StartTime,
		EndTime:              body.EndTime,
		PatientProblem:       body.PatientProblem,
		Status:               body.Status,
		PaymentType:          body.PaymentType,
		PaymentAmount:        float32(body.PaymentAmount),
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateArchive") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Archive{
		Id:                   archive.Id,
		DoctorAvailabilityId: archive.DoctorAvailabilityId,
		StartTime:            archive.StartTime,
		EndTime:              archive.EndTime,
		PatientProblem:       archive.PatientProblem,
		Status:               archive.Status,
		PaymentType:          archive.PaymentType,
		PaymentAmount:        float64(archive.PaymentAmount),
		CreatedAt:            archive.CreatedAt,
		UpdatedAt:            e.UpdateTimeFilter(archive.UpdatedAt),
	})
}

// DeleteArchive ...
// @Summary DeleteArchive
// @Description DeleteArchive - Api for delete archive
// @Tags Archive
// @Accept json
// @Produce json
// @Param DeleteArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/archive [delete]
func (h *HandlerV1) DeleteArchive(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.BookingService().ArchiveService().DeleteArchive(ctx, &pb.ArchiveFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteArchive") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
