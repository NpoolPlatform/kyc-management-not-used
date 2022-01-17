package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	npool "github.com/NpoolPlatform/message/npool/kyc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func kycInfoCheck(in *npool.KycInfo) error {
	if in.GetCardID() == "" {
		return status.Error(codes.InvalidArgument, "user card id can not be empty")
	}

	if len(in.GetCardID()) > 128 {
		return status.Error(codes.InvalidArgument, "user card id is invalid")
	}

	if in.GetCardType() == "" {
		return status.Error(codes.InvalidArgument, "user card type can not be empty")
	}

	if in.GetFrontCardImg() == "" {
		return status.Error(codes.InvalidArgument, "user front card image can not be empty")
	}

	if in.GetUserHandingCardImg() == "" {
		return status.Error(codes.InvalidArgument, "user handing card image can not be empty")
	}

	if in.GetBackCardImg() == "" {
		return status.Error(codes.InvalidArgument, "user back card image can not be empty")
	}
	return nil
}

func (s *Server) CreateKyc(ctx context.Context, in *npool.CreateKycRequest) (*npool.CreateKycResponse, error) {
	appID, err := uuid.Parse(in.GetInfo().GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: app id <%v> is invalid: %v", in.GetInfo().GetAppID(), err)
		return nil, status.Error(codes.InvalidArgument, "app id is invalid")
	}

	userID, err := uuid.Parse(in.GetInfo().GetUserID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: user id <%v> is invalid: %v", in.GetInfo().GetUserID(), err)
		return nil, status.Error(codes.InvalidArgument, "user id is invalid")
	}

	if err := kycInfoCheck(in.GetInfo()); err != nil {
		logger.Sugar().Errorf("CreateKyc error: %v", err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	if exist, err := kyc.ExistKycByUserIDAndAppID(ctx, appID, userID); err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if exist {
		logger.Sugar().Errorf("CreateKyc error: user <%v> has already created a kyc record in app <%v>", in.GetInfo().GetUserID(), in.GetInfo().GetAppID())
		return nil, status.Error(codes.AlreadyExists, "user has already created a kyc record")
	}

	if exist, err := kyc.ExistCradTypeCardIDInAppExceptUserID(ctx, in.GetInfo().GetCardType(), in.GetInfo().GetCardID(), appID, userID); err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if exist {
		logger.Sugar().Errorf("CreateKyc error: card id <%v> of card type <%v> from app <%v> has already been used by others", in.GetInfo().GetCardID(), in.GetInfo().GetCardType(), in.GetInfo().GetAppID())
		return nil, status.Errorf(codes.AlreadyExists, "card id <%v> of card type <%v> from app <%v> has already been used by others", in.GetInfo().GetCardID(), in.GetInfo().GetCardType(), in.GetInfo().GetAppID())
	}

	resp, err := kyc.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
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
		logger.Sugar().Errorf("GetKycByUserID error: invalid userID<%v>: %v", in.GetUserID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "userID <%v> is not invalid", in.GetUserID())
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := kyc.GetKycByUserIDAndAppID(ctx, appID, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			logger.Sugar().Errorf("GetKycByUserID error: user %v record is not exist in app <%v>", in.GetUserID(), in.GetAppID())
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

	resp, total, err := kyc.GetAll(ctx, appID, in.GetPageInfo().GetLimit(), in.GetPageInfo().GetOffset())
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

	resp, total, err := kyc.GetAll(ctx, uuid.Nil, in.GetPageInfo().GetLimit(), in.GetPageInfo().GetOffset())
	if err != nil {
		logger.Sugar().Errorf("GetAllKyc error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &npool.GetAllKycResponse{
		Infos: resp,
		Total: int32(total),
	}, nil
}

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	if err := kycInfoCheck(in.GetInfo()); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: %v", err.Error())
		return nil, err
	}

	userID, err := uuid.Parse(in.GetInfo().GetUserID())
	if err != nil {
		logger.Sugar().Errorf("UpdateKyc error: invalid appID <%v>: %v", in.GetInfo().GetAppID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid appID <%v>", in.GetInfo().GetAppID())
	}

	appID, err := uuid.Parse(in.GetInfo().GetAppID())
	if err != nil {
		logger.Sugar().Errorf("UpdateKyc error: invalid userID <%v>: %v", in.GetInfo().GetUserID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid userID <%v>", in.GetInfo().GetUserID())
	}

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: invalid kyc id <%v>: %v", in.GetInfo().GetID(), err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid kyc id <%v>", in.GetInfo().GetID())
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	if exist, err := kyc.ExistKycByUserIDAndAppID(ctx, appID, userID); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if !exist {
		logger.Sugar().Errorf("UpdateKyc error: user <%v> has never been done kyc in app <%v>", in.GetInfo().GetUserID(), in.GetInfo().GetAppID())
		return nil, status.Errorf(codes.NotFound, "user <%v> has never been done kyc in app <%v>", in.GetInfo().GetUserID(), in.GetInfo().GetAppID())
	}

	if exist, err := kyc.ExistCradTypeCardIDInAppExceptUserID(ctx, in.GetInfo().GetCardType(), in.GetInfo().GetCardID(), appID, userID); err != nil {
		logger.Sugar().Errorf("UpdateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if exist {
		logger.Sugar().Errorf("UpdateKyc error: card id <%v> of card type <%v> from app <%v> has already been used by others", in.GetInfo().GetCardID(), in.GetInfo().GetCardType(), in.GetInfo().GetAppID())
		return nil, status.Errorf(codes.AlreadyExists, "card id <%v> of card type <%v> from app <%v> has already been used by others", in.GetInfo().GetCardID(), in.GetInfo().GetCardType(), in.GetInfo().GetAppID())
	}

	resp, err := kyc.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("UpdateKyc error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetKycByKycIDs(ctx context.Context, in *npool.GetKycByKycIDsRequest) (*npool.GetKycByKycIDsResponse, error) {
	if in.GetKycIDs() == nil || len(in.GetKycIDs()) == 0 {
		logger.Sugar().Error("GetKycByKycIDs error: kyc ids can not be empty")
		return nil, status.Error(codes.InvalidArgument, "kyc ids can not be empty")
	}

	kycIDs := []uuid.UUID{}
	for _, kycID := range in.GetKycIDs() {
		id, err := uuid.Parse(kycID)
		if err != nil {
			logger.Sugar().Errorf("GetKycByKycIDs error: invalid kyc id <%v>, %v", kycID, err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid kyc id <%v>", kycID)
		}
		kycIDs = append(kycIDs, id)
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := kyc.GetKycByKycIDs(ctx, kycIDs)
	if err != nil {
		logger.Sugar().Errorf("GetKycByKycIDs error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.GetKycByKycIDsResponse{
		Infos: resp,
	}, nil
}
