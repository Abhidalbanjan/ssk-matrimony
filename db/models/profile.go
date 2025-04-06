package db

import (
	"database/sql"
)

type Profile struct {
	UserID            string         `json:"user_id"`
	Name              sql.NullString `json:"name"`
	Age               sql.NullInt32  `json:"age"`
	BirthDate         sql.NullString `json:"birth_date"`
	BirthTime         sql.NullString `json:"birth_time"`
	BirthPlace        sql.NullString `json:"birth_place"`
	Nakshatra         sql.NullString `json:"nakshatra"`
	Rashi             sql.NullString `json:"rashi"`
	Gender            sql.NullString `json:"gender"`
	Height            sql.NullString `json:"height"`
	FatherName        sql.NullString `json:"father_name"`
	MotherName        sql.NullString `json:"mother_name"`
	TotalFamilyMembers sql.NullInt32 `json:"total_family_members"`
	Qualification     sql.NullString `json:"qualification"`
	Degree            sql.NullString `json:"degree"`
	College           sql.NullString `json:"college"`
	Designation       sql.NullString `json:"designation"`
	CompanyAndCity    sql.NullString `json:"company_and_city"`
	City              sql.NullString `json:"city"`
	District          sql.NullString `json:"district"`
	State             sql.NullString `json:"state"`
	MaritalStatus     sql.NullString `json:"marital_status"`
	AddDetails        sql.NullString `json:"add_details"`
	MobileNumber      sql.NullString `json:"mobile_number"`
	ProfilePhotoUrl   sql.NullString `json:"profile_photo_url"`
}
