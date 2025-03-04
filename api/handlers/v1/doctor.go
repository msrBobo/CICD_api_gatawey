package v1

import (
	e "CICD_api_gatawey/api/handlers/regtool"
	"CICD_api_gatawey/api/models"
	"CICD_api_gatawey/api/models/model_healthcare_service"
	pb "CICD_api_gatawey/genproto/healthcare-service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
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
		ImageUrl:      body.ImageUrl,
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
		ImageUrl:      body.ImageUrl,
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
// @Param id query string true "id"
// @Success 200 {object} model_healthcare_service.DoctorAndDoctorHours
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor/get [get]
func (h *HandlerV1) GetDoctor(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctor, err := h.serviceManager.HealthcareService().DoctorService().GetDoctorById(ctx, &pb.GetReqStrDoctor{
		Field:    "id",
		Value:    id,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetDoctor") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DoctorAndDoctorHours{
		Id:            doctor.Id,
		Order:         doctor.Order,
		FirstName:     doctor.FirstName,
		LastName:      doctor.LastName,
		ImageUrl:      doctor.ImageUrl,
		Gender:        doctor.Gender,
		BirthDate:     doctor.BirthDate,
		PhoneNumber:   doctor.PhoneNumber,
		Email:         doctor.Email,
		Address:       doctor.Address,
		City:          doctor.City,
		Country:       doctor.Country,
		Salary:        doctor.Salary,
		StartTime:     doctor.StartTime,
		FinishTime:    doctor.FinishTime,
		DayOfWeek:     doctor.DayOfWeek,
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
// @Param search query string false "search" Enums(first_name, last_name, gender, phone_number, email, address, city, country, biography) "search"
// @Success 200 {object} model_healthcare_service.DoctorAndDoctorHours
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor [get]
func (h *HandlerV1) ListDoctors(c *gin.Context) {
	search := c.Query("search")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListDoctors") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctors, err := h.serviceManager.HealthcareService().DoctorService().GetAllDoctors(ctx, &pb.GetAllDoctorS{
		Field:    search,
		Value:    value,
		IsActive: false,
		Page:     int64(pageInt),
		Limit:    int64(limitInt),
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctors") {
		return
	}
	var doctorsRes model_healthcare_service.ListDoctorsAndHours
	for _, doctorRes := range doctors.DoctorHours {
		doctorsRes.Doctors = append(doctorsRes.Doctors, &model_healthcare_service.DoctorAndDoctorHours{
			Id:            doctorRes.Id,
			Order:         doctorRes.Order,
			FirstName:     doctorRes.FirstName,
			LastName:      doctorRes.LastName,
			ImageUrl:      doctorRes.ImageUrl,
			Gender:        doctorRes.Gender,
			BirthDate:     doctorRes.BirthDate,
			PhoneNumber:   doctorRes.PhoneNumber,
			Email:         doctorRes.Email,
			Address:       doctorRes.Address,
			City:          doctorRes.City,
			Country:       doctorRes.Country,
			Salary:        doctorRes.Salary,
			StartTime:     doctorRes.StartTime,
			FinishTime:    doctorRes.FinishTime,
			DayOfWeek:     doctorRes.DayOfWeek,
			Bio:           doctorRes.Bio,
			StartWorkDate: doctorRes.StartWorkDate,
			EndWorkDate:   doctorRes.EndWorkDate,
			WorkYears:     doctorRes.WorkYears,
			DepartmentId:  doctorRes.DepartmentId,
			RoomNumber:    doctorRes.RoomNumber,
			Password:      doctorRes.Password,
			CreatedAt:     doctorRes.CreatedAt,
			UpdatedAt:     doctorRes.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, model_healthcare_service.ListDoctorsAndHours{
		Count:   doctors.Count,
		Doctors: doctorsRes.Doctors,
	})
}

// ListDoctorsBySpecializationId ...
// @Summary ListDoctorsBySpecializationId
// @Description ListDoctorsBySpecializationId - Api for list doctors by specialization id
// @Tags Doctor
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Param specialization_id query string true "specialization_id"
// @Success 200 {object} model_healthcare_service.ListDoctors
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor/spec [get]
func (h *HandlerV1) ListDoctorsBySpecializationId(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	specId := c.Query("specialization_id")
	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListDoctorsBySpecializationId") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctors, err := h.serviceManager.HealthcareService().DoctorService().ListDoctorBySpecializationId(ctx, &pb.GetReqStrSpec{
		Field:            field,
		Value:            value,
		IsActive:         false,
		Page:             int32(pageInt),
		Limit:            int32(limitInt),
		OrderBy:          orderBy,
		SpecializationId: specId,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctorsBySpecializationId") {
		return
	}
	var doctorsRes model_healthcare_service.ListDoctors
	for _, doctorRes := range doctors.Doctors {
		doctorsRes.Doctors = append(doctorsRes.Doctors, &model_healthcare_service.DoctorRes{
			Id:            doctorRes.Id,
			Order:         doctorRes.Order,
			FirstName:     doctorRes.FirstName,
			LastName:      doctorRes.LastName,
			ImageUrl:      doctorRes.ImageUrl,
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
// @Param UpdateDoctorReq body model_healthcare_service.DoctorUpdateReq true "UpdateDoctorReq"
// @Success 200 {object} model_healthcare_service.DoctorRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor [put]
func (h *HandlerV1) UpdateDoctor(c *gin.Context) {
	var (
		body        model_healthcare_service.DoctorUpdateReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateDoctor") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctor, err := h.serviceManager.HealthcareService().DoctorService().UpdateDoctor(ctx, &pb.Doctor{
		Id:            body.Id,
		FirstName:     body.FirstName,
		LastName:      body.LastName,
		ImageUrl:      body.ImageUrl,
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
		ImageUrl:      doctor.ImageUrl,
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
