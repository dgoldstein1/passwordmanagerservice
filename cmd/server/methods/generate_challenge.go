package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	"time"
	// "fmt"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"github.com/spf13/viper"
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
	entry, err := GetEntryFromDB(c, request.User)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching user")
	}
	// is user locked out?
	if int(entry.Auth.FailedLogins) > viper.GetInt("max_failed_logins") {
		return nil, errors.New("'" + request.User + "' is locked out. Please contact an administrator to regain access.")
	}
	// location is not known || valid answer to question?
	locationIsNotKnown := StringInArray(request.Location.Ip, entry.Auth.KnownIps) == false
	invalidResponseToAnswer := AnswerInAuthQuestions(request, entry.Auth.AuthQuestions) == false
	if locationIsNotKnown && invalidResponseToAnswer {
		// increment unsuccessful logins
		entry.Auth.FailedLogins = entry.Auth.FailedLogins + 1
		// return new challenge question
		return &pb.ChallengeResponse{
			Error : "Login Unsuccessful",
			UserQuestion : GetNewAuthQuestion(request, entry.Auth.AuthQuestions),
		}, UpdateEntry(c, entry.Auth.Dn, entry)
	}
	// answer already in db?
	if entry.Auth.AccessToken != "" {
		return nil, errors.New("Challenge request has already been created")
	}
	// generate challenge
	challenge := "foo"
	// solve challenge
	accessToken := "bar"
	// add login to list
	login := &pb.Login{
		Location : request.Location,
		Timestamp : int64(time.Now().Unix()),
	}
	entry.Logins = append(entry.Logins, login)
	entry.Auth.AccessToken = accessToken
	// return response
	return &pb.ChallengeResponse{
		Logins : entry.Logins,
		Challenge : challenge,
		User : entry.User,
	}, UpdateEntry(c, entry.Auth.Dn, entry)
}
