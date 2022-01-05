package kyc

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kyc-management/message/npool"
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

func TestKycCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()
	appID := uuid.New().String()
	kycInfo := &npool.KycInfo{
		UserID:              userID,
		AppID:               appID,
		FirstName:           "test",
		LastName:            "test",
		Region:              "test",
		CardType:            "ID Card",
		CardID:              uuid.New().String(),
		FrontCardImg:        "front" + userID,
		BackCardImg:         "back" + userID,
		UserHandlingCardImg: "user" + userID,
	}

	resp, err := Create(context.Background(), &npool.CreateKycRecordRequest{
		Info: kycInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.UserID, kycInfo.UserID)
		assert.Equal(t, resp.Info.FirstName, kycInfo.FirstName)
		assert.Equal(t, resp.Info.LastName, kycInfo.LastName)
		assert.Equal(t, resp.Info.Region, kycInfo.Region)
		assert.Equal(t, resp.Info.CardType, kycInfo.CardType)
		assert.Equal(t, resp.Info.FrontCardImg, kycInfo.FrontCardImg)
		assert.Equal(t, resp.Info.BackCardImg, kycInfo.BackCardImg)
		assert.Equal(t, resp.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
		kycInfo.ID = resp.Info.ID
	}

	resp1, err := GetAll(context.Background(), &npool.GetAllKycInfosRequest{
		KycIDs: []string{kycInfo.ID},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
	}

	_, err = UpdateReviewStatus(context.Background(), &npool.UpdateKycStatusRequest{
		UserID: kycInfo.UserID,
		Status: 1,
		AppID:  appID,
	})
	assert.NotNil(t, err)

	_, err = UpdateReviewStatus(context.Background(), &npool.UpdateKycStatusRequest{
		KycID:  kycInfo.ID,
		Status: 2,
		AppID:  appID,
	})
	assert.NotNil(t, err)
	resp6, err := Update(context.Background(), &npool.UpdateKycRequest{
		Info: kycInfo,
	})

	if assert.Nil(t, err) {
		assert.NotNil(t, resp6)
	}
}
