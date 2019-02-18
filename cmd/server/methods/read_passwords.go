package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

// reads passwords
func (s *serverData) ReadPasswords(ctx context.Context, request *pb.CrudRequest) (*pb.CrudResponse, error) {
	return nil, errors.New("not implemented")
}
