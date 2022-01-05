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

func dbRowToKyc(row *ent.Kyc) *npool.KycInfo {
	return &npool.KycInfo{
		ID:                  row.ID.String(),
		UserID:              row.UserID.String(),
		FirstName:           row.FirstName,
		LastName:            row.LastName,
		Region:              row.Region,
		CardType:            row.CardType,
		CardID:              row.CardID,
		FrontCardImg:        row.FrontCardImg,
		BackCardImg:         row.BackCardImg,
		UserHandlingCardImg: row.UserHandlingCardImg,
		ReviewStatus:        row.ReviewStatus,
		CreateAT:            row.CreateAt,
		UpdateAT:            row.UpdateAt,
	}
}

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
		return nil, status.Errorf(codes.InvalidArgument, "Invalid kyc id <%v>: %v", in.GetKycID(), err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = kyc.GetKycByID(ctx, kycID)
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("UpdateKycStatus error: %v is not exist in database: %v", in.GetKycID(), err)
		return nil, status.Errorf(codes.NotFound, "This kyc record <%v> can not be found", in.GetKycID())
	}

	kycInfo, err := kyc.UpdateReviewStatus(ctx, kycID, in.GetStatus())
	if err != nil {
		logger.Sugar().Errorf("UpdateKycStatus error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	rowToKycInfo := dbRowToKyc(kycInfo)

	if in.Status == 1 {
		err := grpc.UpdateUserKycStatus(rowToKycInfo.UserID, rowToKycInfo.AppID, true)
		if err != nil {
			logger.Sugar().Errorf("UpdateKycStatus error: %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	} else {
		err := grpc.UpdateUserKycStatus(rowToKycInfo.UserID, rowToKycInfo.AppID, false)
		if err != nil {
			logger.Sugar().Errorf("UpdateKycStatus error: %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}
	return &npool.UpdateKycStatusResponse{
		Info: rowToKycInfo,
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
