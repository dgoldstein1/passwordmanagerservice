package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

// get challenge token
func (s *serverData) GenerateChallenge(ctx context.Context, request *pb.ChallengeRequest) (*pb.ChallengeResponse, error) {
	// validate request
	if err := ValidateChallengeRequest(request); err != nil {
		return nil, errors.Wrap(err, "Invalid request")
	}
	// is user in db?
	c, _, err := CopySessionAndGetCollection(&s.session, "passwords")
	if err != nil {
		return nil, errors.Wrap(err, "Could not get collection")
	}
	_, err = GetEntryFromDB(c, request.User)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching user")
	}

	// is user locked out?
	// location is known || answer is in header?
	// answer already in db?
	// generate challenge
	// add login to list

	
	return nil, errors.New("not implemented")
}