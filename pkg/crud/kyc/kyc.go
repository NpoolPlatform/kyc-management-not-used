package kyc

import (
	"context"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/db"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
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

func Create(ctx context.Context, in *npool.CreateKycRecordRequest) (*npool.CreateKycRecordResponse, error) {
	userID, err := uuid.Parse(in.Info.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	info, err := db.Client().
		Kyc.
		Create().
		SetUserID(userID).
		SetFirstName(in.Info.FirstName).
		SetLastName(in.Info.LastName).
		SetRegion(in.Info.Region).
		SetCardType(in.Info.CardType).
		SetCardID(in.Info.CardID).
		SetFrontCardImg(in.Info.FrontCardImg).
		SetBackCardImg(in.Info.BackCardImg).
		SetUserHandlingCardImg(in.Info.UserHandlingCardImg).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create user kyc: %v", err)
	}

	return &npool.CreateKycRecordResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func parse2ID(userIDString, idString string) (uuid.UUID, uuid.UUID, error) { // nolint
	var userID, id uuid.UUID
	var err error
	if userIDString != "" {
		userID, err = uuid.Parse(userIDString)
		if err != nil {
			return uuid.UUID{}, uuid.UUID{}, xerrors.Errorf("invalid user id: %v", err)
		}
	}
	if idString != "" {
		id, err = uuid.Parse(idString)
		if err != nil {
			return uuid.UUID{}, uuid.UUID{}, xerrors.Errorf("invalid kyc id: %v", err)
		}
	}
	return userID, id, nil
}

func Get(ctx context.Context, in *npool.GetKycInfoRequest) (*npool.GetKycInfoResponse, error) {
	userID, id, err := parse2ID(in.UserID, in.KycID)
	if err != nil {
		return nil, err
	}

	info, err := db.Client().
		Kyc.
		Query().
		Where(
			kyc.Or(
				kyc.UserID(userID),
				kyc.ID(id),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get kyc info: %v", err)
	}
	return &npool.GetKycInfoResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateKycStatusRequest) (*npool.UpdateKycStatusResponse, error) {
	userID, id, err := parse2ID(in.UserID, in.KycID)
	if err != nil {
		return nil, err
	}

	_, err = db.Client().
		Kyc.
		Update().
		Where(
			kyc.Or(
				kyc.UserID(userID),
				kyc.ID(id),
			),
		).
		SetReviewStatus(in.Status).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update kyc status: %v", err)
	}

	info, err := db.Client().
		Kyc.
		Query().
		Where(
			kyc.Or(
				kyc.ID(id),
				kyc.UserID(userID),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get kyc info: %v", err)
	}

	return &npool.UpdateKycStatusResponse{
		Info: dbRowToKyc(info),
	}, nil
}
