package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models"
	"dennic_api_gateway/api/models/model_healthcare_service"
	pb "dennic_api_gateway/genproto/healthcare-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateReasons ...
// @Summary CreateReasons
// @Description CreateReasons - Api for crete reasons
// @Tags Reasons
// @Accept json
// @Produce json
// @Param DoctorWorkingHoursReq body model_healthcare_service.ReasonsReq true "DoctorServiceReq"
// @Success 200 {object} model_healthcare_service.ReasonsRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/reasons [post]
func (h *HandlerV1) CreateReasons(c *gin.Context) {
	var (
		body        model_healthcare_service.ReasonsReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateReasons") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	reasons, err := h.serviceManager.HealthcareService().ReasonsService().CreateReasons(ctx, &pb.Reasons{
		Id:               uuid.NewString(),
		Name:             body.Name,
		SpecializationId: body.SpecializationId,
		ImageUrl:         body.ImageUrl,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateReasons") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.ReasonsRes{
		Id:               reasons.Id,
		Name:             reasons.Name,
		SpecializationId: reasons.SpecializationId,
		ImageUrl:         reasons.ImageUrl,
		CreatedAt:        reasons.CreatedAt,
		UpdatedAt:        e.UpdateTimeFilter(reasons.UpdatedAt),
	})
}

// GetReasons ...
// @Summary GetReasons
// @Description GetReasons - Api for get reasons
// @Tags Reasons
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} model_healthcare_service.ReasonsRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/reasons/get [get]
func (h *HandlerV1) GetReasons(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	reasons, err := h.serviceManager.HealthcareService().ReasonsService().GetReasonsById(ctx, &pb.GetReqStrReasons{
		Field:    "id",
		Value:    id,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetReasons") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.ReasonsRes{
		Id:               reasons.Id,
		Name:             reasons.Name,
		SpecializationId: reasons.SpecializationId,
		ImageUrl:         reasons.ImageUrl,
		CreatedAt:        reasons.CreatedAt,
		UpdatedAt:        e.UpdateTimeFilter(reasons.UpdatedAt),
	})
}

// ListReasons ...
// @Summary ListReasons
// @Description ListReasons - Api for list reasons
// @Tags Reasons
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Param search query string false "search" Enums(name) "search"
// @Success 200 {object} model_healthcare_service.ListReasons
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/reasons [get]
func (h *HandlerV1) ListReasons(c *gin.Context) {
	search := c.Query("search")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListReasons") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	reasons, err := h.serviceManager.HealthcareService().ReasonsService().GetAllReasons(ctx, &pb.GetAllReas{
		Field:    search,
		Value:    value,
		IsActive: false,
		Page:     int32(pageInt),
		Limit:    int32(limitInt),
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListReasons") {
		return
	}
	var reasonsRes model_healthcare_service.ListReasons
	for _, repreasons := range reasons.Reasons {
		reasonsRes.Reasons = append(reasonsRes.Reasons, &model_healthcare_service.ReasonsRes{
			Id:               repreasons.Id,
			Name:             repreasons.Name,
			SpecializationId: repreasons.SpecializationId,
			ImageUrl:         repreasons.ImageUrl,
			CreatedAt:        repreasons.CreatedAt,
			UpdatedAt:        e.UpdateTimeFilter(repreasons.UpdatedAt),
		})
	}

	c.JSON(http.StatusOK, model_healthcare_service.ListReasons{
		Count:   reasons.Count,
		Reasons: reasonsRes.Reasons,
	})
}

// UpdateReasons ...
// @Summary UpdateReasons
// @Description UpdateReasons - Api for update reasons
// @Tags Reasons
// @Accept json
// @Produce json
// @Param UpdateReasonsReq body model_healthcare_service.ReasonsReq true "UpdateReasonsReq"
// @Success 200 {object} model_healthcare_service.ReasonsRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/reasons [put]
func (h *HandlerV1) UpdateReasons(c *gin.Context) {
	var (
		body        model_healthcare_service.ReasonsReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateReasons") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	respreasons, err := h.serviceManager.HealthcareService().ReasonsService().UpdateReasons(ctx, &pb.Reasons{
		Id:               body.Id,
		Name:             body.Name,
		SpecializationId: body.SpecializationId,
		ImageUrl:         body.ImageUrl,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateReasons") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.ReasonsRes{
		Id:               respreasons.Id,
		Name:             respreasons.Name,
		SpecializationId: respreasons.SpecializationId,
		ImageUrl:         respreasons.ImageUrl,
		CreatedAt:        respreasons.CreatedAt,
		UpdatedAt:        e.UpdateTimeFilter(respreasons.UpdatedAt),
	})
}

// DeleteReasons ...
// @Summary DeleteReasons
// @Description DeleteReasons - Api for delete reasons
// @Tags Reasons
// @Accept json
// @Produce json
// @Param DeleteReasonsReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/reasons [delete]
func (h *HandlerV1) DeleteReasons(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.HealthcareService().ReasonsService().DeleteReasons(ctx, &pb.GetReqStrReasons{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteReasons") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
