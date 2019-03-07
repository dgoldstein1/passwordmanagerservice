// utils.go

package methods

import (
	"strings"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
	"math/rand"
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
	if request.Body.UserQuestionResponse == nil {
		return false
	}
	for _, q := range qs {
		if strings.ToLower(q.Q) == strings.ToLower(request.Body.UserQuestionResponse.Q) {
			return strings.ToLower(q.A) == strings.ToLower(request.Body.UserQuestionResponse.A)
		}
	}
	return false
}

// gets a new random auth question
func GetNewAuthQuestion(request *pb.ChallengeRequest, qs []*pb.AuthQuestion) string {
	currQuestion := ""
	if request.Body.UserQuestionResponse != nil {
		currQuestion = request.Body.UserQuestionResponse.Q
	}
	// case by length
	if len(qs) == 0 {
		return ""
	}
	if len(qs) == 1 {
		return qs[0].Q
	}
	// create new temp array without current question
	tempQs := []string{}
	for _, q := range qs {
		if currQuestion != q.Q {
			tempQs = append(tempQs, q.Q)
		}
	}
	// return random index 
	return tempQs[rand.Intn(len(tempQs))]
}