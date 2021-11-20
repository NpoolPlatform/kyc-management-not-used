package kyc

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/db"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	WaitAudit = "auditing"
	PassAudit = "audited"
	FailAudit = "failed"
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
		SetReviewStatus(WaitAudit).
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
	kycIDs := []uuid.UUID{}
	for _, id := range in.KycIDs {
		id, err := uuid.Parse(id)
		if err != nil {
			return nil, xerrors.Errorf("%v user kyc id invalid: %v", id, err)
		}
		kycIDs = append(kycIDs, id)
	}
	response := []*npool.KycInfo{}
	for _, kycid := range kycIDs {
		info, err := db.Client().
			Kyc.
			Query().
			Where(
				kyc.ID(kycid),
			).Only(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to get %v kyc info: %v", kycid, err)
		}

		response = append(response, dbRowToKyc(info))
	}

	return &npool.GetKycInfoResponse{
		Infos: response,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	userID, err := uuid.Parse(in.Info.UserID)
	if err != nil {
		return nil, err
	}
	id := ""
	if in.Info.ID == "" {
		info, err := db.Client().
			Kyc.
			Query().
			Where(
				kyc.UserID(userID),
			).Only(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to query user kyc info: %v", err)
		}

		id = info.ID.String()
	} else {
		id = in.Info.ID
	}

	kycID, err := uuid.Parse(id)
	if err != nil {
		return nil, xerrors.Errorf("invalid kyc record id: %v", err)
	}

	info, err := db.Client().
		Kyc.
		UpdateOneID(kycID).
		SetUserID(userID).
		SetFirstName(in.Info.FirstName).
		SetLastName(in.Info.LastName).
		SetRegion(in.Info.Region).
		SetCardType(in.Info.CardType).
		SetCardID(in.Info.CardID).
		SetFrontCardImg(in.Info.FrontCardImg).
		SetBackCardImg(in.Info.BackCardImg).
		SetUserHandlingCardImg(in.Info.UserHandlingCardImg).
		SetReviewStatus(WaitAudit).
		Save(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, xerrors.Errorf("fail to update user kyc: %v", err)
	}

	return &npool.UpdateKycResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func UpdateReviewStatus(ctx context.Context, in *npool.UpdateKycStatusRequest) (*npool.UpdateKycStatusResponse, error) {
	userID, id, err := parse2ID(in.UserID, in.KycID)
	if err != nil {
		return nil, err
	}
	status := ""
	switch in.Status {
	case 1:
		status = PassAudit
	case 2:
		status = FailAudit
	default:
		status = WaitAudit
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
		SetReviewStatus(status).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update kyc status: %v", err)
	}

	return &npool.UpdateKycStatusResponse{
		Info: status,
	}, nil
}
