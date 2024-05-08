package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models/model_booking_service"
	pb "dennic_api_gateway/genproto/booking_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateArchive ...
// @Summary CreateArchive
// @Description CreateArchive - Api for crete archive
// @Tags Booking Archive
// @Accept json
// @Produce json
// @Param CreateArchiveReq body model_booking_service.CreateArchiveReq true "CreateArchiveReq"
// @Success 200 {object} model_booking_service.Archive
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/booking/archive [post]
func (h *HandlerV1) CreateArchive(c *gin.Context) {
	var (
		body        model_booking_service.CreateArchiveReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "Register") {
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
	})
}

// GetArchive ...
// @Summary GetArchive
// @Description GetArchive - Api for get archive
// @Tags Booking Archive
// @Accept json
// @Produce json
// @Param GetArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_booking_service.Archive
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/booking/archive [get]
func (h *HandlerV1) GetArchive(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	archive, err := h.serviceManager.BookingService().ArchiveService().GetArchive(ctx, &pb.ArchiveFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
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
	})
}

// ListArchive ...
// @Summary ListArchive
// @Description ListArchive - Api for list archive
// @Tags Booking Archive
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_booking_service.ArchivesType
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/booking/archive/get [get]
func (h *HandlerV1) ListArchive(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	archives, err := h.serviceManager.BookingService().ArchiveService().GetAllArchives(ctx, &pb.GetAllArchivesReq{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     ,
		Limit:    0,
		OrderBy:  "",
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateArchive") {
		return
	}

	for _, archiveRes := range archives {
		var archive model_booking_service.Archive
		archive.Id = archiveRes

	}

	c.JSON(http.StatusOK, model_booking_service.ArchivesType{
		Count:    0,
		Archives: nil,
	})
}

// UpdateArchive ...
// @Summary UpdateArchive
// @Description UpdateArchive - Api for update archive
// @Tags Booking Archive
// @Accept json
// @Produce json
// @Param UpdateArchiveReq body model_booking_service.UpdateArchiveReq true "UpdateArchiveReq"
// @Success 200 {object} model_booking_service.Archive
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/booking/archive [put]
func (h *HandlerV1) UpdateArchive(c *gin.Context) {
	var (
		body        model_booking_service.UpdateArchiveReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "Register") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	archive, err := h.serviceManager.BookingService().ArchiveService().UpdateArchive(ctx, &pb.UpdateArchiveReq{
		Field:                body.Field,
		Value:                body.Value,
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
	})
}
