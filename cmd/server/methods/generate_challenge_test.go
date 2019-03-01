// generate_challenge_test.go

package methods

import (
	"golang.org/x/net/context"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"testing"
)

func TestGenerateChallenge(t *testing.T) {
	// place holder test
	s := serverData{}
	ctx := context.TODO()
	request := pb.ChallengeRequest{}
	_, err := s.GenerateChallenge(ctx, &request)
	if err == nil {
		t.Errorf("Expected error not to be nil")
	}
	expectedError := "not implemented"
	if err.Error() != "not implemented" {
		t.Errorf("Expected error to be %s but was %s", expectedError, err.Error())
	}
}

func TestValidateChallengeRequest(t *testing.T) {
	// must pass location
	requestWithoutLocation := pb.ChallengeRequest{
		User : "davd@david.com",
	}
	err := ValidateChallengeRequest(&requestWithoutLocation)
	if (err == nil || err.Error() != "'Location' is a required field.") {
		t.Errorf("Expected error to be thrown")
	}
	// must pass location.ip
	requestWithoutIp := pb.ChallengeRequest{
		User : "davd@david.com",
		Location : &pb.Location{
			Ip : "172.42.74.6",
		},
	}
	err = ValidateChallengeRequest(&requestWithoutIp)
	if (err == nil || err.Error() != "'location.ip' is a required field.") {
		t.Errorf("Expected error to be thrown")
	}
	// must pass user
	requestWithoutUser := pb.ChallengeRequest{
		Location : &pb.Location{},
	}
	err = ValidateChallengeRequest(&requestWithoutUser)
	if (err == nil || err.Error() != "'user' is a required field.") {
		t.Errorf("Expected error to be thrown")
	}
}