package model_healthcare_service

type DoctorReq struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Gender        string  `json:"gender"`
	BirthDate     string  `json:"birth_date"`
	PhoneNumber   string  `json:"phone_number"`
	Email         string  `json:"email"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	Country       string  `json:"country"`
	Salary        float32 `json:"salary"`
	Bio           string  `json:"bio"`
	StartWorkDate string  `json:"start_work_date"`
	EndWorkDate   string  `json:"end_work_date"`
	WorkYears     int32   `json:"work_years"`
	DepartmentId  string  `json:"department_id"`
	RoomNumber    int32   `json:"room_number"`
	Password      string  `json:"password"`
}

type DoctorRes struct {
	Id            string  `json:"id"`
	Order         int32   `json:"order"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Gender        string  `json:"gender"`
	BirthDate     string  `json:"birth_date"`
	PhoneNumber   string  `json:"phone_number"`
	Email         string  `json:"email"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	Country       string  `json:"country"`
	Salary        float32 `json:"salary"`
	Bio           string  `json:"bio"`
	StartWorkDate string  `json:"start_work_date"`
	EndWorkDate   string  `json:"end_work_date"`
	WorkYears     int32   `json:"work_years"`
	DepartmentId  string  `json:"department_id"`
	RoomNumber    int32   `json:"room_number"`
	Password      string  `json:"password"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type ListDoctorByDepIdReq struct {
	DepartmentId string `json:"department_id"`
	IsActive     bool   `json:"is_active"`
	Page         int32  `json:"page"`
	Limit        int32  `json:"limit"`
	Field        string `json:"field"`
	Value        string `json:"value"`
	OrderBy      string `json:"order_by"`
}

type ListDoctors struct {
	Count   int64        `json:"count"`
	Doctors []*DoctorRes `json:"doctors"`
}
