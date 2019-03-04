// utils.go

package methods

import (
	"strings"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

// checks if string is in given array
func StringInArray(a string, arr []string) bool {
	for _, s := range arr {
		if strings.ToLower(a) == strings.ToLower(s) {
			return true
		}
	}
	return false
}

// checks if given answer is correct in database
func AnswerInAuthQuestions(request *pb.ChallengeRequest, qs []*pb.AuthQuestion) bool {
	if request.UserQuestionResponse == nil {
		return false
	}
	for _, q := range qs {
		if strings.ToLower(q.Q) == strings.ToLower(request.UserQuestionResponse.Q) {
			return strings.ToLower(q.A) == strings.ToLower(request.UserQuestionResponse.A)
		}
	}
	return false
}