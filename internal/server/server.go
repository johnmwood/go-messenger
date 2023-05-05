package server

import (
	"context"

	"github.com/johnmwood/go-messenger/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

func (s *Server) Report(ctx context.Context, req *proto.MetricRequest) (*proto.MetricResponse, error) {
	if req.Message == "" {
		return nil, status.Error(codes.InvalidArgument, "unable to report: missing message")
	}
	return &proto.MetricResponse{Message: "Reporting metrics: " + req.Message}, nil
}
