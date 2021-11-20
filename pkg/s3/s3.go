package s3

import (
	"context"
	"encoding/base64"

	"github.com/NpoolPlatform/go-service-framework/pkg/oss"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"golang.org/x/xerrors"
)

func UploadKycImg(ctx context.Context, in *npool.UploadKycImgRequest) (*npool.UploadKycImgResponse, error) {
	encodeImg := base64.StdEncoding.EncodeToString([]byte(in.ImgBase64))
	s3Key := "kyc/" + in.ImgType + in.UserID

	err := oss.PutObject(ctx, s3Key, []byte(encodeImg), true)
	if err != nil {
		return nil, xerrors.Errorf("fail to upload img to s3: %v", err)
	}

	return &npool.UploadKycImgResponse{
		Info: s3Key,
	}, nil
}

func GetKycImg(ctx context.Context, in *npool.GetKycImgRequest) (*npool.GetKycImgResponse, error) {
	resp, err := oss.GetObject(ctx, in.ImgID, true)
	if err != nil {
		return nil, xerrors.Errorf("fail to get img from s3: %v", err)
	}

	if resp == nil {
		return nil, xerrors.Errorf("empty response")
	}

	decodeImg, err := base64.StdEncoding.DecodeString(string(resp))
	if err != nil {
		return nil, xerrors.Errorf("fail to decode img: %v", err)
	}

	return &npool.GetKycImgResponse{
		Info: string(decodeImg),
	}, nil
}
