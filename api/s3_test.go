package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestS3API(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	userID := uuid.New().String()
	appID := uuid.New().String()
	imgType := "test"
	imgBase64 := "iVBORw0KGgoAAAANSUhEUgAAAB4AAAAZCAYAAAAmNZ4aAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAySURBVEhL7c2hAQAgDMTAp/vv3CI6QxDkTGROX3mgtjjHGMcYxxjHGMcYxxjHmN/GyQBA0AQuiLmS2gAAAABJRU5ErkJggg=="

	resposne := npool.UploadKycImageResponse{}
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UploadKycImageRequest{
			UserID:      userID,
			AppID:       appID,
			ImageType:   imgType,
			ImageBase64: imgBase64,
		}).Post("http://localhost:50120/v1/upload/kyc/image")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &resposne)
		if assert.Nil(t, err) {
			assert.Equal(t, "kyc/"+appID+"/"+userID+"/"+imgType, resposne.Info)
		}
	}

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetKycImageRequest{
			AppID:      appID,
			UserID:     userID,
			ImageS3Key: resposne.Info,
		}).Post("http://localhost:50120/v1/get/kyc/image")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		response1 := npool.GetKycImageResponse{}
		err := json.Unmarshal(resp1.Body(), &response1)
		if assert.Nil(t, err) {
			assert.Equal(t, imgBase64, response1.Info)
		}
	}
}
