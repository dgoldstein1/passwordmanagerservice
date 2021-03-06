package methods

import (
	"google.golang.org/grpc"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"github.com/rs/zerolog"
	"github.com/globalsign/mgo"
)

type serverData struct {
	logger zerolog.Logger
	session mgo.Session

}

// NewPasswordserviceServer returns an object that implements the  interface
func CreateAndRegisterServer(
	logger zerolog.Logger,
	grpcServer *grpc.Server,
) {
	// connect to mongo
	session, err := ConnectToMongo(logger)
	if err != nil {
		panic(err)
	}
	
	var server pb.PasswordserviceServer = &serverData{
		logger : logger,
		session : *session,
	}

	pb.RegisterPasswordserviceServer(grpcServer, server)

}
