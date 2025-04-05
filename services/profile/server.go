package profile

import (
	"context"

	"github.com/Abhidalbanjan/ssk-matrimony/db/generated"
	pb "github.com/Abhidalbanjan/ssk-matrimony/proto/profile"
)

type Server struct {
	pb.UnimplementedProfileServiceServer
	queries *generated.Queries
}

func NewProfileServer(queries *generated.Queries) *Server {
	return &Server{queries: queries}
}

func (s *Server) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.ProfileResponse, error) {
	err := s.queries.CreateProfile(ctx, generated.CreateProfileParams{
		UserID:             req.UserId,
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
	})
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

func (s *Server) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileResponse, error) {
	profile, err := s.queries.GetProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		UserId:             profile.UserID,
		Name:               profile.Name,
		Age:                profile.Age,
		BirthDate:          profile.BirthDate,
		BirthTime:          profile.BirthTime,
		BirthPlace:         profile.BirthPlace,
		Nakshatra:          profile.Nakshatra,
		Rashi:              profile.Rashi,
		Gender:             profile.Gender,
		Height:             profile.Height,
		FatherName:         profile.FatherName,
		MotherName:         profile.MotherName,
		TotalFamilyMembers: profile.TotalFamilyMembers,
		Qualification:      profile.Qualification,
		Degree:             profile.Degree,
		College:            profile.College,
		Designation:        profile.Designation,
		CompanyAndCity:     profile.CompanyAndCity,
		City:               profile.City,
		District:           profile.District,
		State:              profile.State,
		MaritalStatus:      profile.MaritalStatus,
		AddDetails:         profile.AddDetails,
	}, nil
}
