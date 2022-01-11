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

func GetKycByAppID(ctx context.Context, appID uuid.UUID) ([]*npool.KycInfo, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	response := []*npool.KycInfo{}
	for _, info := range infos {
		response = append(response, dbRowToKyc(info))
	}
	return response, nil
}

func GetAll(ctx context.Context, in *npool.GetAllKycRequest) ([]*npool.KycInfo, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Kyc.
		Query().All(ctx)
	if err != nil {
		return nil, err
	}

	response := []*npool.KycInfo{}
	for _, info := range infos {
		response = append(response, dbRowToKyc(info))
	}

	return response, nil
}

func Update(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Kyc.
		UpdateOneID(uuid.MustParse(in.GetID())).
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

	return &npool.UpdateKycResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func DeleteUserKycByKycID(ctx context.Context, kycID uuid.UUID) error {
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
