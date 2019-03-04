// utils_test.go

package methods

import (
	"testing"
	pb "github.com/dgoldstein1/passwordservice/protobuf"
)

func TestStringInArray(t *testing.T) {
	// test table
	var tableTests = []struct {
		name string
		str string
		arr []string
		expectedResult bool
	}{
		{"basic example", "a", []string{"a","b"}, true},
		{"lowercase", "A", []string{"a", "B"}, true},
		{"negative test", "X", []string{"a", "B"}, false},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := StringInArray(tt.str, tt.arr)
			AssertEqual(t, actualResult, tt.expectedResult)
		})
	}

}

func TestAnswerInAuthQuestions(t *testing.T) {
	firstQ := "Jolena's least favorite food?"
	questions := []*pb.AuthQuestion{
		&pb.AuthQuestion{
			Q : firstQ, 
			A : "pickles",
		},
	}

	validRequest := &pb.ChallengeRequest{
		UserQuestionResponse : &pb.AuthQuestion{
			Q : firstQ,
			A : "pickles",
		},
	}

	wrongAnswer := &pb.ChallengeRequest{
		UserQuestionResponse : &pb.AuthQuestion{
			Q : firstQ,
			A : "bananas",
		},
	}

	wrongQuestion := &pb.ChallengeRequest{
		UserQuestionResponse : &pb.AuthQuestion{
			Q : "what is my mother's name?",
			A : "pickles",
		},
	}


	// test table
	var tableTests = []struct {
		name string
		request *pb.ChallengeRequest
		qs []*pb.AuthQuestion
		expectedResult bool
	}{
		{"valid answer to a question", validRequest, questions, true},
		{"answer is not correct", wrongAnswer, questions, false},
		{"empty request", wrongQuestion, questions, false},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := AnswerInAuthQuestions(tt.request, tt.qs)
			AssertEqual(t, actualResult, tt.expectedResult)
		})
	}



}