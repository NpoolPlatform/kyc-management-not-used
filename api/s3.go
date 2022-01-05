package api

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/s3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UploadKycImg(ctx context.Context, in *npool.UploadKycImgRequest) (*npool.UploadKycImgResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := s3.UploadKycImg(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to upload img to s3: %v", err)
		return &npool.UploadKycImgResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetKycImg(ctx context.Context, in *npool.GetKycImgRequest) (*npool.GetKycImgResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := s3.GetKycImg(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get img from s3: %v", err)
		return &npool.GetKycImgResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
