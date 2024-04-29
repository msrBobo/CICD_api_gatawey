package booking_service

import (
	errorapi "dennic_api_gateway/api/errors"
	"dennic_api_gateway/api/handlers"
	_ "dennic_api_gateway/api/models/models_booking_service"
	models_booking_service "dennic_api_gateway/api/models/models_booking_service"
	pb "dennic_api_gateway/genproto/booking_service"
	grpcClient "dennic_api_gateway/internal/infrastructure/grpc_service_client"
	"dennic_api_gateway/internal/pkg/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/casbin/casbin/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

const (
	serviceNamePatient = "Booking"
	spanServiceName
)

type patientHandler struct {
	handlers.BaseHandler
	logger          *zap.Logger
	config          *config.Config
	service         grpcClient.ServiceClient
	enforcer        *casbin.CachedEnforcer
	serviceName     string
	spanServiceName string
}

func NewPatientHandler(option *handlers.HandlerOption) http.Handler {
	handler := patientHandler{
		logger:          option.Logger,
		config:          option.Config,
		service:         option.Service,
		enforcer:        option.Enforcer,
		serviceName:     "BookingService",
		spanServiceName: "ApiBookingService",
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	//defer cancel()

	handler.Cache = option.Cache
	handler.Client = option.Service
	handler.Config = option.Config

	policies := [][]string{
		{"investor", "/v1/patients/create", "POST"},
		{"investor", "/v1/patients/get", "GET"},
		{"investor", "/v1/patients/list", "GET"},
		{"investor", "/v1/patients/update", "PUT"},
		{"investor", "/v1/patients/:phone", "PUT"},
		{"investor", "/v1/patients/delete", "DELETE"},
	}
	for _, policy := range policies {
		_, err := option.Enforcer.AddPolicy(policy)
		if err != nil {
			option.Logger.Error("error during investor enforcer add policies", zap.Error(err))
		}
		//key := fmt.Sprintf("%s:%s:%s", policy[0], policy[1], policy[2])
		//if err := handler.Cache.Set(ctx, key, true, 0); err != nil {
		//	option.Logger.Error("error setting policy in cache", zap.Error(err))
		//}
	}

	//for _, policy := range policies {
	//	key := fmt.Sprintf("%s:%s:%s", policy[0], policy[1], policy[2])
	//	res, err := handler.Cache.Get(ctx, key)
	//	if err != nil {
	//		option.Logger.Error("error getting policy from cache", zap.Error(err))
	//	}
	//	fmt.Printf("Key: %s, Value: %s\n", key, res)
	//}

	//policies := [][]string{
	//	{"investor", "/v1/content/categories", "GET"},
	//	{"investor", "/v1/content/chapters", "GET"},
	//	{"investor", "/v1/content/articles/{chapter_id}", "GET"},
	//	{"investor", "/v1/content/news", "GET"},
	//}
	//for _, policy := range policies {
	//	_, err := option.Enforcer.AddPolicy(policy)
	//	if err != nil {
	//		option.Logger.Error("error during investor enforcer add policies", zap.Error(err))
	//	}
	//}

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		// auth
		//r.Use(middleware.Authorizer(option.Enforcer, option.Logger))

		// content
		r.Post("/create", handler.createPatient())
		r.Get("/get/:key", handler.getPatient())
		r.Get("/list", handler.listPatients())
		r.Put("/update", handler.updatePatient())
		r.Put("/phone", handler.updatePhonePatient())
		r.Delete("/delete/:key", handler.deletePatient())

	})
	return router
}

// createPatient
// @Router /v1/patients/create [post]
// @Summary Create Patient
// @Description Patients
// @Tags Patient
// @Accept json
// @Produce json
// @Param Create body models_booking_service.CreatePatientReq true "Create Patient"
// @Success 200 {object} models_booking_service.Patient
// @Failure 404 {object} models_booking_service.Errors
// @Failure 500 {object} models_booking_service.Errors
func (h *patientHandler) createPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var patient models_booking_service.Patient

		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			h.logger.Error("error decoding product", zap.Error(err))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		res, err := h.service.BookingService().PatientService().CreatePatient(ctx, &pb.CreatePatientReq{
			Id:             uuid.New().String(),
			FirstName:      patient.FirstName,
			LastName:       patient.LastName,
			BirthDate:      patient.BirthDate,
			Gender:         patient.Gender,
			Address:        patient.Address,
			BloodGroup:     patient.BloodGroup,
			PhoneNumber:    patient.PhoneNumber,
			City:           patient.City,
			Country:        patient.Country,
			PatientProblem: patient.PatientProblem,
		})

		if err != nil {
			render.Render(w, r, errorapi.Error(err))
			return
		}

		render.JSON(w, r, &models_booking_service.Patient{
			ID:             res.Id,
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
			UpdatedAt:      res.UpdatedAt,
			DeletedAt:      res.DeletedAt,
		})
	}
}

// getPatient
// @Router /v1/patients/get [get]
// @Summary Create Patient
// @Description Patients
// @Tags Patient
// @Accept json
// @Produce json
// @Param key path string true "key"
// @Success 200 {object} models_booking_service.Patient
// @Failure 404 {object} models_booking_service.Errors
// @Failure 500 {object} models_booking_service.Errors
func (h *patientHandler) getPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params := mux.Vars(r)
		patientId, ok := params["key"]
		if !ok {
			render.Render(w, r, errorapi.Error(fmt.Errorf("key not found in path")))
			return
		}

		res, err := h.service.BookingService().PatientService().GetPatient(ctx, &pb.PatientFieldValueReq{
			Field:    "id",
			Value:    patientId,
			IsActive: true,
		})

		if err != nil {
			render.Render(w, r, errorapi.Error(err))
			return
		}

		render.JSON(w, r, &models_booking_service.Patient{
			ID:             res.Id,
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
			UpdatedAt:      res.UpdatedAt,
			DeletedAt:      res.DeletedAt,
		})
	}
}

// listPatient
// @Router /v1/patients/list [get]
// @Summary List Patient
// @Description Patients
// @Tags Patient
// @Accept json
// @Produce json
// @Param request query models_booking_service.GetAllRequest false "request"
// @Success 200 {object} models_booking_service.Patients
// @Failure 404 {object} models_booking_service.Errors
// @Failure 500 {object} models_booking_service.Errors
func (h *patientHandler) listPatients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		field := r.URL.Query().Get("field")
		value := r.URL.Query().Get("value")
		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")
		orderBy := r.URL.Query().Get("orderBy")
		// search := r.URL.Query().Get("search")

		pageInt, _ := strconv.Atoi(page)
		limitInt, _ := strconv.Atoi(limit)

		var (
			patients models_booking_service.Patients
		)

		res, err := h.service.BookingService().PatientService().GetAllPatients(ctx, &pb.GetAllPatientsReq{
			Field:    field,
			Value:    value,
			IsActive: false,
			Page:     uint64(pageInt),
			Limit:    uint64(limitInt),
			OrderBy:  orderBy,
			// Search: search,
		})

		if err != nil {
			render.Render(w, r, errorapi.Error(err))
			return
		}

		for _, patientRes := range res.Patients {
			var patient models_booking_service.Patient
			patient.ID = patientRes.Id
			patient.FirstName = patientRes.FirstName
			patient.LastName = patientRes.LastName
			patient.BirthDate = patientRes.BirthDate
			patient.Gender = patientRes.Gender
			patient.Address = patientRes.Address
			patient.BloodGroup = patientRes.BloodGroup
			patient.PhoneNumber = patientRes.PhoneNumber
			patient.City = patientRes.City
			patient.Country = patientRes.Country
			patient.PatientProblem = patientRes.PatientProblem
			patient.CreatedAt = patientRes.CreatedAt
			patient.UpdatedAt = patientRes.UpdatedAt
			patient.DeletedAt = patientRes.DeletedAt
			patients.Patients = append(patients.Patients, &patient)
		}
		patients.Count = res.Count
		render.JSON(w, r, patients)
	}
}

// updatePatient
// @Router /v1/patients/update [put]
// @Summary Update Patient
// @Description Patients
// @Tags Patient
// @Accept json
// @Produce json
// @Param field query string true "field"
// @Param value query string true "value"
// @Param Create body models_booking_service.UpdatePatientReq true "Update Patient"
// @Success 200 {object} models_booking_service.Patient
// @Failure 404 {object} models_booking_service.Errors
// @Failure 500 {object} models_booking_service.Errors
func (h *patientHandler) updatePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var patient models_booking_service.UpdatePatientReq

		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			h.logger.Error("error decoding product", zap.Error(err))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		field := r.URL.Query().Get("field")
		value := r.URL.Query().Get("value")

		res, err := h.service.BookingService().PatientService().UpdatePatient(ctx, &pb.UpdatePatientReq{
			Field:          field,
			Value:          value,
			FirstName:      patient.FirstName,
			LastName:       patient.LastName,
			BirthDate:      patient.BirthDate,
			Gender:         patient.Gender,
			Address:        patient.Address,
			BloodGroup:     patient.BloodGroup,
			City:           patient.City,
			Country:        patient.Country,
			PatientProblem: patient.PatientProblem,
		})

		if err != nil {
			render.Render(w, r, errorapi.Error(err))
			return
		}

		render.JSON(w, r, &models_booking_service.Patient{
			ID:             res.Id,
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
			UpdatedAt:      res.UpdatedAt,
			DeletedAt:      res.DeletedAt,
		})
	}
}

// updatePhonePatient
// @Router /v1/patients/phone [put]
// @Summary Update Patient
// @Description Patients
// @Tags Patient
// @Accept json
// @Produce json
// @Param Create body models_booking_service.UpdatePhoneNumber true "Update Patient"
// @Success 200 {object} models_booking_service.Patient
// @Failure 404 {object} models_booking_service.Errors
// @Failure 500 {object} models_booking_service.Errors
func (h *patientHandler) updatePhonePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var patient models_booking_service.UpdatePhoneNumber

		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			h.logger.Error("error decoding product", zap.Error(err))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		res, err := h.service.BookingService().PatientService().UpdatePhonePatient(ctx, &pb.UpdatePhoneNumber{
			Field:       patient.Field,
			Value:       patient.Value,
			PhoneNumber: patient.PhoneNumber,
		})

		if err != nil {
			render.Render(w, r, errorapi.Error(err))
			return
		}

		render.JSON(w, r, &models_booking_service.StatusType{Status: res.Status})
	}
}

// deletePatient
// @Router /v1/patient/delete/:key [delete]
// @Summary Delete Patient
// @Description Patients
// @Tags Patient
// @Accept json
// @Produce json
// @Param key path string true "key"
// @Success 200 {object} models_booking_service.DeleteStatus
// @Failure 404 {object} models_booking_service.Errors
// @Failure 500 {object} models_booking_service.Errors
func (h *patientHandler) deletePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params := mux.Vars(r)
		patientId, ok := params["key"]
		if !ok {
			render.Render(w, r, errorapi.Error(fmt.Errorf("key not found in path")))
			return
		}

		res, err := h.service.BookingService().PatientService().DeletePatient(ctx, &pb.PatientFieldValueReq{
			Field:    "id",
			Value:    patientId,
			IsActive: true,
		})

		if err != nil {
			render.Render(w, r, errorapi.Error(err))
			return
		}

		render.JSON(w, r, &models_booking_service.DeleteStatus{Status: res.Status})
	}
}
