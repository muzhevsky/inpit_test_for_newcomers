package testData

import (
	"encoding/json"
	"fmt"
)

type codersData struct {
	CoderQuestions []*CoderQuestion
	Descriptions   map[string]Description
}

func GetCodersData() []byte {
	coderQuestions := make([]*CoderQuestion, 0)
	answers := make([]*CoderAnswer, 0)

	// вопрос 1
	answers = append(answers,
		&CoderAnswer{
			Text:  "Разработка приложений",
			Ifst:  3,
			Pinf:  3,
			Pinj:  1,
			Ivcht: 2,
		},
		&CoderAnswer{
			Text:  "Руководство процессом разработки",
			Ifst:  2,
			Pinf:  2,
			Pinj:  3,
			Ivcht: 1,
		},
		&CoderAnswer{
			Text:  "Работа с сетями и \"железом\"",
			Ifst:  1,
			Pinf:  1,
			Pinj:  1,
			Ivcht: 3,
		},
	)

	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "С чем связана работа вашей мечты?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*CoderAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 2

	answers = append(answers,
		&CoderAnswer{
			Text:  "Веб-сайты и веб-приложения",
			Ifst:  3,
			Pinf:  2,
			Pinj:  0,
			Ivcht: 1,
		},
		&CoderAnswer{
			Text:  "Игры и приложения виртуальной и дополненной реальности",
			Ifst:  2,
			Pinf:  3,
			Pinj:  0,
			Ivcht: 1,
		},
		&CoderAnswer{
			Text:  "Базы данных",
			Ifst:  3,
			Pinf:  2,
			Pinj:  0,
			Ivcht: 2,
		},
		&CoderAnswer{
			Text:  "Технологии умного дома",
			Ifst:  0,
			Pinf:  3,
			Pinj:  0,
			Ivcht: 0,
		},
	)
	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "В разработке какого вида программных продуктов вы бы хотели участвовать?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*CoderAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 3

	answers = append(answers,
		&CoderAnswer{
			Text:  "Важно",
			Ifst:  2,
			Pinf:  3,
			Pinj:  2,
			Ivcht: 1,
		},
		&CoderAnswer{
			Text:  "Могу работать и с консолью",
			Ifst:  2,
			Pinf:  0,
			Pinj:  0,
			Ivcht: 3,
		},
	)
	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "Важно ли видеть красивое форматирование результата вашей работы?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*CoderAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 4

	answers = append(answers,
		&CoderAnswer{
			Text:  "Хочу",
			Ifst:  2,
			Pinf:  2,
			Pinj:  0,
			Ivcht: 0,
		},
		&CoderAnswer{
			Text:  "Не хочу",
			Ifst:  1,
			Pinf:  0,
			Pinj:  2,
			Ivcht: 2,
		},
	)
	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "Хотите ли освоить двухмерную и трехмерную графику?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*CoderAnswer, 0)

	// вопрос 5

	answers = append(answers,
		&CoderAnswer{
			Text:  "Отлично",
			Ifst:  3,
			Pinf:  2,
			Pinj:  2,
			Ivcht: 2,
		},
		&CoderAnswer{
			Text:  "Хорошо",
			Ifst:  2,
			Pinf:  3,
			Pinj:  2,
			Ivcht: 2,
		},
		&CoderAnswer{
			Text:  "Удовлетворительно",
			Ifst:  1,
			Pinf:  2,
			Pinj:  3,
			Ivcht: 3,
		},
	)
	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "Как хорошо вы учитесь?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*CoderAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	descriptions := make(map[string]Description)
	descriptions["ifst"] = Description{
		Title: "Ваше направление: ИФСТ.",
		Text: "Это самое престижное направление ИНПиТа, связанное с программированием, разработкой баз данных, " +
			"администрированием сетей и многими другими видами деятельности в IT. никальное сочетание знаний в области " +
			"компьютерных технологий, программирования,  операционных систем, мультмедиа-технологий позволяет выпускникам стать" +
			" самыми востребованными специалистами на рынке труда не только в России, но и за рубежом. <br> " +
			"Желаем успехов!",
	}
	descriptions["ivcht"] = Description{
		Title: "Ваше направление: ИВЧТ.",
		Text:  "Описание потом придумаю",
	}
	descriptions["pinf"] = Description{
		Title: "Ваше направление: ПИНФ.",
		Text:  "Описание pinf потом придумаю",
	}
	descriptions["pinj"] = Description{
		Title: "Ваше направление: ПИНЖ.",
		Text:  "Описание pinj потом придумаю",
	}

	codersData := codersData{
		CoderQuestions: coderQuestions,
		Descriptions:   descriptions,
	}
	result, err := json.Marshal(codersData)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(string(result))

	return result
}
