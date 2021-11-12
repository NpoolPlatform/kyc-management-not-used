package kyc

import (
	"context"
	"fmt"
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
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestKycCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()
	kycInfo := &npool.KycInfo{
		UserID:              userID,
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

	resp1, err := Get(context.Background(), &npool.GetKycInfoRequest{
		UserID: kycInfo.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, kycInfo.ID)
		assert.Equal(t, resp1.Info.UserID, kycInfo.UserID)
		assert.Equal(t, resp1.Info.FirstName, kycInfo.FirstName)
		assert.Equal(t, resp1.Info.LastName, kycInfo.LastName)
		assert.Equal(t, resp1.Info.Region, kycInfo.Region)
		assert.Equal(t, resp1.Info.CardType, kycInfo.CardType)
		assert.Equal(t, resp1.Info.FrontCardImg, kycInfo.FrontCardImg)
		assert.Equal(t, resp1.Info.BackCardImg, kycInfo.BackCardImg)
		assert.Equal(t, resp1.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
	}

	resp2, err := Get(context.Background(), &npool.GetKycInfoRequest{
		KycID: kycInfo.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, kycInfo.ID)
		assert.Equal(t, resp2.Info.UserID, kycInfo.UserID)
		assert.Equal(t, resp2.Info.FirstName, kycInfo.FirstName)
		assert.Equal(t, resp2.Info.LastName, kycInfo.LastName)
		assert.Equal(t, resp2.Info.Region, kycInfo.Region)
		assert.Equal(t, resp2.Info.CardType, kycInfo.CardType)
		assert.Equal(t, resp2.Info.FrontCardImg, kycInfo.FrontCardImg)
		assert.Equal(t, resp2.Info.BackCardImg, kycInfo.BackCardImg)
		assert.Equal(t, resp2.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
	}

	resp3, err := Get(context.Background(), &npool.GetKycInfoRequest{
		KycID:  kycInfo.ID,
		UserID: kycInfo.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, kycInfo.ID)
		assert.Equal(t, resp3.Info.UserID, kycInfo.UserID)
		assert.Equal(t, resp3.Info.FirstName, kycInfo.FirstName)
		assert.Equal(t, resp3.Info.LastName, kycInfo.LastName)
		assert.Equal(t, resp3.Info.Region, kycInfo.Region)
		assert.Equal(t, resp3.Info.CardType, kycInfo.CardType)
		assert.Equal(t, resp3.Info.FrontCardImg, kycInfo.FrontCardImg)
		assert.Equal(t, resp3.Info.BackCardImg, kycInfo.BackCardImg)
		assert.Equal(t, resp3.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
	}

	resp4, err := Update(context.Background(), &npool.UpdateKycStatusRequest{
		UserID: kycInfo.UserID,
		Status: true,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, kycInfo.ID)
		assert.Equal(t, resp4.Info.UserID, kycInfo.UserID)
		assert.Equal(t, resp4.Info.FirstName, kycInfo.FirstName)
		assert.Equal(t, resp4.Info.LastName, kycInfo.LastName)
		assert.Equal(t, resp4.Info.Region, kycInfo.Region)
		assert.Equal(t, resp4.Info.CardType, kycInfo.CardType)
		assert.Equal(t, resp4.Info.FrontCardImg, kycInfo.FrontCardImg)
		assert.Equal(t, resp4.Info.BackCardImg, kycInfo.BackCardImg)
		assert.Equal(t, resp4.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
	}

	resp5, err := Update(context.Background(), &npool.UpdateKycStatusRequest{
		KycID:  kycInfo.ID,
		Status: true,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.Info.ID, kycInfo.ID)
		assert.Equal(t, resp5.Info.UserID, kycInfo.UserID)
		assert.Equal(t, resp5.Info.FirstName, kycInfo.FirstName)
		assert.Equal(t, resp5.Info.LastName, kycInfo.LastName)
		assert.Equal(t, resp5.Info.Region, kycInfo.Region)
		assert.Equal(t, resp5.Info.CardType, kycInfo.CardType)
		assert.Equal(t, resp5.Info.FrontCardImg, kycInfo.FrontCardImg)
		assert.Equal(t, resp5.Info.BackCardImg, kycInfo.BackCardImg)
		assert.Equal(t, resp5.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
	}
}
