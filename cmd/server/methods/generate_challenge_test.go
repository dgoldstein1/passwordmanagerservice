// generate_challenge_test.go

package methods

import (
	"golang.org/x/net/context"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"testing"
	"reflect"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
	// "github.com/davecgh/go-spew/spew"
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
			FailedLogins : 0,
			AuthQuestions : []*pb.AuthQuestion{
				&pb.AuthQuestion{
					Q : "what is your favorite color?",
					A : "yello",		
				},
			},
			KnownIps : []string{"192.0.0.1"},
		},
		Logins : []*pb.Login{},
		Passwords : "lskjdflskdjflskjdf",
	}

	nonExistentUserRequest := pb.ChallengeRequest{
		User : "sdlkjf239jf-23jsdf",
		Location : &pb.Location{
			Ip : "192.0.0.1",
		},
	}

	lockedOutRequest := pb.ChallengeRequest{
		User : "locked@out.com",
		Location : &pb.Location{
			Ip : "192.0.0.1",
		},
	}
	lockedOutUser := pb.DBEntry{
		User : &pb.User{
			First : "david",
			Last : "goldstein",
			Email : lockedOutRequest.User,
		},
		Auth : &pb.Auth{
			Dn : lockedOutRequest.User,
			FailedLogins : 6,
		},
	}

	unknownLocationRequest := pb.ChallengeRequest{
		User : "unkown@location.com",
		Location : &pb.Location{
			Ip : "192.0.0.2",
		},
	}

	unknownLocationResponse := pb.ChallengeResponse{
		Error : "Login Unsuccessful",
		UserQuestion : "what is your favorite color?",
	}

	unknownLocation := pb.DBEntry{
		Auth : &pb.Auth{
			Dn : unknownLocationRequest.User,
			FailedLogins : 0,
			AuthQuestions : []*pb.AuthQuestion{
				&pb.AuthQuestion{
					Q : "what is your favorite color?",
					A : "yello",		
				},
			},
			KnownIps : []string{"192.0.0.1"},
		},
	}

	wrongAnswerRequest := pb.ChallengeRequest{
		User : "wrong@answer.com",
		Location : &pb.Location{
			Ip : "192.0.0.2",
		},
		UserQuestionResponse : &pb.AuthQuestion{
			Q : "what is your favorite color?",
			A : "yellow",
		},
	}

	wrongAnswerResponse := pb.ChallengeResponse{
		Error : "Login Unsuccessful",
		UserQuestion : "what is your favorite color?",
	}

	wrongAnswer := pb.DBEntry{
		Auth : &pb.Auth{
			Dn : wrongAnswerRequest.User,
			FailedLogins : 0,
			AuthQuestions : []*pb.AuthQuestion{
				&pb.AuthQuestion{
					Q : "what is your favorite color?",
					A : "yello",		
				},
			},
			KnownIps : []string{"192.0.0.1"},
		},
	}

	alreadyExistsRequest := pb.ChallengeRequest{
		User : "already@exists.com",
		Location : &pb.Location{
			Ip : "192.0.0.1",
		},
	}
	alreadyExistsEntry := pb.DBEntry{
		Auth : &pb.Auth{
			Dn : alreadyExistsRequest.User,
			FailedLogins : 0,
			KnownIps : []string{"192.0.0.1"},
			AccessToken : "sdfslkjlkje",
		},
		Logins : []*pb.Login{},
		Passwords : "lskjdflskdjflskjdf",
	}

	// setup
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(5)
	viper.Set("mongodb_endpoint", "mongodb://localhost:27017")
	viper.Set("mongodb_endpoint", "mongodb://localhost:27017")
	viper.Set("mongodb_timeout", 1)
	viper.Set("max_failed_logins", 5)
	viper.Set("mongodb_name", "passwords")
	sess, _ := ConnectToMongo(logger)
	s := serverData{logger, *sess}
	ctx := context.TODO()
	// insert data
	c, sess, _ := CopySessionAndGetCollection(sess, "passwords")
	c.Insert(validDBEntry)
	c.Insert(lockedOutUser)
	c.Insert(unknownLocation)
	c.Insert(wrongAnswer)
	c.Insert(alreadyExistsEntry)
	// test table
	var tableTests = []struct {
		name string
		request *pb.ChallengeRequest
		expectedResponse *pb.ChallengeResponse
		expectedError error
		stringifiedComparison bool
	}{

		{"bad user generate challenge request", &pb.ChallengeRequest{}, nil, errors.New("Invalid request: 'user' is a required field."), false},
		{"error fetching user", &nonExistentUserRequest, nil, errors.New("Error fetching user: no userDn found for: '" + nonExistentUserRequest.User + "'"), false},
		{"locked out user", &lockedOutRequest, nil, errors.New("'locked@out.com' is locked out. Please contact an administrator to regain access."), false},
		{"unknown location", &unknownLocationRequest, &unknownLocationResponse, nil, true},
		{"bad response to question and unknown location", &wrongAnswerRequest, &wrongAnswerResponse, nil, true},
		{"challenge token already exists", &alreadyExistsRequest, nil, errors.New("Challenge request has already been created"), false},
		{"valid request", &validRequest, nil, errors.New("not implemented"), false},
	}

	//cleanup
	defer func() {
		c.RemoveAll(bson.M{"auth.dn": validDBEntry.Auth.Dn})
		c.RemoveAll(bson.M{"auth.dn": lockedOutUser.Auth.Dn})
		c.RemoveAll(bson.M{"auth.dn": unknownLocation.Auth.Dn})
		c.RemoveAll(bson.M{"auth.dn": wrongAnswer.Auth.Dn})
		c.RemoveAll(bson.M{"auth.dn": alreadyExistsEntry.Auth.Dn})
	}()

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			actualResponse, actualError := s.GenerateChallenge(ctx, tt.request)
			if tt.stringifiedComparison {
				AssertEqual(t, fmt.Sprint(actualResponse), fmt.Sprint(tt.expectedResponse))
			} else {
				AssertEqual(t, actualResponse, tt.expectedResponse)
			}
			AssertErrorEqual(t, actualError, tt.expectedError)
		})
	}

	// test error fecthing user
	s.session = mgo.Session{}
	t.Run("Could not get collection", func(t *testing.T) {
		_, actualError := s.GenerateChallenge(ctx, &validRequest)
		AssertErrorEqual(t, actualError, errors.New("Could not get collection: Bad mongo session"))
	})
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