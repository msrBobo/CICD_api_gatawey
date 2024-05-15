package model_healthcare_service

type DoctorReq struct {
	FirstName     string  `json:"first_name" example:"First Name"`
	LastName      string  `json:"last_name" example:"Last Name"`
	ImageUrl      string  `json:"image_url" example:"http://example.com/image.png"`
	Gender        string  `json:"gender" example:"male"`
	BirthDate     string  `json:"birth_date" example:"2012-12-12"`
	PhoneNumber   string  `json:"phone_number" example:"+998901234567"`
	Email         string  `json:"email" example:"email@gmail.com"`
	Address       string  `json:"address" example:"Addres"`
	City          string  `json:"city" example:"City"`
	Country       string  `json:"country" example:"Country"`
	Salary        float32 `json:"salary" example:"10"`
	Bio           string  `json:"bio" example:"Biography"`
	StartWorkDate string  `json:"start_work_date" example:"2012-12-12"`
	WorkYears     int32   `json:"work_years" example:"4"`
	DepartmentId  string  `json:"department_id" example:"123e4567-e89b-12d3-a456-426614174001"`
	RoomNumber    int32   `json:"room_number" example:"1"`
	Password      string  `json:"password" example:"password"`
}

type DoctorUpdateReq struct {
	Id            string  `json:"id" example:"123e4567-e89b-12d3-a456-426614274001"`
	FirstName     string  `json:"first_name" example:"First Name"`
	LastName      string  `json:"last_name" example:"Last Name"`
	ImageUrl      string  `json:"image_url" example:"http://example.com/image.png"`
	Gender        string  `json:"gender" example:"male"`
	BirthDate     string  `json:"birth_date" example:"2012-12-12"`
	PhoneNumber   string  `json:"phone_number" example:"+998901234567"`
	Email         string  `json:"email" example:"email@gmail.com"`
	Address       string  `json:"address" example:"Addres"`
	City          string  `json:"city" example:"City"`
	Country       string  `json:"country" example:"Country"`
	Salary        float32 `json:"salary" example:"10"`
	Bio           string  `json:"bio" example:"Biography"`
	StartWorkDate string  `json:"start_work_date" example:"2012-12-12"`
	EndWorkDate   string  `json:"end-work-date" example:"2022-12-12"`
	WorkYears     int32   `json:"work_years" example:"4"`
	DepartmentId  string  `json:"department_id" example:"123e4567-e89b-12d3-a456-426614174001"`
	RoomNumber    int32   `json:"room_number" example:"1"`
	Password      string  `json:"password" example:"password"`
}

type DoctorRes struct {
	Id            string  `json:"id"`
	Order         int32   `json:"order"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	ImageUrl      string  `json:"image_url"`
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

type ListDoctors struct {
	Count   int64        `json:"count"`
	Doctors []*DoctorRes `json:"doctors"`
}

type DoctorAndDoctorHours struct {
	Id            string  `json:"id"`
	Order         int32   `json:"order"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	ImageUrl      string  `json:"image_url"`
	Gender        string  `json:"gender"`
	BirthDate     string  `json:"birth_date"`
	PhoneNumber   string  `json:"phone_number"`
	Email         string  `json:"email"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	Country       string  `json:"country"`
	Salary        float32 `json:"salary"`
	StartTime     string  `json:"start_time"`
	FinishTime    string  `json:"finish_time"`
	DayOfWeek     string  `json:"day_of_week"`
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

type ListDoctorsAndHours struct {
	Count   int64                   `json:"count"`
	Doctors []*DoctorAndDoctorHours `json:"doctors"`
}
