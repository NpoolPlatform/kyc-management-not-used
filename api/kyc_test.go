package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestKycAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

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

	response := npool.CreateKycRecordResponse{}
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateKycRecordRequest{
			Info: kycInfo,
		}).
		Post("http://localhost:32759/v1/create/kyc/record")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Info.ID, uuid.UUID{})
			assert.Equal(t, response.Info.UserID, kycInfo.UserID)
			assert.Equal(t, response.Info.FirstName, kycInfo.FirstName)
			assert.Equal(t, response.Info.LastName, kycInfo.LastName)
			assert.Equal(t, response.Info.Region, kycInfo.Region)
			assert.Equal(t, response.Info.CardType, kycInfo.CardType)
			assert.Equal(t, response.Info.FrontCardImg, kycInfo.FrontCardImg)
			assert.Equal(t, response.Info.BackCardImg, kycInfo.BackCardImg)
			assert.Equal(t, response.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
			kycInfo.ID = response.Info.ID
		}
	}

	response1 := npool.GetKycInfoResponse{}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetKycInfoRequest{
			UserID: kycInfo.UserID,
		}).
		Post("http://localhost:32759/v1/get/kyc/info")
	if assert.Nil(t, err) {
		fmt.Println("resp1 is", resp1)
		assert.Equal(t, 200, resp1.StatusCode())
		err := json.Unmarshal(resp1.Body(), &response1)
		if assert.Nil(t, err) {
			assert.Equal(t, response1.Info.ID, kycInfo.ID)
			assert.Equal(t, response1.Info.UserID, kycInfo.UserID)
			assert.Equal(t, response1.Info.FirstName, kycInfo.FirstName)
			assert.Equal(t, response1.Info.LastName, kycInfo.LastName)
			assert.Equal(t, response1.Info.Region, kycInfo.Region)
			assert.Equal(t, response1.Info.CardType, kycInfo.CardType)
			assert.Equal(t, response1.Info.FrontCardImg, kycInfo.FrontCardImg)
			assert.Equal(t, response1.Info.BackCardImg, kycInfo.BackCardImg)
			assert.Equal(t, response1.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
		}
	}

	response2 := npool.UpdateKycStatusResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateKycStatusRequest{
			UserID: kycInfo.UserID,
			Status: true,
		}).
		Post("http://localhost:32759/v1/update/kyc/status")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		err := json.Unmarshal(resp2.Body(), &response2)
		if assert.Nil(t, err) {
			assert.Equal(t, response2.Info.ID, kycInfo.ID)
			assert.Equal(t, response2.Info.UserID, kycInfo.UserID)
			assert.Equal(t, response2.Info.FirstName, kycInfo.FirstName)
			assert.Equal(t, response2.Info.LastName, kycInfo.LastName)
			assert.Equal(t, response2.Info.Region, kycInfo.Region)
			assert.Equal(t, response2.Info.CardType, kycInfo.CardType)
			assert.Equal(t, response2.Info.FrontCardImg, kycInfo.FrontCardImg)
			assert.Equal(t, response2.Info.BackCardImg, kycInfo.BackCardImg)
			assert.Equal(t, response2.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
		}
	}
}
