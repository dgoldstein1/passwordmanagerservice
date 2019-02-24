package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

// get challenge token
func (s *serverData) GenerateChallenge(ctx context.Context, request *pb.ChallengeRequest) (*pb.ChallengeResponse, error) {

	
	return nil, errors.New("not implemented")
}
