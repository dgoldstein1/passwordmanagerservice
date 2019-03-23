// new_server_test.go

// crud_passwords_test.go

package methods

import (
	"testing"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"os"
	"github.com/spf13/viper"
)

func TestCreateAndRegisterServer(t *testing.T) {
	// positive test
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)

	server := grpc.NewServer()
	CreateAndRegisterServer(logger, server)

}

func TestGivingBadMongodbEndpoint(t *testing.T) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)
	// recover
	defer func() {
	    if r := recover(); r == nil {
	        t.Errorf("The code did not panic")
	    }
	}()
	// test panic
	server := grpc.NewServer()
	viper.Set("mongodb_endpoint", "mongodb://badendpoint")
	CreateAndRegisterServer(logger, server)
}