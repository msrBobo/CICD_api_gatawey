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

// CreateBookedAppointment ...
// @Summary CreateBookedAppointment
// @Description CreateBookedAppointment - Api for create booked appointment
// @Tags Appointment
// @Accept json
// @Produce json
// @Param CreateAppointmentReq body model_booking_service.CreateAppointmentReq true "CreateAppointmentReq"
// @Success 200 {object} model_booking_service.Appointment
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/appointment [post]
func (h *HandlerV1) CreateBookedAppointment(c *gin.Context) {
	var (
		body        model_booking_service.CreateAppointmentReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateBookedAppointment") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().BookedAppointment().CreateAppointment(ctx, &pb.CreateAppointmentReq{
		DepartmentId:    body.DepartmentId,
		DoctorId:        body.DoctorId,
		PatientId:       body.PatientId,
		AppointmentDate: body.AppointmentDate,
		AppointmentTime: body.AppointmentTime,
		Duration:        body.Duration,
		Key:             body.Key,
		ExpiresAt:       body.ExpiresAt,
		PatientStatus:   false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateBookedAppointment") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Appointment{
		Id:              res.Id,
		DepartmentId:    res.DepartmentId,
		DoctorId:        res.DoctorId,
		PatientId:       res.PatientId,
		AppointmentDate: res.AppointmentDate,
		AppointmentTime: res.AppointmentTime,
		Duration:        res.Duration,
		Key:             res.Key,
		ExpiresAt:       res.ExpiresAt,
		PatientStatus:   res.PatientStatus,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// GetBookedAppointment ...
// @Summary GetBookedAppointment
// @Description GetBookedAppointment - API to get Booked appointment by ID
// @Tags Appointment
// @Accept json
// @Produce json
// @Param GetArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_booking_service.Appointment
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/appointment/get [get]
func (h *HandlerV1) GetBookedAppointment(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().BookedAppointment().GetAppointment(ctx, &pb.AppointmentFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetBookedAppointment") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Appointment{
		Id:              res.Id,
		DepartmentId:    res.DepartmentId,
		DoctorId:        res.DoctorId,
		PatientId:       res.PatientId,
		AppointmentDate: res.AppointmentDate,
		AppointmentTime: res.AppointmentTime,
		Duration:        res.Duration,
		Key:             res.Key,
		ExpiresAt:       res.ExpiresAt,
		PatientStatus:   res.PatientStatus,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// ListBookedAppointments ...
// @Summary ListBookedAppointments
// @Description ListBookedAppointments - API to list doctor notes
// @Tags Appointment
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_booking_service.Appointment
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/appointment [get]
func (h *HandlerV1) ListBookedAppointments(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListBookedAppointments") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().BookedAppointment().GetAllAppointment(ctx, &pb.GetAllAppointmentsReq{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     pageInt,
		Limit:    limitInt,
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListBookedAppointments") {
		return
	}

	var response model_booking_service.AppointmentsType
	for _, appointment := range res.Appointments {
		var app model_booking_service.Appointment
		app.Id = appointment.Id
		app.DepartmentId = appointment.DepartmentId
		app.DoctorId = appointment.DoctorId
		app.PatientId = appointment.PatientId
		app.AppointmentDate = appointment.AppointmentDate
		app.AppointmentTime = appointment.AppointmentTime
		app.Duration = appointment.Duration
		app.Key = appointment.Key
		app.ExpiresAt = appointment.ExpiresAt
		app.PatientStatus = appointment.PatientStatus
		app.CreatedAt = appointment.CreatedAt
		app.UpdatedAt = e.UpdateTimeFilter(appointment.UpdatedAt)
		response.Appointments = append(response.Appointments, &app)
	}

	c.JSON(http.StatusOK, &model_booking_service.AppointmentsType{
		Appointments: response.Appointments,
		Count:        res.Count,
	})
}

// UpdateBookedAppointment ...
// @Summary UpdateBookedAppointment
// @Description UpdateDoctorNote - API to update appointment
// @Tags Appointment
// @Accept json
// @Produce json
// @Param appointment_id  query string true "appointment_id"
// @Param UpdateAppointmentReq body model_booking_service.UpdateAppointmentReq true "UpdateAppointmentReq"
// @Success 200 {object} model_booking_service.Appointment
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/appointment [put]
func (h *HandlerV1) UpdateBookedAppointment(c *gin.Context) {
	id := c.Query("appointment_id")
	var (
		body        model_booking_service.UpdateAppointmentReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateBookedAppointment") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().BookedAppointment().UpdateAppointment(ctx, &pb.UpdateAppointmentReq{
		AppointmentDate: body.AppointmentDate,
		AppointmentTime: body.AppointmentTime,
		Duration:        body.Duration,
		Key:             body.Key,
		ExpiresAt:       body.ExpiresAt,
		PatientStatus:   body.PatientStatus,
		Field:           "id",
		Value:           id,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateBookedAppointment") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Appointment{
		Id:              res.Id,
		DepartmentId:    res.DepartmentId,
		DoctorId:        res.DoctorId,
		PatientId:       res.PatientId,
		AppointmentDate: res.AppointmentDate,
		AppointmentTime: res.AppointmentTime,
		Duration:        res.Duration,
		Key:             res.Key,
		ExpiresAt:       res.ExpiresAt,
		PatientStatus:   res.PatientStatus,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// DeleteBookedAppointment ...
// @Summary DeleteBookedAppointment
// @Description DeleteBookedAppointment - API to delete a appointment
// @Tags Appointment
// @Accept json
// @Produce json
// @Param DeleteArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/appointment [delete]
func (h *HandlerV1) DeleteBookedAppointment(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.BookingService().BookedAppointment().DeleteAppointment(ctx, &pb.AppointmentFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteBookedAppointment") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
