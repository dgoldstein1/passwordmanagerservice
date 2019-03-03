// generate_challenge_test.go

package methods

import (
	"golang.org/x/net/context"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"testing"
	"reflect"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"os"
	"github.com/spf13/viper"
	"github.com/globalsign/mgo/bson"
)

func TestGenerateChallenge(t *testing.T) {
	// data used in tests
	validRequest := pb.ChallengeRequest{
		User : "davd@david.com",
		Location : &pb.Location{
			Ip : "192.0.0.1",
		},
	}
	validDBEntry := pb.DBEntry{
		User : &pb.User{
			First : "david",
			Last : "goldstein",
			Email : validRequest.User,
		},
		Auth : &pb.Auth{
			Dn : validRequest.User,
		},
		Logins : []*pb.Login{},
		Passwords : "lskjdflskdjflskjdf",
	}

	// setup
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)
	viper.Set("mongodb_endpoint", "mongodb://localhost:27017")
	viper.Set("mongodb_timeout", 1)
	viper.Set("mongodb_name", "passwords")
	sess, _ := ConnectToMongo(logger)
	s := serverData{logger, *sess}
	ctx := context.TODO()
	// insert data
	c, sess, _ := CopySessionAndGetCollection(sess, "passwords")
	c.Insert(validDBEntry)
	// test table
	var tableTests = []struct {
		name string
		request *pb.ChallengeRequest
		expectedResponse *pb.ChallengeResponse
		expectedError error
	}{

		{"bad user generate challenge request", &pb.ChallengeRequest{}, nil, errors.New("Invalid request: 'user' is a required field.")},
		{"valid request", &validRequest, nil, errors.New("not implemented")},
	}

	//cleanup
	defer func() {
		c.RemoveAll(bson.M{"auth.dn": validDBEntry.Auth.Dn})
	}()

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			actualResponse, actualError := s.GenerateChallenge(ctx, tt.request)
			AssertEqual(t, actualResponse, tt.expectedResponse, )
			AssertErrorEqual(t, actualError, tt.expectedError)
		})
	}
}

// adopted taken from https://gist.github.com/samalba/6059502
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b { 
		return
	}
	// debug.PrintStack()
	t.Errorf("Received '%v' (type %v), expected '%v' (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func AssertErrorEqual(t *testing.T, a error, b error) {
	if (a == nil || b == nil) {
		AssertEqual(t, a, b)
		return
	}
	if (a.Error() == b.Error()) {
		return
	}
	t.Errorf("Received '%v' (type %v), expected '%v' (type %v)", a.Error(), reflect.TypeOf(a), b.Error(), reflect.TypeOf(b))
}