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
		CreateAt:            row.CreateAt,
		UpdateAt:            row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateKycRequest) (*npool.CreateKycResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Kyc.
		Create().
		SetUserID(uuid.MustParse(in.GetUserID())).
		SetAppID(uuid.MustParse(in.GetAppID())).
		SetFirstName(in.GetFirstName()).
		SetLastName(in.GetLastName()).
		SetRegion(in.GetRegion()).
		SetCardType(in.GetCardType()).
		SetCardID(in.GetCardID()).
		SetFrontCardImg(in.GetFrontCardImg()).
		SetBackCardImg(in.GetBackCardImg()).
		SetUserHandlingCardImg(in.GetUserHandlingCardImg()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateKycResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func GetKycByUserIDAndAppID(ctx context.Context, appID, userID uuid.UUID) (*npool.KycInfo, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	info, err := cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.UserID(userID),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	return dbRowToKyc(info), nil
}

func ExistKycByUserIDAndAppID(ctx context.Context, appID, userID uuid.UUID) (bool, error) {
	cli, err := db.Client()
	if err != nil {
		return false, xerrors.Errorf("fail get db client: %v", err)
	}

	return cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.UserID(userID),
			),
		).
		Exist(ctx)
}

func GetKycByID(ctx context.Context, kycID uuid.UUID) (*npool.KycInfo, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.Kyc.Query().
		Where(
			kyc.ID(kycID),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	return dbRowToKyc(info), nil
}

func GetAll(ctx context.Context, in *npool.GetAllKycInfosRequest) (*npool.GetAllKycInfosResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

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
		info, err := cli.
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

	return &npool.GetAllKycInfosResponse{
		Infos: response,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	userID, err := uuid.Parse(in.Info.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	appID, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}
	id := ""
	if in.Info.ID == "" {
		info, err := cli.
			Kyc.
			Query().
			Where(
				kyc.And(
					kyc.UserID(userID),
					kyc.AppID(appID),
				),
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

	info, err := cli.
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
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user kyc: %v", err)
	}

	return &npool.UpdateKycResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func DeleteUserKycRecordByKycID(ctx context.Context, kycID uuid.UUID) error {
	cli, err := db.Client()
	if err != nil {
		return xerrors.Errorf("fail to get db client: %v", err)
	}

	return cli.Kyc.DeleteOneID(kycID).Exec(ctx)
}

func ExistCradTypeCardIDInApp(ctx context.Context, cardType, cardID string, appID uuid.UUID) (bool, error) {
	cli, err := db.Client()
	if err != nil {
		return false, xerrors.Errorf("fail to get db client: %v", err)
	}

	return cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.CardType(cardType),
				kyc.CardID(cardID),
			),
		).
		Exist(ctx)
}
