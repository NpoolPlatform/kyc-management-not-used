package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	mygrpc "github.com/NpoolPlatform/kyc-management/pkg/grpc"
	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	mkyc "github.com/NpoolPlatform/kyc-management/pkg/middleware/kyc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateKycRecord(ctx context.Context, in *npool.CreateKycRecordRequest) (*npool.CreateKycRecordResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	if in.GetInfo() == nil {
		logger.Sugar().Error("CreateKycRecord error: kyc create info can not be empty")
		return nil, status.Error(codes.InvalidArgument, "kyc create info can not be empty")
	}

	if in.GetInfo().GetCardID() == "" {
		logger.Sugar().Error("CreateKycRecord error: user card id can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user card id can not be empty")
	}

	if in.GetInfo().GetCardType() == "" {
		logger.Sugar().Error("CreateKycRecord error: user card type can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user card type can not be empty")
	}

	if in.GetInfo().GetFirstName() == "" {
		logger.Sugar().Error("CreateKycRecord error: user first name can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user first name can not be empty")
	}

	if in.GetInfo().GetLastName() == "" {
		logger.Sugar().Error("CreateKycRecord error: user last name can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user last name can not be empty")
	}

	if in.GetInfo().GetRegion() == "" {
		logger.Sugar().Error("CreateKycRecord error: user region can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user region can not be empty")
	}

	if in.GetInfo().GetFrontCardImg() == "" {
		logger.Sugar().Error("CreateKycRecord error: user front card image can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user front card image can not be empty")
	}

	if in.GetInfo().GetUserHandlingCardImg() == "" {
		logger.Sugar().Error("CreateKycRecord error: user handling card image can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user front card image can not be empty")
	}

	if in.GetInfo().GetBackCardImg() == "" {
		logger.Sugar().Error("CreateKycRecord error: user back card image can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user back card image can not be empty")
	}

	appid, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateKycRecord error: app id is invalid: %v", err)
		return nil, status.Error(codes.InvalidArgument, "app id is invalid")
	}

	userid, err := uuid.Parse(in.GetUserID())
	if err != nil {
		logger.Sugar().Errorf("CreateKycRecord error: user id is invalid: %v", err)
		return nil, status.Error(codes.InvalidArgument, "user id is invalid")
	}

	kycInfo, err := kyc.GetKycByUserIDAndAppID(ctx, appid, userid)
	if kycInfo != nil && err == nil {
		logger.Sugar().Error("CreateKycRecord error: user has been done kyc in this app")
		return nil, status.Error(codes.InvalidArgument, "user has been done kyc in this app")
	}

	if !ent.IsNotFound(err) {
		logger.Sugar().Errorf("CreateKycRecord error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	resp, err := kyc.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("CreateKycRecord error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	_, err = mygrpc.CreateKycReview(ctx, resp.GetInfo().GetID(), resp.GetInfo().GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateKycRecord call CreateReview error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return resp, nil
}

func (s *Server) GetAllKycInfos(ctx context.Context, in *npool.GetAllKycInfosRequest) (*npool.GetAllKycInfosResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := mkyc.GetKycInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := kyc.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to update kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
