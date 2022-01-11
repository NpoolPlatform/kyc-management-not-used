package s3

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/oss"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"golang.org/x/xerrors"
)

func UploadKycImg(ctx context.Context, in *npool.UploadKycImgRequest) (*npool.UploadKycImgResponse, error) {
	s3Key := "kyc/" + in.GetAppID() + "/" + in.GetUserID() + "/" + in.GetImgType()

	err := oss.PutObject(ctx, s3Key, []byte(in.GetImgBase64()), true)
	if err != nil {
		return nil, xerrors.Errorf("fail to upload img to s3: %v", err)
	}

	return &npool.UploadKycImgResponse{
		Info: s3Key,
	}, nil
}

func GetKycImg(ctx context.Context, in *npool.GetKycImgRequest) (*npool.GetKycImgResponse, error) {
	resp, err := oss.GetObject(ctx, in.GetImgID(), true)
	if err != nil {
		return nil, xerrors.Errorf("fail to get img from s3: %v", err)
	}

	if resp == nil {
		return nil, xerrors.Errorf("empty response")
	}

	return &npool.GetKycImgResponse{
		Info: string(resp),
	}, nil
}
