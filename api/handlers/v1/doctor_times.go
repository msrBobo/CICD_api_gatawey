package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models"
	"dennic_api_gateway/api/models/model_booking_service"
	pb "dennic_api_gateway/genproto/booking_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateDoctorTimes ...
// @Summary CreateDoctorTimes
// @Description CreateDoctorTimes - Api for crete doctor time
// @Tags Doctor Time
// @Accept json
// @Produce json
// @Param CreateDoctorTimeReq body model_booking_service.CreateDoctorTimeReq true "CreateDoctorTimeReq"
// @Success 200 {object} model_booking_service.DoctorTime
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-time [post]
func (h *HandlerV1) CreateDoctorTimes(c *gin.Context) {
	var (
		body        model_booking_service.CreateDoctorTimeReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateDoctorTimes") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().DoctorTimes().CreateDoctorTime(ctx, &pb.CreateDoctorTimeReq{
		DepartmentId: body.DepartmentId,
		DoctorId:     body.DoctorId,
		DoctorDate:   body.DoctorDate,
		StartTime:    body.StartTime,
		EndTime:      body.EndTime,
		Status:       body.Status,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateDoctorTimes") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.DoctorTime{
		Id:           res.Id,
		DepartmentId: res.DepartmentId,
		DoctorId:     res.DoctorId,
		DoctorDate:   res.DoctorDate,
		StartTime:    res.StartTime,
		EndTime:      res.EndTime,
		Status:       res.Status,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// GetDoctorTimes ...
// @Summary GetDoctorTimes
// @Description GetDoctorTimes - Api for get doctor time
// @Tags Doctor Time
// @Accept json
// @Produce json
// @Param GetDoctorTimeReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_booking_service.DoctorTime
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-time/get [get]
func (h *HandlerV1) GetDoctorTimes(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().DoctorTimes().GetDoctorTime(ctx, &pb.DoctorTimeFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetDoctorTimes") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.DoctorTime{
		Id:           res.Id,
		DepartmentId: res.DepartmentId,
		DoctorId:     res.DoctorId,
		DoctorDate:   res.DoctorDate,
		StartTime:    res.StartTime,
		EndTime:      res.EndTime,
		Status:       res.Status,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// ListDoctorTimes ...
// @Summary ListDoctorTimes
// @Description ListDoctorTimes - Api for list doctor time
// @Tags Doctor Time
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_booking_service.DoctorTimesType
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-time [get]
func (h *HandlerV1) ListDoctorTimes(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctorTimes") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().DoctorTimes().GetAllDoctorTimes(ctx, &pb.GetAllDoctorTimesReq{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     pageInt,
		Limit:    limitInt,
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctorTimes") {
		return
	}

	var doctorTimes model_booking_service.DoctorTimesType
	for _, times := range res.DoctorTimes {
		var doctorTime model_booking_service.DoctorTime
		doctorTime.Id = times.Id
		doctorTime.DepartmentId = times.DepartmentId
		doctorTime.DoctorId = times.DoctorId
		doctorTime.DoctorDate = times.DoctorDate
		doctorTime.StartTime = times.StartTime
		doctorTime.EndTime = times.EndTime
		doctorTime.Status = times.Status
		doctorTime.CreatedAt = times.CreatedAt
		doctorTime.UpdatedAt = e.UpdateTimeFilter(times.UpdatedAt)
		doctorTimes.DoctorTimes = append(doctorTimes.DoctorTimes, &doctorTime)
	}

	c.JSON(http.StatusOK, model_booking_service.DoctorTimesType{
		Count:       res.Count,
		DoctorTimes: doctorTimes.DoctorTimes,
	})
}

// UpdateDoctorTimes ...
// @Summary UpdateDoctorTimes
// @Description UpdateDoctorTimes - Api for update doctor time
// @Tags Doctor Time
// @Accept json
// @Produce json
// @Param doctor_time_id query string true "doctor_time_id"
// @Param UpdateDoctorTimeReq body model_booking_service.UpdateDoctorTimeReq true "UpdateDoctorTimeReq"
// @Success 200 {object} model_booking_service.DoctorTime
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-time [put]
func (h *HandlerV1) UpdateDoctorTimes(c *gin.Context) {
	id := c.Query("doctor_time_id")
	var (
		body        model_booking_service.UpdateDoctorTimeReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateDoctorTimes") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().DoctorTimes().UpdateDoctorTime(ctx, &pb.UpdateDoctorTimeReq{
		Field:        "id",
		Value:        id,
		DepartmentId: body.DepartmentId,
		DoctorId:     body.DoctorId,
		DoctorDate:   body.DoctorDate,
		StartTime:    body.StartTime,
		EndTime:      body.EndTime,
		Status:       body.Status,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateDoctorTimes") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.DoctorTime{
		Id:           res.Id,
		DepartmentId: res.DepartmentId,
		DoctorId:     res.DoctorId,
		DoctorDate:   res.DoctorDate,
		StartTime:    res.StartTime,
		EndTime:      res.EndTime,
		Status:       res.Status,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// DeleteDoctorTimes ...
// @Summary DeleteDoctorTimes
// @Description DeleteDoctorTimes - Api for delete Doctor time
// @Tags Doctor Time
// @Accept json
// @Produce json
// @Param DeleteArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-time [delete]
func (h *HandlerV1) DeleteDoctorTimes(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.BookingService().DoctorTimes().DeleteDoctorTime(ctx, &pb.DoctorTimeFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteDoctorTimes") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
