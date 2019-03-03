// mongo_test.go

package methods

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"testing"
	"os"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
)


func TestConnectToMongo(t *testing.T) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)

	// positive test
	viper.Set("mongodb_endpoint", "mongodb://localhost:27017")
	viper.Set("mongodb_timeout", 1)
	sess, err := ConnectToMongo(logger)

	if (sess == nil || err != nil) {
		t.Errorf("could not connect to mongo")
	}
}

func TestCopySessionAndGetCollection(t *testing.T) {
	_, _, err := CopySessionAndGetCollection(nil, "null")
	if err == nil {
		t.Errorf("expected error to be thrown passing nil mongo session")		
	}

	// try to connect 
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)

	// positive test
	sess, err := ConnectToMongo(logger)

	c, newSess, err := CopySessionAndGetCollection(sess, "passwords")
	if err != nil {
		t.Error(err)
	}

	if c == nil {
		t.Errorf("expected collection to be returned")
	}

	if newSess == nil {
		t.Errorf("expected session to be returned")
	}

	if newSess == sess {
		t.Errorf("expected new session to be different than old session")
	}

	// return error when connecting to bad db
	viper.Set("mongodb_endpoint", "mongodb://lskdjflskjdflksjdf:27017")
	sess, err = ConnectToMongo(logger)

	_, _, err = CopySessionAndGetCollection(sess, "passwords")
	if err == nil {
		t.Errorf("expected error to be thrown connecting to bad endpoint")
	}
}

func TestGetUserFromDB(t *testing.T) {
	// data used in tests
	
	// setup
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)
	sess, err := ConnectToMongo(logger)
	c, newSess, err := CopySessionAndGetCollection(sess, "user")
	if (err != nil || c == nil || newSess == nil) {
		t.Errorf("Could not connect to mongodb")
	}
	// test table
	var tableTests = []struct {
		name string
		userName string
		expectedUser &pb.User
		expectedError error

	}{
		{"Valid Retrieval Of user"},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			actualResponse, actualError := s.GenerateChallenge(ctx, tt.request)
			AssertEqual(t, actualResponse, tt.expectedResponse, )
			AssertErrorEqual(t, actualError, tt.expectedError)
		})
	}
}