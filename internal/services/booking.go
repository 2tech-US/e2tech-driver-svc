package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lntvan166/e2tech-driver-svc/internal/db"
	"github.com/lntvan166/e2tech-driver-svc/internal/pb"
	"github.com/lntvan166/e2tech-driver-svc/internal/utils"
)

func (s *Server) GetDriverNearby(ctx context.Context, req *pb.GetDriverNearbyRequest) (*pb.GetDriverNearbyResponse, error) {
	arg := db.GetDriverNearbyParams{
		Limit:     req.NumberOfDrivers,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	drivers, err := s.DB.GetDriverNearby(ctx, arg)
	if err != nil {
		return &pb.GetDriverNearbyResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("svc: get driver nearby error: %s", err),
		}, nil
	}

	dataRsp := []*pb.DriverNearby{}
	for _, driver := range drivers {
		dataRsp = append(dataRsp, &pb.DriverNearby{
			Phone:     driver.Phone,
			Distance:  driver.Distance,
			Latitude:  utils.Float64FromNull(driver.Latitude),
			Longitude: utils.Float64FromNull(driver.Longitude),
		})
	}

	return &pb.GetDriverNearbyResponse{
		Status:  http.StatusOK,
		Drivers: dataRsp,
	}, nil
}
