package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

// get challenge token
func (s *serverData) GenerateChallenge(ctx context.Context, request *pb.ChallengeRequest) (*pb.ChallengeResponse, error) {
	// validate request
	// is user in db?
	// is user locked out?
	// location is known || answer is in header?
	// answer already in db?
	// generate challenge
	// add login to list

	
	return nil, errors.New("not implemented")
}