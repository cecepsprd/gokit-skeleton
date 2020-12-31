package grpc

import (
	"context"
	"strconv"

	pb "github.com/cecepsprd/gokit-skeleton/api/proto/v1"
	"github.com/cecepsprd/gokit-skeleton/internal/endpoint"
	"github.com/cecepsprd/gokit-skeleton/internal/model"
	gkit "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	getPerson  gkit.Handler
	getPersons gkit.Handler
}

func NewGRPCServer(ep endpoint.PersonEndpoint, opt []gkit.ServerOption) pb.PersonServer {
	return &grpcServer{
		getPerson: gkit.NewServer(
			ep.GetHandler,
			decodeGetPersonRequest,
			encodeGetPersonResponse,
		),
		getPersons: gkit.NewServer(
			ep.GetAllHandler,
			decodeGetPersonsRequest,
			encodeGetPersonsResponse,
		),
	}
}

// GRPC Request handler
func (s *grpcServer) GetPerson(ctx context.Context, req *pb.GetPersonRequest) (*pb.GetPersonResponse, error) {
	_, resp, err := s.getPerson.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetPersonResponse), nil
}

func (s *grpcServer) GetPersons(ctx context.Context, req *pb.ReadAllPersonRequest) (*pb.GetPersonsResponse, error) {
	_, resp, err := s.getPersons.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetPersonsResponse), nil
}

func decodeGetPersonRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetPersonRequest)
	return model.GetPersonRequest{
		ID: req.Id,
	}, nil
}

// encodeCreateCustomerResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetPersonResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(model.Person)
	id := strconv.Itoa(int(res.ID))
	return &pb.GetPersonResponse{
		Id:    id,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func decodeGetPersonsRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}
func encodeGetPersonsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.([]*pb.GetPersonResponse)
	return &pb.GetPersonsResponse{Persons: res}, nil
}
