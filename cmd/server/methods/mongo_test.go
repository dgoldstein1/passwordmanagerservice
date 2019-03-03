// mongo_test.go

package methods

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"testing"
	"os"
	"errors"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	// "github.com/davecgh/go-spew/spew"
	"github.com/globalsign/mgo/bson"
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

func TestGetEntryFromDB(t *testing.T) {
	// data used in tests
	user1 := pb.DBEntry{
		User : &pb.User{
			First : "test",
			Last : "user1",
			Email : "test@user1.com",
		},
		Auth : &pb.Auth{
			Dn : "test@user1.com",
		},
		Logins : []*pb.Login{},
		Passwords : "lskjdflskdjflskjdf",
	}
	user2 := pb.DBEntry{
		Auth : &pb.Auth{
			Dn : "test@user2.com",
		},
	}
	// initialize connection
	viper.Set("mongodb_endpoint", "mongodb://localhost:27017")
	viper.Set("mongodb_timeout", 1)
	viper.Set("mongodb_name", "passwords")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)
	sess, err := ConnectToMongo(logger)
	c, sess, err := CopySessionAndGetCollection(sess, "passwords")
	if err != nil {
		t.Error(err)
	}
	// insert test data
	c.Insert(user1)
	c.Insert(user2)
	c.Insert(user2)
	// run tests
	var tableTests = []struct {
		name string
		userDn string
		expectedResponse *pb.DBEntry
		expectedError error
	}{
		{"valid retrieval", user1.Auth.Dn, &user1, nil},
		{"finding more than one user", user2.Auth.Dn, nil, errors.New("more than one userDn found for: 'test@user2.com'")},
		{"find no user", "this is a bad user dn", nil, errors.New("no userDn found for: 'this is a bad user dn'")},
	}

	defer func() {
		c.RemoveAll(bson.M{"auth.dn": user1.Auth.Dn})
		c.RemoveAll(bson.M{"auth.dn": user2.Auth.Dn})		
	}()


	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			actualResponse, actualError := GetEntryFromDB(c, tt.userDn)
			if (actualResponse != nil && actualResponse.Auth != nil) {
				AssertEqual(t, actualResponse.Auth.Dn, tt.expectedResponse.Auth.Dn)
			}
			AssertErrorEqual(t, actualError, tt.expectedError)
		})
	}
}
