package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateKycRecord(ctx context.Context, in *npool.CreateKycRecordRequest) (*npool.CreateKycRecordResponse, error) {
	resp, err := kyc.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to create kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetKycInfo(ctx context.Context, in *npool.GetKycInfoRequest) (*npool.GetKycInfoResponse, error) {
	resp, err := kyc.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UpdateKycStatus(ctx context.Context, in *npool.UpdateKycStatusRequest) (*npool.UpdateKycStatusResponse, error) {
	resp, err := kyc.UpdateReviewStatus(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to update kyc status: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	resp, err := kyc.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to update kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
