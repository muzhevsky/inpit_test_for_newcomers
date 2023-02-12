package testData

func NewCoderQuestion() *CoderQuestion {
	result := &CoderQuestion{}
	result.Text = "Какой-то вопрос"

	answer := &CoderAnswer{}
	answer.Text = "Какой-то ответ 1"
	answer.Ifst = 5
	answer.Pinj = 4
	answer.Pinf = 3
	answer.Ivcht = 2
	result.Answers = append(result.Answers, answer)

	answer = &CoderAnswer{}
	answer.Text = "Какой-то ответ 2"
	answer.Ifst = 1
	answer.Pinj = 2
	answer.Pinf = 3
	answer.Ivcht = 4
	result.Answers = append(result.Answers, answer)

	return result
}
