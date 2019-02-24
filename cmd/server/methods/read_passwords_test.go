// read_passwords_test.go

package methods

import (
	"golang.org/x/net/context"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"testing"
)

func TestReadPasswords(t *testing.T) {
	// place holder test
	s := serverData{}
	ctx := context.TODO()
	request := pb.CrudRequest{}
	_, err := s.ReadPasswords(ctx, &request)
	if err == nil {
		t.Errorf("Expected error not to be nil")
	}
	expectedError := "not implemented"
	if err.Error() != "not implemented" {
		t.Errorf("Expected error to be %s but was %s", expectedError, err.Error())
	}
}