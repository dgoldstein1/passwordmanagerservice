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
		t.Errorf("Incorrect error thrown %s", err.Error())
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
		t.Errorf("Incorrect error thrown %s", err.Error())
	}
	// must pass user
	requestWithoutUser := pb.ChallengeRequest{
		Location : &pb.Location{},
	}
	err = ValidateChallengeRequest(&requestWithoutUser)
	if (err == nil || err.Error() != "'user' is a required field.") {
		t.Errorf("Incorrect error thrown %s", err.Error())
	}

	err = ValidateChallengeRequest(&pb.ChallengeRequest{
		User : "davd@david.com",
		Location : &pb.Location{
			Ip : "192.0.0.1",
		},
	})
	if (err == nil || err.Error() != "'location.latitude and location.longitude are required fields") {
		t.Errorf("Incorrect error thrown %s", err.Error())
	}
	err = ValidateChallengeRequest(&pb.ChallengeRequest{
		User : "davd@david.com",
		Location : &pb.Location{
			Ip : "128.3.5.1",
			Latitude : 25249.24,
		},
	})
	if (err == nil || err.Error() != "'location.latitude and location.longitude are required fields") {
		t.Errorf("Incorrect error thrown %s", err.Error())
	}
	err = ValidateChallengeRequest(&pb.ChallengeRequest{
		User : "davd@david.com",
		Location : &pb.Location{
			Ip : "128.3.5.1",
			Latitude : 25249.24,
			Longitude : -252.36,
		},
	})
	if (err == nil || err.Error() != "'location.countryCode' is a required field.") {
		t.Errorf("Incorrect error thrown %s", err.Error())
	}

	err = ValidateChallengeRequest(&pb.ChallengeRequest{
		User : "david@david.com",
		Location : &pb.Location{
			Ip : "128.3.5.1",
			Latitude : 25249.24,
			Longitude : -252.36,
			CountryCode : "RU",
		},
	})
	if (err != nil) {
		t.Error(err)
	}
}