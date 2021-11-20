package kyc

import (
	"context"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	crudkyc "github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/s3"
)

func GetKycInfo(ctx context.Context, in *npool.GetKycInfoRequest) (*npool.GetKycInfoResponse, error) {
	resp, err := crudkyc.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	response := []*npool.KycInfo{}
	for _, info := range resp.Infos {
		frontCardImg, err := s3.GetKycImg(ctx, &npool.GetKycImgRequest{
			ImgID: info.FrontCardImg,
		})
		if err != nil {
			return nil, err
		}
		info.FrontCardImg = frontCardImg.Info

		backCardImg, err := s3.GetKycImg(ctx, &npool.GetKycImgRequest{
			ImgID: info.BackCardImg,
		})
		if err != nil {
			return nil, err
		}
		info.BackCardImg = backCardImg.Info

		userHandlingCardImg, err := s3.GetKycImg(ctx, &npool.GetKycImgRequest{
			ImgID: info.UserHandlingCardImg,
		})
		if err != nil {
			return nil, err
		}
		info.UserHandlingCardImg = userHandlingCardImg.Info
		response = append(response, info)
	}
	return &npool.GetKycInfoResponse{
		Infos: response,
	}, nil
}
