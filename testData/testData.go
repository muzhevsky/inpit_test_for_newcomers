package testData

type CoderQuestion struct {
	Text    string
	Answers []*CoderAnswer
}

type CoderAnswer struct {
	Text  string
	Ifst  byte
	Pinf  byte
	Pinj  byte
	Ivcht byte
}
