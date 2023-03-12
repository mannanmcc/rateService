package v1

import (
	"context"
	"errors"

	protos "github.com/mannanmcc/proto/rates/rate"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mannanmcc/rateService/internal/logger"
	"github.com/mannanmcc/rateService/internal/rateservice"
)

type Handler struct {
	service *rateservice.Service
	protos.UnimplementedRateServiceServer
}

func New(service *rateservice.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Register register the handler in the GRPC server
func (h *Handler) Register(server *grpc.Server) {
	protos.RegisterRateServiceServer(server, h)
}

func (h *Handler) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	response, err := h.service.GetRate(ctx, transformRequest(req))
	if err != nil {
		if errors.Is(err, rateservice.ErrInvalidRequest) {
			logger.Print(ctx, "invalid request", zap.Error(err))
			return transformResponse(response), status.Error(codes.InvalidArgument, err.Error())
		}
		logger.Print(ctx, "request failed to process", zap.Error(err))
		return transformResponse(response), status.Error(codes.Internal, err.Error())
	}
	logger.Print(ctx, "request processed", zap.Error(err))
	return transformResponse(response), nil
}
