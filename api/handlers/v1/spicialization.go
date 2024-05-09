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

// CreateSpecialization ...
// @Summary CreateSpecialization
// @Description CreateSpecialization - Api for crete specialization
// @Tags Specialization
// @Accept json
// @Produce json
// @Param SpecializationReq body model_healthcare_service.SpecializationReq true "SpecializationReq"
// @Success 200 {object} model_healthcare_service.SpecializationRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/specialization [post]
func (h *HandlerV1) CreateSpecialization(c *gin.Context) {
	var (
		body        model_healthcare_service.SpecializationReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateSpecialization") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	specialization, err := h.serviceManager.HealthcareService().SpecializationService().CreateSpecialization(ctx, &pb.Specializations{
		Id:           uuid.NewString(),
		Name:         body.Name,
		Description:  body.Description,
		DepartmentId: body.DepartmentId,
		ImageUrl:     body.ImageUrl,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateSpecialization") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.SpecializationRes{
		ID:           specialization.Id,
		Order:        specialization.Order,
		Name:         specialization.Name,
		Description:  specialization.Description,
		DepartmentId: specialization.DepartmentId,
		ImageUrl:     specialization.ImageUrl,
		CreatedAt:    specialization.CreatedAt,
		UpdatedAt:    specialization.UpdatedAt,
	})
}

// GetSpecialization ...
// @Summary GetSpecialization
// @Description GetSpecialization - Api for get specialization
// @Tags Specialization
// @Accept json
// @Produce json
// @Param GetSpecialization query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_healthcare_service.SpecializationRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/specialization [get]
func (h *HandlerV1) GetSpecialization(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	specialization, err := h.serviceManager.HealthcareService().SpecializationService().GetSpecializationById(ctx, &pb.GetReqStrSpecialization{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetSpecialization") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.SpecializationRes{
		ID:           specialization.Id,
		Order:        specialization.Order,
		Name:         specialization.Name,
		Description:  specialization.Description,
		DepartmentId: specialization.DepartmentId,
		ImageUrl:     specialization.ImageUrl,
		CreatedAt:    specialization.CreatedAt,
		UpdatedAt:    specialization.UpdatedAt,
	})
}

// ListSpecializations ...
// @Summary ListSpecializations
// @Description ListSpecializations - Api for list specialization
// @Tags Specialization
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_healthcare_service.ListSpecializations
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/specialization/get [get]
func (h *HandlerV1) ListSpecializations(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListSpecializations") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	specializations, err := h.serviceManager.HealthcareService().SpecializationService().GetAllSpecializations(ctx, &pb.GetAllSpecialization{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     int32(pageInt),
		Limit:    int32(limitInt),
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListSpecializations") {
		return
	}

	var specializationsRes model_healthcare_service.ListSpecializations
	for _, specializationRes := range specializations.Specializations {
		specializationsRes.Specializations = append(specializationsRes.Specializations, &model_healthcare_service.SpecializationRes{
			ID:           specializationRes.Id,
			Order:        specializationRes.Order,
			Name:         specializationRes.Name,
			Description:  specializationRes.Description,
			DepartmentId: specializationRes.DepartmentId,
			ImageUrl:     specializationRes.ImageUrl,
			CreatedAt:    specializationRes.CreatedAt,
			UpdatedAt:    specializationRes.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, model_healthcare_service.ListSpecializations{
		Count:           specializations.Count,
		Specializations: specializationsRes.Specializations,
	})
}

// UpdateSpecialization ...
// @Summary UpdateSpecialization
// @Description UpdateSpecialization - Api for update specialization
// @Tags Specialization
// @Accept json
// @Produce json
// @Param UpdateSpecializationReq body model_healthcare_service.SpecializationReq true "UpdateSpecializationReq"
// @Param id query string true "id"
// @Success 200 {object} model_healthcare_service.SpecializationRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/specialization [put]
func (h *HandlerV1) UpdateSpecialization(c *gin.Context) {
	var (
		body        model_healthcare_service.SpecializationReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	id := c.Query("id")

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateSpecialization") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	specialization, err := h.serviceManager.HealthcareService().SpecializationService().UpdateSpecialization(ctx, &pb.Specializations{
		Id:           id,
		Name:         body.Name,
		Description:  body.Description,
		DepartmentId: body.DepartmentId,
		ImageUrl:     body.ImageUrl,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateSpecialization") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.SpecializationRes{
		ID:           specialization.Id,
		Order:        specialization.Order,
		Name:         specialization.Name,
		Description:  specialization.Description,
		DepartmentId: specialization.DepartmentId,
		ImageUrl:     specialization.ImageUrl,
		CreatedAt:    specialization.CreatedAt,
		UpdatedAt:    specialization.UpdatedAt,
	})
}

// DeleteSpecialization ...
// @Summary DeleteSpecialization
// @Description DeleteSpecialization - Api for delete specialization
// @Tags Specialization
// @Accept json
// @Produce json
// @Param DeleteSpecializationReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/specialization [delete]
func (h *HandlerV1) DeleteSpecialization(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.HealthcareService().SpecializationService().DeleteSpecialization(ctx, &pb.GetReqStrSpecialization{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteSpecialization") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
