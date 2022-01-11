package s3

import (
	"context"
	"errors"

	"github.com/NpoolPlatform/go-service-framework/pkg/oss"
	"github.com/NpoolPlatform/kyc-management/message/npool"
)

var ErrNoImage = errors.New("there is no image from this s3 key")

func UploadKycImage(ctx context.Context, in *npool.UploadKycImageRequest) (*npool.UploadKycImageResponse, error) {
	s3Key := "kyc/" + in.GetAppID() + "/" + in.GetUserID() + "/" + in.GetImageType()

	err := oss.PutObject(ctx, s3Key, []byte(in.GetImageBase64()), true)
	if err != nil {
		return nil, err
	}

	return &npool.UploadKycImageResponse{
		Info: s3Key,
	}, nil
}

func GetKycImage(ctx context.Context, in *npool.GetKycImageRequest) (*npool.GetKycImageResponse, error) {
	resp, err := oss.GetObject(ctx, in.GetImageS3Key(), true)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, ErrNoImage
	}

	return &npool.GetKycImageResponse{
		Info: string(resp),
	}, nil
}
