module github.com/NpoolPlatform/kyc-management

go 1.16

require (
	entgo.io/ent v0.10.0
	github.com/NpoolPlatform/application-management v0.0.0-20211228043636-766772748ce7
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211222114515-4928e6cf3f1f
	github.com/NpoolPlatform/message v0.0.0-20220118090327-926885a280ec
	github.com/boombuler/barcode v1.0.1 // indirect
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	github.com/spf13/cobra v1.3.0 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/t-yuki/gocover-cobertura v0.0.0-20180217150009-aaee18c8195c // indirect
	github.com/tebeka/go2xunit v1.4.10 // indirect
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/tools v0.1.9-0.20211216111533-8d383106f7e7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.42.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
