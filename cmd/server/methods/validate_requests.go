// validateRequests.go

package methods

import (
	"github.com/pkg/errors"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

func ValidateChallengeRequest(request *pb.ChallengeRequest) error {
	return errors.New("not implemented")
}