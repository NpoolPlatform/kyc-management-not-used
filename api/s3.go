package api

import (
	"context"
	"errors"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	"github.com/NpoolPlatform/kyc-management/pkg/s3"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UploadKycImage(ctx context.Context, in *npool.UploadKycImageRequest) (*npool.UploadKycImageResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorf("UploadKycImage error: invalid app id <%v>: %v", in.GetAppID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid app id <%v>", in.GetAppID())
	}

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorf("UploadKycImage error: invalid user id <%v>: %v", in.GetUserID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id <%v>", in.GetUserID())
	}

	if in.GetImageType() == "" {
		logger.Sugar().Errorf("UploadKycImage error: image type can not be empty")
		return nil, status.Error(codes.InvalidArgument, "image type can not be empty")
	}

	if in.GetImageBase64() == "" {
		logger.Sugar().Errorf("UploadKycImage error: image base64 url can not be empty")
		return nil, status.Error(codes.InvalidArgument, "image base64 url can not be empty")
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := s3.UploadKycImage(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("UploadKycImage error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetKycImage(ctx context.Context, in *npool.GetKycImageRequest) (*npool.GetKycImageResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorf("GetKycImage error: invalid app id <%v>: %v", in.GetAppID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid app id <%v>", in.GetAppID())
	}

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorf("GetKycImage error: invalid user id <%v>: %v", in.GetUserID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id <%v>", in.GetUserID())
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := s3.GetKycImage(ctx, in)
	if err != nil {
		if errors.Is(err, s3.ErrNoImage) {
			logger.Sugar().Errorf("GetKycImage error: %v", err)
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		logger.Sugar().Errorf("GetKycImage error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	return resp, nil
}
