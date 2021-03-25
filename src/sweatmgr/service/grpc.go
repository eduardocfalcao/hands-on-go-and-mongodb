package service

import (
	"context"

	google_protobuf "github.com/golang/protobuf/ptypes"

	pb "github.com/eduardocfalcao/hands-on-go-and-mongodb/src/proto/sweatmgr"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/logger"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/models"
)

type GrpcServer struct {
}

func (s *GrpcServer) GetSweatStats(ctx context.Context, req *pb.SweatStatsRequest) (res *pb.SweatStatsResponse, err error) {
	ctx = context.WithValue(ctx, "UserID", req.Userid)

	sweats, err := models.ListAllSweat(ctx)

	if err != nil {
		logger.Get().Info("Error fetching data", err)
	}

	res = &pb.SweatStatsResponse{
		Sweat: []*pb.Sweat{},
	}

	res.Userid = req.Userid
	for _, sw := range sweats {
		tmp := pb.Sweat{
			Glucose:         sw.Glucose,
			Chloride:        sw.Chloride,
			Sodium:          sw.Sodium,
			Potassium:       sw.Potassium,
			Magnesium:       sw.Magnesium,
			Calcium:         sw.Calcium,
			Humidity:        sw.Humidity,
			RoomTemperature: sw.RoomTemperature,
			BodyTemperature: sw.BodyTemperature,
			Heartbeat:       sw.HeartBeat,
		}
		tmp.CreatedAt, _ = google_protobuf.TimestampProto(sw.CreatedAt)
		res.Sweat = append(res.Sweat, &tmp)
	}
	return
}
