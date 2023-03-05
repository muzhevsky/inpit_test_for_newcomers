package testData

import (
	"encoding/json"
	"fmt"
)

type codersData struct {
	CoderQuestions []*CoderQuestion
}

func GetCodersData() []byte {
	coderQuestions := make([]*CoderQuestion, 0)
	answers := make([]*CoderAnswer, 0)

	answers = append(answers,
		&CoderAnswer{
			Text:  "ответ1",
			Ifst:  0,
			Pinf:  0,
			Pinj:  0,
			Ivcht: 0,
		},
		&CoderAnswer{
			Text:  "ответ2",
			Ifst:  0,
			Pinf:  0,
			Pinj:  0,
			Ivcht: 0,
		},
	)

	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "вопрос",
		Answers: answers,
	})

	codersData := codersData{CoderQuestions: coderQuestions}
	result, err := json.Marshal(codersData)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(string(result))

	return result
}
