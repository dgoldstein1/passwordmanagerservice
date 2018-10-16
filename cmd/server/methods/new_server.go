package methods

import (
	"google.golang.org/grpc"

	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"github.com/rs/zerolog"
)

type serverData struct {
	zerolog.Logger
}

// NewPasswordserviceServer returns an object that implements the  interface
func CreateAndRegisterServer(
	logger zerolog.Logger,
	grpcServer *grpc.Server,
) {
	var server pb.PasswordserviceServer = &serverData{
		logger,
	}

	pb.RegisterPasswordserviceServer(grpcServer, server)

}
