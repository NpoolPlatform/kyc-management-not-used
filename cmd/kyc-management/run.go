package main

import (
	"time"

	"github.com/NpoolPlatform/kyc-management/api"
	db "github.com/NpoolPlatform/kyc-management/pkg/db"
	msgcli "github.com/NpoolPlatform/kyc-management/pkg/message/client"
	msglistener "github.com/NpoolPlatform/kyc-management/pkg/message/listener"
	msg "github.com/NpoolPlatform/kyc-management/pkg/message/message"
	msgsrv "github.com/NpoolPlatform/kyc-management/pkg/message/server"
	"golang.org/x/xerrors"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/oss"

	apimgrcli "github.com/NpoolPlatform/api-manager/pkg/client"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	ossconst "github.com/NpoolPlatform/go-service-framework/pkg/oss/const"
	"google.golang.org/grpc"
)

const BukectKey = "kyc_bucket"

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		if err := oss.Init(ossconst.SecretStoreKey, BukectKey); err != nil {
			return xerrors.Errorf("fail to init s3: %v", err)
		}

		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()

		if err := msgsrv.Init(); err != nil {
			return err
		}
		if err := msgcli.Init(); err != nil {
			return err
		}

		go msglistener.Listen()
		go msgSender()

		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	apimgrcli.Register(mux)

	return nil
}

func msgSender() {
	id := 0
	for {
		err := msgsrv.PublishExample(&msg.Example{
			ID:      id,
			Example: "hello world",
		})
		if err != nil {
			logger.Sugar().Errorf("fail to send example: %v", err)
			return
		}
		id++
		time.Sleep(3 * time.Second)
	}
}
