package grpc

import (
	"context"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
)

func newApplicationGrpcConn() (*grpc.ClientConn, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func UpdateUserKycStatus(userID, appID string, status bool) error {
	conn, err := newApplicationGrpcConn()
	if err != nil {
		return err
	}

	client := pbApplication.NewApplicationManagementClient(conn)
	_, err = client.UpdateUserKYCStatus(context.Background(), &pbApplication.UpdateUserKYCStatusRequest{
		UserID: userID,
		AppID:  appID,
		Status: status,
	})
	if err != nil {
		return err
	}
	return nil
}
