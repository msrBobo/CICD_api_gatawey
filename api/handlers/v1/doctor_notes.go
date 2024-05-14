package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models"
	"dennic_api_gateway/api/models/model_booking_service"
	pb "dennic_api_gateway/genproto/booking_service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateDoctorNote ...
// @Summary CreateDoctorNote
// @Description CreateDoctorNote - Api for create doctor note
// @Tags Doctor Note
// @Accept json
// @Produce json
// @Param CreateDoctorNotesReq body model_booking_service.CreateDoctorNotesReq true "CreateDoctorNotesReq"
// @Success 200 {object} model_booking_service.DoctorNote
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-notes [post]
func (h *HandlerV1) CreateDoctorNote(c *gin.Context) {
	var (
		body        model_booking_service.CreateDoctorNotesReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateDoctorNote") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorNote, err := h.serviceManager.BookingService().DoctorNotes().CreateDoctorNote(ctx, &pb.CreateDoctorNoteReq{
		AppointmentId: body.AppointmentId,
		DoctorId:      body.DoctorId,
		PatientId:     body.PatientId,
		Prescription:  body.Prescription,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateDoctorNote") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.DoctorNote{
		Id:            doctorNote.Id,
		AppointmentId: doctorNote.AppointmentId,
		DoctorId:      doctorNote.DoctorId,
		PatientId:     doctorNote.PatientId,
		Prescription:  doctorNote.Prescription,
		CreatedAt:     doctorNote.CreatedAt,
		UpdatedAt:     e.UpdateTimeFilter(doctorNote.UpdatedAt),
	})
}

// GetDoctorNote ...
// @Summary GetDoctorNote
// @Description GetDoctorNote - API to get doctor note by ID
// @Tags Doctor Note
// @Accept json
// @Produce json
// @Param GetArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_booking_service.DoctorNote
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-notes/get [get]
func (h *HandlerV1) GetDoctorNote(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorNote, err := h.serviceManager.BookingService().DoctorNotes().GetDoctorNote(ctx, &pb.FieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetDoctorNote") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.DoctorNote{
		Id:            doctorNote.Id,
		AppointmentId: doctorNote.AppointmentId,
		DoctorId:      doctorNote.DoctorId,
		PatientId:     doctorNote.PatientId,
		Prescription:  doctorNote.Prescription,
		CreatedAt:     doctorNote.CreatedAt,
		UpdatedAt:     e.UpdateTimeFilter(doctorNote.UpdatedAt),
	})
}

// ListDoctorNotes ...
// @Summary ListDoctorNotes
// @Description ListDoctorNotes - API to list doctor notes
// @Tags Doctor Note
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_booking_service.DoctorNotesType
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-notes [get]
func (h *HandlerV1) ListDoctorNotes(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctorNotes") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorNotes, err := h.serviceManager.BookingService().DoctorNotes().GetAllNotes(ctx, &pb.GetAllReq{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     pageInt,
		Limit:    limitInt,
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctorNotes") {
		return
	}

	var doctorNotesRes model_booking_service.DoctorNotesType
	for _, doctorNoteRes := range doctorNotes.DoctorNotes {
		var doctorNote model_booking_service.DoctorNote
		doctorNote.Id = doctorNoteRes.Id
		doctorNote.AppointmentId = doctorNoteRes.AppointmentId
		doctorNote.DoctorId = doctorNoteRes.DoctorId
		doctorNote.PatientId = doctorNoteRes.PatientId
		doctorNote.Prescription = doctorNoteRes.Prescription
		doctorNote.CreatedAt = doctorNoteRes.CreatedAt
		doctorNote.UpdatedAt = e.UpdateTimeFilter(doctorNoteRes.UpdatedAt)
		doctorNotesRes.DoctorNotes = append(doctorNotesRes.DoctorNotes, &doctorNote)
	}

	c.JSON(http.StatusOK, doctorNotesRes)
}

// UpdateDoctorNote ...
// @Summary UpdateDoctorNote
// @Description UpdateDoctorNote - API to update a doctor note
// @Tags Doctor Note
// @Accept json
// @Produce json
// @Param notes_id query string true "notes_id"
// @Param UpdateDoctorNoteReq body model_booking_service.UpdateDoctorNoteReq true "UpdateDoctorNoteReq"
// @Success 200 {object} model_booking_service.DoctorNote
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-notes [put]
func (h *HandlerV1) UpdateDoctorNote(c *gin.Context) {
	id := c.Query("notes_id")
	var (
		body        model_booking_service.UpdateDoctorNoteReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateDoctorNote") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorNote, err := h.serviceManager.BookingService().DoctorNotes().UpdateDoctorNote(ctx, &pb.UpdateDoctorNoteReq{
		Field:         "id",
		Value:         id,
		AppointmentId: body.AppointmentId,
		DoctorId:      body.DoctorId,
		PatientId:     body.PatientId,
		Prescription:  body.Prescription,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateDoctorNote") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.DoctorNote{
		Id:            doctorNote.Id,
		AppointmentId: doctorNote.AppointmentId,
		DoctorId:      doctorNote.DoctorId,
		PatientId:     doctorNote.PatientId,
		Prescription:  doctorNote.Prescription,
		CreatedAt:     doctorNote.CreatedAt,
		UpdatedAt:     e.UpdateTimeFilter(doctorNote.UpdatedAt),
	})
}

// DeleteDoctorNote ...
// @Summary DeleteDoctorNote
// @Description DeleteDoctorNote - API to delete a doctor note
// @Tags Doctor Note
// @Accept json
// @Produce json
// @Param DeleteArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-notes [delete]
func (h *HandlerV1) DeleteDoctorNote(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.BookingService().DoctorNotes().DeleteDoctorNote(ctx, &pb.FieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteDoctorNote") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
