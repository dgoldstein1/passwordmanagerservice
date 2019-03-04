// validateRequests.go

package methods

import (
	"github.com/pkg/errors"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

func ValidateChallengeRequest(request *pb.ChallengeRequest) error {
	// must have 'User'
	if request.User == "" {
		return errors.New("'user' is a required field.")
	}
	// must have 'Location'
	if request.Location == nil {
		return errors.New("'location' is a required field.")
	}
	// must have 'Location.Ip'
	if request.Location.Ip == "" {
		return errors.New("'location.ip' is a required field.")
	}
	if (request.Location.Latitude == 0 || request.Location.Longitude == 0) {
		return errors.New("'location.latitude and location.longitude are required fields")
	}
	if (request.Location.CountryCode == "") {
		return errors.New("'location.countryCode' is a required field.")
	}

	return nil
}