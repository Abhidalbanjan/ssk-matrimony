package profile

import (
	"context"
	"log"

	pb "github.com/Abhidalbanjan/ssk-matrimony/proto/profile"
)

type Server struct {
	pb.UnimplementedProfileServiceServer
}

func (s *Server) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.ProfileResponse, error) {
	log.Printf("CreateProfile: %+v\n", req)

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
	log.Printf("GetProfile: %+v\n", req)

	// TODO: Replace with actual database lookup
	return &pb.ProfileResponse{
		UserId:             req.UserId,
		Name:               "Sample Name",
		Age:                30,
		BirthDate:          "01-01-1990",
		BirthTime:          "12:00",
		BirthPlace:         "Sample City",
		Nakshatra:          "Sample Nakshatra",
		Rashi:              "Sample Rashi",
		Gender:             "Sample Gender",
		Height:             "Sample Height",
		FatherName:         "Sample Father",
		MotherName:         "Sample Mother",
		TotalFamilyMembers: 4,
		Qualification:      "Sample Qualification",
		Degree:             "Sample Degree",
		College:            "Sample College",
		Designation:        "Sample Designation",
		CompanyAndCity:     "Sample Company",
		City:               "Sample City",
		District:           "Sample District",
		State:              "Sample State",
		MaritalStatus:      "Single",
		AddDetails:         "Sample Details",
	}, nil
}
