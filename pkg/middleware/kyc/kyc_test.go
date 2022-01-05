package kyc

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	crudkyc "github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/s3"
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

func TestKycMiddleware(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()

	resp, err := s3.UploadKycImg(context.Background(), &npool.UploadKycImgRequest{
		UserID:    userID,
		ImgType:   "test",
		ImgBase64: "hjoasdidhjasihdasiodhsaiofhjasiofjioasjfiopasjfpoasjfopasjfasopfjpasfjasipjfpj",
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	kycInfo := &npool.KycInfo{
		AppID:               uuid.New().String(),
		UserID:              userID,
		FirstName:           "test",
		LastName:            "test",
		Region:              "test",
		CardType:            "ID Card",
		CardID:              uuid.New().String(),
		FrontCardImg:        resp.Info,
		BackCardImg:         resp.Info,
		UserHandlingCardImg: resp.Info,
	}

	resp1, err := crudkyc.Create(context.Background(), &npool.CreateKycRecordRequest{
		Info: kycInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.Info.ID, uuid.UUID{})
		assert.Equal(t, resp1.Info.UserID, kycInfo.UserID)
		assert.Equal(t, resp1.Info.FirstName, kycInfo.FirstName)
		assert.Equal(t, resp1.Info.LastName, kycInfo.LastName)
		assert.Equal(t, resp1.Info.Region, kycInfo.Region)
		assert.Equal(t, resp1.Info.CardType, kycInfo.CardType)
		assert.Equal(t, resp1.Info.FrontCardImg, kycInfo.FrontCardImg)
		assert.Equal(t, resp1.Info.BackCardImg, kycInfo.BackCardImg)
		assert.Equal(t, resp1.Info.UserHandlingCardImg, kycInfo.UserHandlingCardImg)
		kycInfo.ID = resp1.Info.ID
	}

	resp2, err := GetKycInfo(context.Background(), &npool.GetAllKycInfosRequest{
		KycIDs: []string{kycInfo.ID},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp2.Infos)
	}
}
