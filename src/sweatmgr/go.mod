module github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr

go 1.14

require (
	github.com/eduardocfalcao/hands-on-go-and-mongodb/src/proto v1.0.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/urfave/negroni v1.0.0
	go.mongodb.org/mongo-driver v1.4.6
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.36.0
)

replace github.com/eduardocfalcao/hands-on-go-and-mongodb/src/proto => ../proto
