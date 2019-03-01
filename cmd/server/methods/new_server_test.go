// new_server_test.go

// crud_passwords_test.go

package methods

import (
	"testing"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"os"
)

func TestCreateAndRegisterServer(t *testing.T) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)

	server := grpc.NewServer()
	CreateAndRegisterServer(logger, server)
}