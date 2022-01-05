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
	appID := uuid.New().String()

	imgType := "test"
	imgBase64 := "iVBORw0KGgoAAAANSUhEUgAAAB4AAAAZCAYAAAAmNZ4aAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAySURBVEhL7c2hAQAgDMTAp/vv3CI6QxDkTGROX3mgtjjHGMcYxxjHGMcYxxjHmN/GyQBA0AQuiLmS2gAAAABJRU5ErkJggg=="
	imgID := imgType + userID

	resposne := npool.UploadKycImgResponse{}
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UploadKycImgRequest{
			UserID:    userID,
			ImgType:   imgType,
			ImgBase64: imgBase64,
		}).Post("http://localhost:50120/v1/upload/kyc/img")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &resposne)
		if assert.Nil(t, err) {
			assert.Equal(t, resposne.Info, "kyc/"+imgID)
		}
	}

	kycInfo := &npool.KycInfo{
		UserID:              userID,
		AppID:               appID,
		FirstName:           "test",
		LastName:            "test",
		Region:              "test",
		CardType:            "ID Card",
		CardID:              uuid.New().String(),
		FrontCardImg:        resposne.Info,
		BackCardImg:         resposne.Info,
		UserHandlingCardImg: resposne.Info,
	}

	response := npool.CreateKycRecordResponse{}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateKycRecordRequest{
			Info: kycInfo,
		}).
		Post("http://localhost:50120/v1/create/kyc/record")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		err := json.Unmarshal(resp1.Body(), &response)
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

	response1 := npool.GetAllKycInfosResponse{}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetAllKycInfosRequest{
			KycIDs: []string{kycInfo.ID},
		}).
		Post("http://localhost:50120/v1/get/all/kyc/infos")
	if assert.Nil(t, err) {
		fmt.Println("resp1 is", resp2)
		assert.Equal(t, 200, resp2.StatusCode())
		err := json.Unmarshal(resp2.Body(), &response1)
		if assert.Nil(t, err) {
			assert.NotNil(t, &response1)
		}
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateKycStatusRequest{
			UserID: kycInfo.UserID,
			Status: 1,
		}).
		Post("http://localhost:50120/v1/update/kyc/status")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp3.StatusCode())
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateKycRequest{
			Info: kycInfo,
		}).
		Post("http://localhost:50120/v1/update/kyc")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
	}
}
