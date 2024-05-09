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

// CreateDepartment ...
// @Summary CreateDepartment
// @Description CreateDepartment - Api for crete department
// @Tags Department
// @Accept json
// @Produce json
// @Param DepartmentReq body model_healthcare_service.DepartmentReq true "DepartmentReq"
// @Success 200 {object} model_healthcare_service.DepartmentRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/department [post]
func (h *HandlerV1) CreateDepartment(c *gin.Context) {
	var (
		body        model_healthcare_service.DepartmentReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateDepartment") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	department, err := h.serviceManager.HealthcareService().DepartmentService().CreateDepartment(ctx, &pb.Department{
		Id:               uuid.NewString(),
		Name:             body.Name,
		Description:      body.Description,
		ImageUrl:         body.ImageUrl,
		FloorNumber:      body.FloorNumber,
		ShortDescription: body.ShortDescription,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateDepartment") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DepartmentRes{
		Id:               department.Id,
		Name:             department.Name,
		Description:      department.Description,
		ImageUrl:         department.ImageUrl,
		FloorNumber:      department.FloorNumber,
		ShortDescription: department.ShortDescription,
		CreatedAt:        department.CreatedAt,
		UpdatedAt:        department.UpdatedAt,
	})
}

// GetDepartment ...
// @Summary GetDepartment
// @Description GetDepartment - Api for get department
// @Tags Department
// @Accept json
// @Produce json
// @Param GetDepartment query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_healthcare_service.DepartmentRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/department [get]
func (h *HandlerV1) GetDepartment(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	department, err := h.serviceManager.HealthcareService().DepartmentService().GetDepartmentById(ctx, &pb.GetReqStrDepartment{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetDepartment") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DepartmentRes{
		Id:               department.Id,
		Name:             department.Name,
		Description:      department.Description,
		ImageUrl:         department.ImageUrl,
		FloorNumber:      department.FloorNumber,
		ShortDescription: department.ShortDescription,
		CreatedAt:        department.CreatedAt,
		UpdatedAt:        department.UpdatedAt,
	})
}

// ListDepartments ...
// @Summary ListDepartments
// @Description ListDepartments - Api for list department
// @Tags Department
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_healthcare_service.ListDepartments
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/department/get [get]
func (h *HandlerV1) ListDepartments(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDepartments") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	departments, err := h.serviceManager.HealthcareService().DepartmentService().GetAllDepartments(ctx, &pb.GetAllDepartment{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     int64(pageInt),
		Limit:    int64(limitInt),
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDepartments") {
		return
	}
	var departmentsRes model_healthcare_service.ListDepartments
	for _, departmentRes := range departments.Departments {
		departmentsRes.Departments = append(departmentsRes.Departments, &model_healthcare_service.DepartmentRes{
			Id:               departmentRes.Id,
			Name:             departmentRes.Name,
			Description:      departmentRes.Description,
			ImageUrl:         departmentRes.ImageUrl,
			FloorNumber:      departmentRes.FloorNumber,
			ShortDescription: departmentRes.ShortDescription,
			CreatedAt:        departmentRes.CreatedAt,
			UpdatedAt:        departmentRes.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, model_healthcare_service.ListDepartments{
		Count:       int32(departments.Count),
		Departments: departmentsRes.Departments,
	})
}

// UpdateDepartment ...
// @Summary UpdateDepartment
// @Description UpdateDepartment - Api for update department
// @Tags Department
// @Accept json
// @Produce json
// @Param UpdateDepartmentReq body model_healthcare_service.DepartmentReq true "UpdateDepartmentReq"
// @Param id query string true "id"
// @Success 200 {object} model_healthcare_service.DepartmentReq
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/department [put]
func (h *HandlerV1) UpdateDepartment(c *gin.Context) {
	var (
		body        model_healthcare_service.DepartmentReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	id := c.Query("id")

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateDepartment") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	department, err := h.serviceManager.HealthcareService().DepartmentService().UpdateDepartment(ctx, &pb.Department{
		Id:               id,
		Name:             body.Name,
		Description:      body.Description,
		ImageUrl:         body.ImageUrl,
		FloorNumber:      body.FloorNumber,
		ShortDescription: body.ShortDescription,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateDepartment") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DepartmentRes{
		Id:               department.Id,
		Name:             department.Name,
		Description:      department.Description,
		ImageUrl:         department.ImageUrl,
		FloorNumber:      department.FloorNumber,
		ShortDescription: department.ShortDescription,
		CreatedAt:        department.CreatedAt,
		UpdatedAt:        department.UpdatedAt,
	})
}

// DeleteDepartment ...
// @Summary DeleteDepartment
// @Description DeleteDepartment - Api for delete department
// @Tags Department
// @Accept json
// @Produce json
// @Param DeleteDepartmentReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/department [delete]
func (h *HandlerV1) DeleteDepartment(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.HealthcareService().DepartmentService().DeleteDepartment(ctx, &pb.GetReqStrDepartment{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteDepartment") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
