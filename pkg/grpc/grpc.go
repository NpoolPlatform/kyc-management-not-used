package grpc

import (
	"context"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
)

func UpdateUserKycStatus(ctx context.Context, userID, appID string, status bool) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbApplication.NewApplicationManagementClient(conn)
	_, err = client.UpdateUserKYCStatus(ctx, &pbApplication.UpdateUserKYCStatusRequest{
		UserID: userID,
		AppID:  appID,
		Status: status,
	})
	if err != nil {
		return err
	}
	return nil
}
