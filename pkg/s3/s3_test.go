package s3

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

func TestS3(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()
	appID := uuid.New().String()
	imgType := "test"
	imgBase64 := "iVBORw0KGgoAAAANSUhEUgAAAB4AAAAZCAYAAAAmNZ4aAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAySURBVEhL7c2hAQAgDMTAp/vv3CI6QxDkTGROX3mgtjjHGMcYxxjHGMcYxxjHmN/GyQBA0AQuiLmS2gAAAABJRU5ErkJggg=="

	resp, err := UploadKycImage(context.Background(), &npool.UploadKycImageRequest{
		UserID:      userID,
		AppID:       appID,
		ImageType:   imgType,
		ImageBase64: imgBase64,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
		assert.Equal(t, "kyc/"+appID+"/"+userID+"/"+imgType, resp.Info)
	}

	resp1, err := GetKycImage(context.Background(), &npool.GetKycImageRequest{
		ImageS3Key: resp.Info,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
		assert.Equal(t, imgBase64, resp1.Info)
	}
}
