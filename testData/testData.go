package testData

type CoderQuestion struct {
	Text      string
	Answers   []*CoderAnswer
	ImageLink string
}

type CoderAnswer struct {
	Text  string
	Ifst  byte
	Pinf  byte
	Pinj  byte
	Ivcht byte
}

type DesignerQuestion struct {
	Text      string
	Answers   []*DesignerAnswer
	ImageLink string
}

type DesignerAnswer struct {
	Text string
	Rklm byte
	Tlvd byte
	Dizn byte
}

type Description struct {
	Title string
	Text  string
}
