package kyc

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	testinit "github.com/NpoolPlatform/kyc-management/pkg/test-init"
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
	firstName := "test1"
	lastName := "test2"
	region := "cn"
	cardType := "test3"
	cardID := "test123131"
	frontCardImage := "31231313112"
	backCardImage := "312313213121"
	userHandlingCardImage := "12312312321312"

	createKycRequest := &npool.CreateKycRequest{
		UserID:              userID,
		AppID:               appID,
		FirstName:           firstName,
		LastName:            lastName,
		Region:              region,
		CardType:            cardType,
		CardID:              cardID,
		FrontCardImg:        frontCardImage,
		BackCardImg:         backCardImage,
		UserHandlingCardImg: userHandlingCardImage,
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
		assert.Equal(t, firstName, createResp.GetInfo().GetFirstName())
		assert.Equal(t, lastName, createResp.GetInfo().GetLastName())
		assert.Equal(t, region, createResp.GetInfo().GetRegion())
		assert.Equal(t, cardType, createResp.GetInfo().GetCardType())
		assert.Equal(t, cardID, createResp.GetInfo().GetCardID())
		assert.Equal(t, frontCardImage, createResp.GetInfo().GetFrontCardImg())
		assert.Equal(t, backCardImage, createResp.GetInfo().GetBackCardImg())
		assert.Equal(t, userHandlingCardImage, createResp.GetInfo().GetUserHandlingCardImg())
	}

	getKycByUserIDAndAppIDResp, err := GetKycByUserIDAndAppID(ctx, uuid.MustParse(appID), uuid.MustParse(userID))
	if assert.Nil(t, err) {
		assert.Equal(t, kycID, getKycByUserIDAndAppIDResp.GetID())
		assert.Equal(t, userID, getKycByUserIDAndAppIDResp.GetUserID())
		assert.Equal(t, appID, getKycByUserIDAndAppIDResp.GetAppID())
		assert.Equal(t, firstName, getKycByUserIDAndAppIDResp.GetFirstName())
		assert.Equal(t, lastName, getKycByUserIDAndAppIDResp.GetLastName())
		assert.Equal(t, region, getKycByUserIDAndAppIDResp.GetRegion())
		assert.Equal(t, cardType, getKycByUserIDAndAppIDResp.GetCardType())
		assert.Equal(t, cardID, getKycByUserIDAndAppIDResp.GetCardID())
		assert.Equal(t, frontCardImage, getKycByUserIDAndAppIDResp.GetFrontCardImg())
		assert.Equal(t, backCardImage, getKycByUserIDAndAppIDResp.GetBackCardImg())
		assert.Equal(t, userHandlingCardImage, getKycByUserIDAndAppIDResp.GetUserHandlingCardImg())
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
		assert.Equal(t, firstName, getKycByIDResp.GetFirstName())
		assert.Equal(t, lastName, getKycByIDResp.GetLastName())
		assert.Equal(t, region, getKycByIDResp.GetRegion())
		assert.Equal(t, cardType, getKycByIDResp.GetCardType())
		assert.Equal(t, cardID, getKycByIDResp.GetCardID())
		assert.Equal(t, frontCardImage, getKycByIDResp.GetFrontCardImg())
		assert.Equal(t, backCardImage, getKycByIDResp.GetBackCardImg())
		assert.Equal(t, userHandlingCardImage, getKycByIDResp.GetUserHandlingCardImg())
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
		ID:                  kycID,
		UserID:              userID,
		AppID:               appID,
		FirstName:           firstName,
		LastName:            lastName,
		Region:              region,
		CardType:            cardType,
		CardID:              cardID,
		FrontCardImg:        frontCardImage,
		BackCardImg:         backCardImage,
		UserHandlingCardImg: userHandlingCardImage,
	}

	updateKycResp, err := Update(ctx, updateKycRequest)
	if assert.Nil(t, err) {
		assert.Equal(t, kycID, updateKycResp.GetInfo().GetID())
		assert.Equal(t, userID, updateKycResp.GetInfo().GetUserID())
		assert.Equal(t, appID, updateKycResp.GetInfo().GetAppID())
		assert.Equal(t, firstName, updateKycResp.GetInfo().GetFirstName())
		assert.Equal(t, lastName, updateKycResp.GetInfo().GetLastName())
		assert.Equal(t, region, updateKycResp.GetInfo().GetRegion())
		assert.Equal(t, cardType, updateKycResp.GetInfo().GetCardType())
		assert.Equal(t, cardID, updateKycResp.GetInfo().GetCardID())
		assert.Equal(t, frontCardImage, updateKycResp.GetInfo().GetFrontCardImg())
		assert.Equal(t, backCardImage, updateKycResp.GetInfo().GetBackCardImg())
		assert.Equal(t, userHandlingCardImage, updateKycResp.GetInfo().GetUserHandlingCardImg())
	}

	exist, err := ExistCradTypeCardIDInAppExceptUserID(ctx, cardType, cardID, uuid.MustParse(appID), uuid.New())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}
}
