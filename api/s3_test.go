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
	imgType := "test"
	imgBase64 := "iVBORw0KGgoAAAANSUhEUgAAAB4AAAAZCAYAAAAmNZ4aAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAySURBVEhL7c2hAQAgDMTAp/vv3CI6QxDkTGROX3mgtjjHGMcYxxjHGMcYxxjHmN/GyQBA0AQuiLmS2gAAAABJRU5ErkJggg=="
	imgID := imgType + userID

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UploadKycImgRequest{
			UserID:    userID,
			ImgType:   imgType,
			ImgBase64: imgBase64,
		}).Post("http://localhost:50120/v1/upload/kyc/img")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		resposne := npool.UploadKycImgResponse{}
		err := json.Unmarshal(resp.Body(), &resposne)
		if assert.Nil(t, err) {
			assert.Equal(t, resposne.Info, "kyc/"+imgID)
		}
	}

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetKycImgRequest{
			ImgID: imgID,
		}).Post("http://localhost:50120/v1/get/kyc/img")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		response := npool.GetKycImgResponse{}
		err := json.Unmarshal(resp1.Body(), &response)
		if assert.Nil(t, err) {
			assert.Equal(t, response.Info, imgBase64)
		}
	}
}
