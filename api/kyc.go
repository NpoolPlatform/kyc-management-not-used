package api

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	"github.com/NpoolPlatform/kyc-management/pkg/grpc"
	mkyc "github.com/NpoolPlatform/kyc-management/pkg/middleware/kyc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateKycRecord(ctx context.Context, in *npool.CreateKycRecordRequest) (*npool.CreateKycRecordResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := kyc.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to create kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetAllKycInfos(ctx context.Context, in *npool.GetAllKycInfosRequest) (*npool.GetAllKycInfosResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := mkyc.GetKycInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UpdateKycStatus(ctx context.Context, in *npool.UpdateKycStatusRequest) (*npool.UpdateKycStatusResponse, error) {
	kycID, err := uuid.Parse(in.GetKycID())
	if err != nil {
		logger.Sugar().Errorf("UpdateKycStatus error: %v is not a valid uuid: %v", in.GetKycID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid kyc id <%v>", in.GetKycID())
	}

	updateStatus, err := kyc.UintToKycState(in.GetStatus())
	if err != nil {
		logger.Sugar().Errorf("UpdateKycStatus error: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid kyc state: %v", err)
	}

	if updateStatus == kyc.WaitState {
		logger.Sugar().Errorf("UpdateKycStatus error: you must change the kyc review status to pass or fail")
		return nil, status.Error(codes.InvalidArgument, "You must change the kyc review status to pass or fail")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	originalKycInfo, err := kyc.GetKycByID(ctx, kycID)
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("UpdateKycStatus error: %v is not exist in database: %v", in.GetKycID(), err)
		return nil, status.Errorf(codes.NotFound, "This kyc record <%v> can not be found", in.GetKycID())
	}

	if originalKycInfo.ReviewStatus != uint32(kyc.WaitState) {
		logger.Sugar().Error("UpdateKycStatus error: kyc status must be waiting!")
		return nil, status.Error(codes.InvalidArgument, "Invalid kyc status, kyc status must be waiting")
	}

	kycInfo, err := kyc.UpdateReviewStatus(ctx, kycID, updateStatus)
	if err != nil {
		logger.Sugar().Errorf("UpdateKycStatus error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if kycInfo.ReviewStatus == uint32(kyc.PassState) {
		err := grpc.UpdateUserKycStatus(ctx, kycInfo.UserID, kycInfo.AppID, true)
		if err != nil {
			logger.Sugar().Errorf("UpdateKycStatus call UpdateUserKycStatus error: %v", err)
			_, err := kyc.UpdateReviewStatus(ctx, kycID, kyc.WaitState)
			if err != nil {
				logger.Sugar().Errorf("UpdateKycStatus call UpdateKycStatus error: %v", err)
				return nil, status.Error(codes.Internal, "internal server error")
			}
			return nil, status.Error(codes.Internal, "internal server error")
		}
	} else {
		err := grpc.UpdateUserKycStatus(ctx, kycInfo.UserID, kycInfo.AppID, false)
		if err != nil {
			logger.Sugar().Errorf("UpdateKycStatus call UpdateUserKycStatus error: %v", err)
			_, err := kyc.UpdateReviewStatus(ctx, kycID, kyc.WaitState)
			if err != nil {
				logger.Sugar().Errorf("UpdateKycStatus call UpdateKycStatus error: %v", err)
				return nil, status.Error(codes.Internal, "internal server error")
			}
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}
	return &npool.UpdateKycStatusResponse{
		Info: kycInfo,
	}, nil
}

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := kyc.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to update kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
