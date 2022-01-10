package grpc

import (
	"context"

	pbapplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbreview "github.com/NpoolPlatform/review-service/message/npool"
	reviewconst "github.com/NpoolPlatform/review-service/pkg/message/const"
)

func UpdateUserKycStatus(ctx context.Context, userID, appID string, status bool) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbapplication.NewApplicationManagementClient(conn)
	_, err = client.UpdateUserKYCStatus(ctx, &pbapplication.UpdateUserKYCStatusRequest{
		UserID: userID,
		AppID:  appID,
		Status: status,
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateKycReview(ctx context.Context, kycID, appID string) (*pbreview.CreateReviewResponse, error) {
	conn, err := mygrpc.GetGRPCConn(reviewconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbreview.NewReviewServiceClient(conn)
	resp, err := client.CreateReview(ctx, &pbreview.CreateReviewRequest{
		Info: &pbreview.Review{
			ObjectType: "kyc",
			ObjectID:   kycID,
			AppID:      appID,
			Domain:     "kyc-management-npool-top",
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
