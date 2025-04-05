package profile

import (
	"context"
	"database/sql"

	sqlcdb "github.com/Abhidalbanjan/ssk-matrimony/db/generated"
	pb "github.com/Abhidalbanjan/ssk-matrimony/proto/profile"
)

type Service struct {
	queries *sqlcdb.Queries
}

func NewService(queries *sqlcdb.Queries) *Service {
	return &Service{queries: queries}
}

func toSQLNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func toSQLNullInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: i, Valid: true}
}

func (s *Service) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.ProfileResponse, error) {
	params := sqlcdb.CreateProfileParams{
		UserID:             req.UserId,
		Name:               toSQLNullString(req.Name),
		Age:                toSQLNullInt32(req.Age),
		BirthDate:          toSQLNullString(req.BirthDate),
		BirthTime:          toSQLNullString(req.BirthTime),
		BirthPlace:         toSQLNullString(req.BirthPlace),
		Nakshatra:          toSQLNullString(req.Nakshatra),
		Rashi:              toSQLNullString(req.Rashi),
		Gender:             toSQLNullString(req.Gender),
		Height:             toSQLNullString(req.Height),
		FatherName:         toSQLNullString(req.FatherName),
		MotherName:         toSQLNullString(req.MotherName),
		TotalFamilyMembers: toSQLNullInt32(req.TotalFamilyMembers),
		Qualification:      toSQLNullString(req.Qualification),
		Degree:             toSQLNullString(req.Degree),
		College:            toSQLNullString(req.College),
		Designation:        toSQLNullString(req.Designation),
		CompanyAndCity:     toSQLNullString(req.CompanyAndCity),
		City:               toSQLNullString(req.City),
		District:           toSQLNullString(req.District),
		State:              toSQLNullString(req.State),
		MaritalStatus:      toSQLNullString(req.MaritalStatus),
		AddDetails:         toSQLNullString(req.AddDetails),
	}

	err := s.queries.CreateProfile(ctx, params)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		UserId:             req.UserId,
		Name:               req.Name,
		Age:                req.Age,
		BirthDate:          req.BirthDate,
		BirthTime:          req.BirthTime,
		BirthPlace:         req.BirthPlace,
		Nakshatra:          req.Nakshatra,
		Rashi:              req.Rashi,
		Gender:             req.Gender,
		Height:             req.Height,
		FatherName:         req.FatherName,
		MotherName:         req.MotherName,
		TotalFamilyMembers: req.TotalFamilyMembers,
		Qualification:      req.Qualification,
		Degree:             req.Degree,
		College:            req.College,
		Designation:        req.Designation,
		CompanyAndCity:     req.CompanyAndCity,
		City:               req.City,
		District:           req.District,
		State:              req.State,
		MaritalStatus:      req.MaritalStatus,
		AddDetails:         req.AddDetails,
	}, nil
}

func (s *Service) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileResponse, error) {
	profile, err := s.queries.GetProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		UserId:             profile.UserID,
		Name:               profile.Name.String,
		Age:                profile.Age.Int32,
		BirthDate:          profile.BirthDate.String,
		BirthTime:          profile.BirthTime.String,
		BirthPlace:         profile.BirthPlace.String,
		Nakshatra:          profile.Nakshatra.String,
		Rashi:              profile.Rashi.String,
		Gender:             profile.Gender.String,
		Height:             profile.Height.String,
		FatherName:         profile.FatherName.String,
		MotherName:         profile.MotherName.String,
		TotalFamilyMembers: profile.TotalFamilyMembers.Int32,
		Qualification:      profile.Qualification.String,
		Degree:             profile.Degree.String,
		College:            profile.College.String,
		Designation:        profile.Designation.String,
		CompanyAndCity:     profile.CompanyAndCity.String,
		City:               profile.City.String,
		District:           profile.District.String,
		State:              profile.State.String,
		MaritalStatus:      profile.MaritalStatus.String,
		AddDetails:         profile.AddDetails.String,
	}, nil
}

func (s *Service) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.ProfileResponse, error) {
	params := sqlcdb.UpdateProfileParams{
		UserID:             req.UserId,
		Name:               toSQLNullString(req.Name),
		Age:                toSQLNullInt32(req.Age),
		BirthDate:          toSQLNullString(req.BirthDate),
		BirthTime:          toSQLNullString(req.BirthTime),
		BirthPlace:         toSQLNullString(req.BirthPlace),
		Nakshatra:          toSQLNullString(req.Nakshatra),
		Rashi:              toSQLNullString(req.Rashi),
		Gender:             toSQLNullString(req.Gender),
		Height:             toSQLNullString(req.Height),
		FatherName:         toSQLNullString(req.FatherName),
		MotherName:         toSQLNullString(req.MotherName),
		TotalFamilyMembers: toSQLNullInt32(req.TotalFamilyMembers),
		Qualification:      toSQLNullString(req.Qualification),
		Degree:             toSQLNullString(req.Degree),
		College:            toSQLNullString(req.College),
		Designation:        toSQLNullString(req.Designation),
		CompanyAndCity:     toSQLNullString(req.CompanyAndCity),
		City:               toSQLNullString(req.City),
		District:           toSQLNullString(req.District),
		State:              toSQLNullString(req.State),
		MaritalStatus:      toSQLNullString(req.MaritalStatus),
		AddDetails:         toSQLNullString(req.AddDetails),
	}

	err := s.queries.UpdateProfile(ctx, params)
	if err != nil {
		return nil, err
	}

	return s.GetProfile(ctx, &pb.GetProfileRequest{UserId: req.UserId})
}
