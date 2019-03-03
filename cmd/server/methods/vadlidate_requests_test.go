// vadlidate_requests_test.go

package methods

import (
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"testing"
)

func TestValidateChallengeRequest(t *testing.T) {
	// must pass location
	requestWithoutLocation := pb.ChallengeRequest{
		User : "davd@david.com",
	}
	err := ValidateChallengeRequest(&requestWithoutLocation)
	if (err == nil || err.Error() != "'location' is a required field.") {
		t.Errorf("Expected error to be thrown")
	}
	// must pass location.ip
	requestWithoutIp := pb.ChallengeRequest{
		User : "davd@david.com",
		Location : &pb.Location{
			Ip : "",
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
	// returns nil on normal request
	validRequest := pb.ChallengeRequest{
		User : "davd@david.com",
		Location : &pb.Location{
			Ip : "192.0.0.1",
		},
	}
	err = ValidateChallengeRequest(&validRequest)
	if err != nil {
		t.Error(err)
	}
}