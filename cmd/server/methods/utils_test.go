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

func TestGetNewAuthQuestion(t *testing.T) {
	firstQ := "Jolena's least favorite food?"
	secondQ := "Jolena's favorite food?"
	thirdQ := "Jolena's favoriate snack?"
	questions := []*pb.AuthQuestion{
		&pb.AuthQuestion{
			Q : firstQ, 
			A : "pickles",
		},
	}

	questions2 := []*pb.AuthQuestion{
		&pb.AuthQuestion{
			Q: firstQ,
			A : "pickles",
		},
		&pb.AuthQuestion{
			Q: secondQ,
			A: "steak",
		},
	}

	validRequest := &pb.ChallengeRequest{
		UserQuestionResponse : &pb.AuthQuestion{
			Q : firstQ,
			A : "pickles",
		},
	}

	// test table
	var tableTests = []struct {
		name string
		request *pb.ChallengeRequest
		qs []*pb.AuthQuestion
		expectedResult string
	}{
		{"returns same question when length is one", validRequest, questions, questions[0].Q},
		{"returns nothing when length is 0", validRequest, []*pb.AuthQuestion{}, ""},
		{"returns other question when length is 2", validRequest, questions2, secondQ},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := GetNewAuthQuestion(tt.request, tt.qs)
			AssertEqual(t, actualResult, tt.expectedResult)
		})
	}

	questions3 := []*pb.AuthQuestion{
		&pb.AuthQuestion{
			Q: firstQ,
			A : "pickles",
		},
		&pb.AuthQuestion{
			Q: secondQ,
			A: "steak",
		},
		&pb.AuthQuestion{
			Q: thirdQ,
			A: "cheddar cheese",
		},
	}
	t.Run("returns random with 2 + questions", func(t *testing.T) {
		actualResult := GetNewAuthQuestion(validRequest, questions3)
		if (actualResult != secondQ && actualResult != thirdQ) {
			t.Errorf("Expected result to be %s or %s but was %s", secondQ, thirdQ, actualResult)
		} 
	})

	t.Run("returns random with 2 + questions and no current question", func(t *testing.T) {
		actualResult1 := GetNewAuthQuestion(&pb.ChallengeRequest{}, questions3)
		actualResult2 := GetNewAuthQuestion(&pb.ChallengeRequest{}, questions3)
		actualResult3 := GetNewAuthQuestion(&pb.ChallengeRequest{}, questions3)
		actualResult4 := GetNewAuthQuestion(&pb.ChallengeRequest{}, questions3)
		if (actualResult1 == actualResult2 && actualResult2 == actualResult3 && actualResult3 == actualResult4) {
			t.Errorf("Expected results to be different")			
		}
	})


}