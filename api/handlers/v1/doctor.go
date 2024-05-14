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

// CreateDoctor ...
// @Summary CreateDoctor
// @Description CreateDoctor - Api for crete doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param DoctorReq body model_healthcare_service.DoctorReq true "DoctorReq"
// @Success 200 {object} model_healthcare_service.DoctorRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor [post]
func (h *HandlerV1) CreateDoctor(c *gin.Context) {
	var (
		body        model_healthcare_service.DoctorReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateDoctor") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctor, err := h.serviceManager.HealthcareService().DoctorService().CreateDoctor(ctx, &pb.Doctor{
		Id:            uuid.NewString(),
		FirstName:     body.FirstName,
		LastName:      body.LastName,
		Gender:        body.Gender,
		BirthDate:     body.BirthDate,
		PhoneNumber:   body.PhoneNumber,
		Email:         body.Email,
		Password:      body.Password,
		Address:       body.Address,
		City:          body.City,
		Country:       body.Country,
		Salary:        body.Salary,
		Bio:           body.Bio,
		StartWorkDate: body.StartWorkDate,
		EndWorkDate:   body.EndWorkDate,
		WorkYears:     body.WorkYears,
		DepartmentId:  body.DepartmentId,
		RoomNumber:    body.RoomNumber,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateDoctor") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DoctorRes{
		Id:            doctor.Id,
		Order:         doctor.Order,
		FirstName:     doctor.FirstName,
		LastName:      doctor.LastName,
		Gender:        doctor.Gender,
		BirthDate:     doctor.BirthDate,
		PhoneNumber:   doctor.PhoneNumber,
		Email:         doctor.Email,
		Address:       doctor.Address,
		City:          doctor.City,
		Country:       doctor.Country,
		Salary:        doctor.Salary,
		Bio:           doctor.Bio,
		StartWorkDate: doctor.StartWorkDate,
		EndWorkDate:   doctor.EndWorkDate,
		WorkYears:     doctor.WorkYears,
		DepartmentId:  doctor.DepartmentId,
		RoomNumber:    doctor.RoomNumber,
		Password:      doctor.Password,
		CreatedAt:     doctor.CreatedAt,
		UpdatedAt:     e.UpdateTimeFilter(doctor.UpdatedAt),
	})
}

// GetDoctor ...
// @Summary GetDoctor
// @Description GetDoctor - Api for get doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param GetDoctor query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_healthcare_service.DoctorRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor/get [get]
func (h *HandlerV1) GetDoctor(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctor, err := h.serviceManager.HealthcareService().DoctorService().GetDoctorById(ctx, &pb.GetReqStrDoctor{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetDoctor") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DoctorRes{
		Id:            doctor.Id,
		Order:         doctor.Order,
		FirstName:     doctor.FirstName,
		LastName:      doctor.LastName,
		Gender:        doctor.Gender,
		BirthDate:     doctor.BirthDate,
		PhoneNumber:   doctor.PhoneNumber,
		Email:         doctor.Email,
		Address:       doctor.Address,
		City:          doctor.City,
		Country:       doctor.Country,
		Salary:        doctor.Salary,
		Bio:           doctor.Bio,
		StartWorkDate: doctor.StartWorkDate,
		EndWorkDate:   doctor.EndWorkDate,
		WorkYears:     doctor.WorkYears,
		DepartmentId:  doctor.DepartmentId,
		RoomNumber:    doctor.RoomNumber,
		Password:      doctor.Password,
		CreatedAt:     doctor.CreatedAt,
		UpdatedAt:     e.UpdateTimeFilter(doctor.UpdatedAt),
	})
}

// ListDoctors ...
// @Summary ListDoctors
// @Description ListDoctors - Api for list doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_healthcare_service.ListDoctors
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor [get]
func (h *HandlerV1) ListDoctors(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctors") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctors, err := h.serviceManager.HealthcareService().DoctorService().GetAllDoctors(ctx, &pb.GetAllDoctorS{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     int64(pageInt),
		Limit:    int64(limitInt),
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctors") {
		return
	}
	var doctorsRes model_healthcare_service.ListDoctors
	for _, doctorRes := range doctors.Doctors {
		doctorsRes.Doctors = append(doctorsRes.Doctors, &model_healthcare_service.DoctorRes{
			Id:            doctorRes.Id,
			Order:         doctorRes.Order,
			FirstName:     doctorRes.FirstName,
			LastName:      doctorRes.LastName,
			Gender:        doctorRes.Gender,
			BirthDate:     doctorRes.BirthDate,
			PhoneNumber:   doctorRes.PhoneNumber,
			Email:         doctorRes.Email,
			Address:       doctorRes.Address,
			City:          doctorRes.City,
			Country:       doctorRes.Country,
			Salary:        doctorRes.Salary,
			Bio:           doctorRes.Bio,
			StartWorkDate: doctorRes.StartWorkDate,
			EndWorkDate:   doctorRes.EndWorkDate,
			WorkYears:     doctorRes.WorkYears,
			DepartmentId:  doctorRes.DepartmentId,
			RoomNumber:    doctorRes.RoomNumber,
			Password:      doctorRes.Password,
			CreatedAt:     doctorRes.CreatedAt,
			UpdatedAt:     e.UpdateTimeFilter(doctorRes.UpdatedAt),
		})
	}
	c.JSON(http.StatusOK, model_healthcare_service.ListDoctors{
		Count:   doctors.Count,
		Doctors: doctorsRes.Doctors,
	})
}

// UpdateDoctor ...
// @Summary UpdateDoctor
// @Description UpdateDoctor - Api for update doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param UpdateDoctorReq body model_healthcare_service.DoctorReq true "UpdateDoctorReq"
// @Param id query string true "id"
// @Success 200 {object} model_healthcare_service.DoctorRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor [put]
func (h *HandlerV1) UpdateDoctor(c *gin.Context) {
	var (
		body        model_healthcare_service.DoctorReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	id := c.Query("id")

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateDoctor") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctor, err := h.serviceManager.HealthcareService().DoctorService().UpdateDoctor(ctx, &pb.Doctor{
		Id:            id,
		FirstName:     body.FirstName,
		LastName:      body.LastName,
		Gender:        body.Gender,
		BirthDate:     body.BirthDate,
		PhoneNumber:   body.PhoneNumber,
		Email:         body.Email,
		Password:      body.Password,
		Address:       body.Address,
		City:          body.City,
		Country:       body.Country,
		Salary:        body.Salary,
		Bio:           body.Bio,
		StartWorkDate: body.StartWorkDate,
		EndWorkDate:   body.EndWorkDate,
		WorkYears:     body.WorkYears,
		DepartmentId:  body.DepartmentId,
		RoomNumber:    body.RoomNumber,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateDoctor") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DoctorRes{
		Id:            doctor.Id,
		FirstName:     doctor.FirstName,
		LastName:      doctor.LastName,
		Gender:        doctor.Gender,
		BirthDate:     doctor.BirthDate,
		PhoneNumber:   doctor.PhoneNumber,
		Email:         doctor.Email,
		Password:      doctor.Password,
		Address:       doctor.Address,
		City:          doctor.City,
		Country:       doctor.Country,
		Salary:        doctor.Salary,
		Bio:           doctor.Bio,
		StartWorkDate: doctor.StartWorkDate,
		EndWorkDate:   doctor.EndWorkDate,
		WorkYears:     doctor.WorkYears,
		DepartmentId:  doctor.DepartmentId,
		RoomNumber:    doctor.RoomNumber,
		CreatedAt:     doctor.CreatedAt,
		UpdatedAt:     e.UpdateTimeFilter(doctor.UpdatedAt),
	})
}

// DeleteDoctor ...
// @Summary DeleteDoctor
// @Description DeleteDoctor - Api for delete doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param DeleteDoctorReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor [delete]
func (h *HandlerV1) DeleteDoctor(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.HealthcareService().DoctorService().DeleteDoctor(ctx, &pb.GetReqStrDoctor{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteDoctor") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
