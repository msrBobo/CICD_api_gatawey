package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models"
	"dennic_api_gateway/api/models/model_booking_service"
	pb "dennic_api_gateway/genproto/booking_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreatePatient ...
// @Summary CreatePatient
// @Description CreatePatient - Api for crete patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param CreatePatientReq body model_booking_service.CreatePatientReq true "CreatePatientReq"
// @Success 200 {object} model_booking_service.Patient
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/patient [post]
func (h *HandlerV1) CreatePatient(c *gin.Context) {
	var (
		body        model_booking_service.CreatePatientReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreatePatient") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().PatientService().CreatePatient(ctx, &pb.CreatePatientReq{
		Id:             uuid.New().String(),
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		BirthDate:      body.BirthDate,
		Gender:         body.Gender,
		Address:        body.Address,
		BloodGroup:     body.BloodGroup,
		PhoneNumber:    body.PhoneNumber,
		City:           body.City,
		Country:        body.Country,
		PatientProblem: body.PatientProblem,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreatePatient") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Patient{
		Id:             res.Id,
		FirstName:      res.FirstName,
		LastName:       res.LastName,
		BirthDate:      res.BirthDate,
		Gender:         res.Gender,
		Address:        res.Address,
		BloodGroup:     res.BloodGroup,
		PhoneNumber:    res.PhoneNumber,
		City:           res.City,
		Country:        res.Country,
		PatientProblem: res.PatientProblem,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// GetPatient ...
// @Summary GetPatient
// @Description GetPatient - Api for get patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param GetDoctorTimeReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} model_booking_service.Patient
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/patient/get [get]
func (h *HandlerV1) GetPatient(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().PatientService().GetPatient(ctx, &pb.PatientFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetPatient") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Patient{
		Id:             res.Id,
		FirstName:      res.FirstName,
		LastName:       res.LastName,
		BirthDate:      res.BirthDate,
		Gender:         res.Gender,
		Address:        res.Address,
		BloodGroup:     res.BloodGroup,
		PhoneNumber:    res.PhoneNumber,
		City:           res.City,
		Country:        res.Country,
		PatientProblem: res.PatientProblem,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// ListPatient ...
// @Summary ListPatient
// @Description ListPatient - Api for list patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} model_booking_service.PatientsType
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/patient [get]
func (h *HandlerV1) ListPatient(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListPatient") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().PatientService().GetAllPatients(ctx, &pb.GetAllPatientsReq{
		Field:    field,
		Value:    value,
		IsActive: false,
		Page:     pageInt,
		Limit:    limitInt,
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListPatient") {
		return
	}

	var patients model_booking_service.PatientsType
	for _, patient := range res.Patients {
		var patientRes model_booking_service.Patient
		patientRes.Id = patient.Id
		patientRes.FirstName = patient.FirstName
		patientRes.LastName = patient.LastName
		patientRes.BirthDate = patient.BirthDate
		patientRes.Gender = patient.Gender
		patientRes.Address = patient.Address
		patientRes.BloodGroup = patient.BloodGroup
		patientRes.PhoneNumber = patient.PhoneNumber
		patientRes.City = patient.City
		patientRes.Country = patient.Country
		patientRes.PatientProblem = patient.PatientProblem
		patientRes.CreatedAt = patient.CreatedAt
		patientRes.UpdatedAt = e.UpdateTimeFilter(patient.UpdatedAt)
		patients.Patients = append(patients.Patients, &patientRes)
	}

	c.JSON(http.StatusOK, model_booking_service.PatientsType{
		Count:    res.Count,
		Patients: patients.Patients,
	})
}

// UpdatePatient ...
// @Summary UpdatePatient
// @Description UpdatePatient - Api for update patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param patient_id query string true "patient_id"
// @Param UpdatePatientReq body model_booking_service.UpdatePatientReq true "UpdatePatientReq"
// @Success 200 {object} model_booking_service.Patient
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/patient [put]
func (h *HandlerV1) UpdatePatient(c *gin.Context) {
	id := c.Query("patient_id")
	var (
		body        model_booking_service.Patient
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdatePatient") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().PatientService().UpdatePatient(ctx, &pb.UpdatePatientReq{
		Field:          "id",
		Value:          id,
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		BirthDate:      body.BirthDate,
		Gender:         body.Gender,
		Address:        body.Address,
		BloodGroup:     body.BloodGroup,
		City:           body.City,
		Country:        body.Country,
		PatientProblem: body.PatientProblem,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdatePatient") {
		return
	}

	c.JSON(http.StatusOK, model_booking_service.Patient{
		Id:             res.Id,
		FirstName:      res.FirstName,
		LastName:       res.LastName,
		BirthDate:      res.BirthDate,
		Gender:         res.Gender,
		Address:        res.Address,
		BloodGroup:     res.BloodGroup,
		PhoneNumber:    res.PhoneNumber,
		City:           res.City,
		Country:        res.Country,
		PatientProblem: res.PatientProblem,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      e.UpdateTimeFilter(res.UpdatedAt),
	})
}

// UpdatePhonePatient ...
// @Summary UpdatePhonePatient
// @Description UpdatePhonePatient - Api for update phone patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param phone_number query string true "phone_number"
// @Param new_phone_number query string true "new_phone_number"
// @Success 200 {object} model_booking_service.Patient
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/patient/phone [put]
func (h *HandlerV1) UpdatePhonePatient(c *gin.Context) {
	phone := c.Query("phone_number")
	newPhone := c.Query("new_phone_number")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.serviceManager.BookingService().PatientService().UpdatePhonePatient(ctx, &pb.UpdatePhoneNumber{
		Field:       "phone_number",
		Value:       phone,
		PhoneNumber: newPhone,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdatePhonePatient") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: res.Status})
}

// DeletePatient ...
// @Summary DeletePatient
// @Description DeletePatient - Api for delete patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param DeleteArchiveReq query models.FieldValueReq true "FieldValueReq"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/patient [delete]
func (h *HandlerV1) DeletePatient(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.BookingService().PatientService().DeletePatient(ctx, &pb.PatientFieldValueReq{
		Field:    field,
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeletePatient") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
