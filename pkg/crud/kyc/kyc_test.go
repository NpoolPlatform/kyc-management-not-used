package kyc

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	testinit "github.com/NpoolPlatform/kyc-management/pkg/test-init"
	npool "github.com/NpoolPlatform/message/npool/kyc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		log.Fatal(err)
	}
}

func TestKycCrud(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()
	appID := uuid.New().String()
	cardType := "test3"
	cardID := "test123131"
	frontCardImage := "31231313112"
	backCardImage := "312313213121"
	userHandingCardImage := "12312312321312"

	createKycRequest := &npool.CreateKycRequest{
		Info: &npool.KycInfo{
			UserID:             userID,
			AppID:              appID,
			CardType:           cardType,
			CardID:             cardID,
			FrontCardImg:       frontCardImage,
			BackCardImg:        backCardImage,
			UserHandingCardImg: userHandingCardImage,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), myconst.GrpcTimeout)
	defer cancel()

	kycID := ""

	createResp, err := Create(ctx, createKycRequest)
	if assert.Nil(t, err) {
		assert.NotNil(t, createResp.GetInfo().GetID())
		kycID = createResp.GetInfo().GetID()
		assert.Equal(t, userID, createResp.GetInfo().GetUserID())
		assert.Equal(t, appID, createResp.GetInfo().GetAppID())
		assert.Equal(t, cardType, createResp.GetInfo().GetCardType())
		assert.Equal(t, cardID, createResp.GetInfo().GetCardID())
		assert.Equal(t, frontCardImage, createResp.GetInfo().GetFrontCardImg())
		assert.Equal(t, backCardImage, createResp.GetInfo().GetBackCardImg())
		assert.Equal(t, userHandingCardImage, createResp.GetInfo().GetUserHandingCardImg())
	}

	getKycByUserIDAndAppIDResp, err := GetKycByUserIDAndAppID(ctx, uuid.MustParse(appID), uuid.MustParse(userID))
	if assert.Nil(t, err) {
		assert.Equal(t, kycID, getKycByUserIDAndAppIDResp.GetID())
		assert.Equal(t, userID, getKycByUserIDAndAppIDResp.GetUserID())
		assert.Equal(t, appID, getKycByUserIDAndAppIDResp.GetAppID())
		assert.Equal(t, cardType, getKycByUserIDAndAppIDResp.GetCardType())
		assert.Equal(t, cardID, getKycByUserIDAndAppIDResp.GetCardID())
		assert.Equal(t, frontCardImage, getKycByUserIDAndAppIDResp.GetFrontCardImg())
		assert.Equal(t, backCardImage, getKycByUserIDAndAppIDResp.GetBackCardImg())
		assert.Equal(t, userHandingCardImage, getKycByUserIDAndAppIDResp.GetUserHandingCardImg())
	}

	existKycByUserIDAndAppIDResp, err := ExistKycByUserIDAndAppID(ctx, uuid.MustParse(appID), uuid.MustParse(userID))
	if assert.Nil(t, err) {
		assert.Equal(t, true, existKycByUserIDAndAppIDResp)
	}

	getKycByIDResp, err := GetKycByID(ctx, uuid.MustParse(kycID))
	if assert.Nil(t, err) {
		assert.Equal(t, kycID, getKycByIDResp.GetID())
		assert.Equal(t, userID, getKycByIDResp.GetUserID())
		assert.Equal(t, appID, getKycByIDResp.GetAppID())
		assert.Equal(t, cardType, getKycByIDResp.GetCardType())
		assert.Equal(t, cardID, getKycByIDResp.GetCardID())
		assert.Equal(t, frontCardImage, getKycByIDResp.GetFrontCardImg())
		assert.Equal(t, backCardImage, getKycByIDResp.GetBackCardImg())
		assert.Equal(t, userHandingCardImage, getKycByIDResp.GetUserHandingCardImg())
	}

	infos, total, err := GetAll(ctx, uuid.MustParse(appID), 5, 0)
	if assert.Nil(t, err) {
		assert.NotEqual(t, 0, total)
		assert.NotNil(t, infos)
		assert.LessOrEqual(t, len(infos), 5)
	}

	infos, total, err = GetAll(ctx, uuid.Nil, 5, 0)
	if assert.Nil(t, err) {
		assert.NotEqual(t, 0, total)
		assert.NotNil(t, infos)
		assert.LessOrEqual(t, len(infos), 5)
	}

	updateKycRequest := &npool.UpdateKycRequest{
		Info: &npool.KycInfo{
			ID:                 kycID,
			UserID:             userID,
			AppID:              appID,
			CardType:           cardType,
			CardID:             cardID,
			FrontCardImg:       frontCardImage,
			BackCardImg:        backCardImage,
			UserHandingCardImg: userHandingCardImage,
		},
	}

	updateKycResp, err := Update(ctx, updateKycRequest)
	if assert.Nil(t, err) {
		assert.Equal(t, kycID, updateKycResp.GetInfo().GetID())
		assert.Equal(t, userID, updateKycResp.GetInfo().GetUserID())
		assert.Equal(t, appID, updateKycResp.GetInfo().GetAppID())
		assert.Equal(t, cardType, updateKycResp.GetInfo().GetCardType())
		assert.Equal(t, cardID, updateKycResp.GetInfo().GetCardID())
		assert.Equal(t, frontCardImage, updateKycResp.GetInfo().GetFrontCardImg())
		assert.Equal(t, backCardImage, updateKycResp.GetInfo().GetBackCardImg())
		assert.Equal(t, userHandingCardImage, updateKycResp.GetInfo().GetUserHandingCardImg())
	}

	exist, err := ExistCradTypeCardIDInAppExceptUserID(ctx, cardType, cardID, uuid.MustParse(appID), uuid.New())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	kycIDs := []uuid.UUID{uuid.MustParse(kycID)}
	infos, err = GetKycByKycIDs(ctx, kycIDs)
	if assert.Nil(t, err) {
		assert.NotNil(t, infos)
	}

	kycIDs = []uuid.UUID{uuid.MustParse("00000000-0000-0000-0000-000000000000")}
	infos, err = GetKycByKycIDs(ctx, kycIDs)
	if assert.Nil(t, err) {
		assert.Equal(t, 0, len(infos))
	}
}
