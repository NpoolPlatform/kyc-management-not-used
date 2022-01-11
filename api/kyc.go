package api

import (
	"context"
	"unicode/utf8"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	mygrpc "github.com/NpoolPlatform/kyc-management/pkg/grpc"
	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func kycInfoCheck(in *npool.KycInfo) error {
	if in.GetCardID() == "" {
		return status.Error(codes.InvalidArgument, "user card id can not be empty")
	}

	if len(in.GetCardID()) > 30 {
		return status.Error(codes.InvalidArgument, "user card id is not invalid")
	}

	if in.GetCardType() == "" {
		return status.Error(codes.InvalidArgument, "user card type can not be empty")
	}

	if in.GetFirstName() == "" {
		return status.Error(codes.InvalidArgument, "user first name can not be empty")
	}

	if utf8.RuneCountInString(in.GetFirstName()) > 50 {
		return status.Error(codes.InvalidArgument, "user first name is not invalid")
	}

	if in.GetLastName() == "" {
		return status.Error(codes.InvalidArgument, "user last name can not be empty")
	}

	if utf8.RuneCountInString(in.GetLastName()) > 50 {
		return status.Error(codes.InvalidArgument, "user last name is not invalid")
	}

	if in.GetRegion() == "" {
		return status.Error(codes.InvalidArgument, "user region can not be empty")
	}

	if in.GetFrontCardImg() == "" {
		return status.Error(codes.InvalidArgument, "user front card image can not be empty")
	}

	if in.GetUserHandlingCardImg() == "" {
		return status.Error(codes.InvalidArgument, "user front card image can not be empty")
	}

	if in.GetBackCardImg() == "" {
		return status.Error(codes.InvalidArgument, "user back card image can not be empty")
	}
	return nil
}

func (s *Server) CreateKyc(ctx context.Context, in *npool.CreateKycRequest) (*npool.CreateKycResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: app id <%v> is invalid: %v", in.GetAppID(), err)
		return nil, status.Error(codes.InvalidArgument, "app id is invalid")
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: user id <%v> is invalid: %v", in.GetUserID(), err)
		return nil, status.Error(codes.InvalidArgument, "user id is invalid")
	}

	if err := kycInfoCheck(&npool.KycInfo{
		FirstName:           in.GetFirstName(),
		LastName:            in.GetLastName(),
		Region:              in.GetRegion(),
		CardType:            in.GetCardType(),
		CardID:              in.GetCardID(),
		FrontCardImg:        in.GetFrontCardImg(),
		BackCardImg:         in.GetBackCardImg(),
		UserHandlingCardImg: in.GetUserHandlingCardImg(),
	}); err != nil {
		logger.Sugar().Errorf("CreateKyc error: %v", err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	if exist, err := kyc.ExistKycByUserIDAndAppID(ctx, appID, userID); err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if exist {
		logger.Sugar().Error("CreateKyc error: user has already created a kyc record")
		return nil, status.Error(codes.AlreadyExists, "user has already created a kyc record")
	}

	if exist, err := kyc.ExistCradTypeCardIDInApp(ctx, in.GetCardType(), in.GetCardID(), appID); err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if exist {
		logger.Sugar().Error("CreayeKyc error: this card type card id has been existed in this app")
		return nil, status.Error(codes.AlreadyExists, "this card type card id has been existed in this app")
	}

	resp, err := kyc.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	_, err = mygrpc.CreateKycReview(ctx, resp.GetInfo().GetID(), resp.GetInfo().GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc call CreateReview error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return resp, nil
}

func (s *Server) GetKycByUserID(ctx context.Context, in *npool.GetKycByUserIDRequest) (*npool.GetKycByUserIDResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetKycByUserID error: invalid appID<%v>: %v", in.GetAppID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "appID <%v> is not invalid", in.GetAppID())
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		logger.Sugar().Errorf("GetKycByUserID error: invalid userID<%v>: %v", in.GetAppID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "userID <%v> is not invalid", in.GetAppID())
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := kyc.GetKycByUserIDAndAppID(ctx, appID, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			logger.Sugar().Errorf("GetKycByUserID error: user %v record is not exist", in.GetUserID())
			return nil, status.Errorf(codes.NotFound, "user %v record is not exist", in.GetUserID())
		}
		logger.Sugar().Errorf("GetKycByUserIDAndAppID error: internal sever error: %v", appID, err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.GetKycByUserIDResponse{
		Info: resp,
	}, nil
}

func (s *Server) GetKycByAppID(ctx context.Context, in *npool.GetKycByAppIDRequest) (*npool.GetKycByAppIDResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetKycByAppID error: invalid appID<%v>: %v", in.GetAppID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "appID <%v> is not invalid", in.GetAppID())
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, total, err := kyc.GetAll(ctx, appID, in.Limit, in.Offset)
	if err != nil {
		logger.Sugar().Errorf("GetKycByAppID <%v> error: internal sever error: %v", appID, err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.GetKycByAppIDResponse{
		Infos: resp,
		Total: int32(total),
	}, nil
}

func (s *Server) GetAllKyc(ctx context.Context, in *npool.GetAllKycRequest) (*npool.GetAllKycResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, total, err := kyc.GetAll(ctx, uuid.Nil, in.Limit, in.Offset)
	if err != nil {
		logger.Sugar().Errorf("fail to get kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &npool.GetAllKycResponse{
		Infos: resp,
		Total: int32(total),
	}, nil
}

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: invalid appID <%v>: %v", in.GetAppID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid appID <%v>", in.GetAppID())
	}

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: invalid userID <%v>: %v", in.GetUserID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid userID <%v>", in.GetUserID())
	}

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: invalid kyc id <%v>: %v", in.GetID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid kyc id <%v>", in.GetID())
	}

	if err := kycInfoCheck(&npool.KycInfo{
		FirstName:           in.GetFirstName(),
		LastName:            in.GetLastName(),
		Region:              in.GetRegion(),
		CardType:            in.GetCardType(),
		CardID:              in.GetCardID(),
		FrontCardImg:        in.GetFrontCardImg(),
		BackCardImg:         in.GetBackCardImg(),
		UserHandlingCardImg: in.GetUserHandlingCardImg(),
	}); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: %v", err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := kyc.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to update kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
