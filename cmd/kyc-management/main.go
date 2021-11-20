package main

import (
	"fmt"
	"os"

	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	ossconst "github.com/NpoolPlatform/go-service-framework/pkg/oss/const"
	rabbitmqconst "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
	cli "github.com/urfave/cli/v2"
)

const serviceName = "Kyc Management"

func main() {
	commands := cli.Commands{
		runCmd,
	}

	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)
	err := app.Init(serviceName, description, "", "", "./", nil, commands,
		config.ServiceNameToNamespace(mysqlconst.MysqlServiceName),
		config.ServiceNameToNamespace(redisconst.RedisServiceName),
		config.ServiceNameToNamespace(rabbitmqconst.RabbitMQServiceName),
		config.ServiceNameToNamespace(applicationconst.ServiceName),
		config.ServiceNameToNamespace(ossconst.S3NameSpace),
	)
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", serviceName, err)
		return
	}
	err = app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorf("fail to run %v: %v", serviceName, err)
	}
}
