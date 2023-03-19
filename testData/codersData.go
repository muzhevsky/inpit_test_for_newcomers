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
			Text:  "Быть лидером команды, продумывать план действий",
			Ifst:  3,
			Pinf:  2,
			Pinj:  2,
			Ivcht: 2,
		},
		&CoderAnswer{
			Text:  "Решать сложные и интересные математические задачи",
			Ifst:  3,
			Pinf:  3,
			Pinj:  2,
			Ivcht: 2,
		},
		&CoderAnswer{
			Text:  "Разбираться в устройстве технически сложных устройств",
			Ifst:  1,
			Pinf:  2,
			Pinj:  3,
			Ivcht: 3,
		},
	)
	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "Что вам ближе?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*CoderAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 2

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

	// вопрос 3

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

	// вопрос 4

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

	// вопрос 5

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

	// вопрос 6

	answers = append(answers,
		&CoderAnswer{
			Text:  "Хочу",
			Ifst:  2,
			Pinf:  2,
			Pinj:  3,
			Ivcht: 0,
		},
		&CoderAnswer{
			Text:  "Не хочу",
			Ifst:  1,
			Pinf:  1,
			Pinj:  0,
			Ivcht: 3,
		},
	)
	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "Хотели бы вы контролировать процесс разработки через поиск ошибок в чужом коде?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*CoderAnswer, 0)

	// вопрос 7

	answers = append(answers,
		&CoderAnswer{
			Text:  "Отлично",
			Ifst:  3,
			Pinf:  2,
			Pinj:  1,
			Ivcht: 0,
		},
		&CoderAnswer{
			Text:  "Хорошо",
			Ifst:  1,
			Pinf:  3,
			Pinj:  2,
			Ivcht: 2,
		},
		&CoderAnswer{
			Text:  "Удовлетворительно",
			Ifst:  0,
			Pinf:  0,
			Pinj:  1,
			Ivcht: 1,
		},
		&CoderAnswer{
			Text:  "Плохо",
			Ifst:  0,
			Pinf:  0,
			Pinj:  0,
			Ivcht: 0,
		},
	)
	coderQuestions = append(coderQuestions, &CoderQuestion{
		Text:    "Хорошо ли вы учитесь?",
		Answers: answers,
		//ImageLink: ""
	})

	//////////////////////////////////////////////////////////////////////////////////////

	descriptions := make(map[string]Description)
	descriptions["ifst"] = Description{
		Title: "Ваше направление: ИФСТ (Информационные системы и технологии)",
		Text: "Уникальное сочетание знаний в области компьютерных технологий, программирования, операционных систем, мультмедиа-технологий позволяет выпускникам стать самыми востребованными специалистами на рынке труда не только в России, но и за рубежом." +
			"<br>Выпускник может работать:<br>  <ui><li>Директором по информационным технологиям </li><li>Начальником АСУ и отделов информатизации </li> <li>Разработчиком мобильных решений </li>" +
			"<li>Программистом (.NET, Java, Perl, PHP)</li><li>Администратором и разработчиком баз данных (MS SQL, Oracle, MYSQL)</li><li>Веб-программистом</li><li>Тестировщиком программного обеспечения</li><li>Разработчиком решений 1С</li></ui><br>Желаем успехов!",
	}
	descriptions["ivcht"] = Description{
		Title: "Ваше направление: ИВЧТ (Информатика и вычислительная техника)",
		Text: "Объектами профессиональной деятельности бакалавров являются вычислительные машины, комплексы, системы и сети. Выпускник может работать: <ul class=\"text-grey\"><li>Программистом</li>" +
			"<li>Администратором компьютерных сетей</li><li>Разработчиком приложений баз данных</li><li>Веб-программистом</li><li>Начальником отдела информатизации</li><li>Директором по информационным технологиям</li></ul><br> Желаем успехов!",
	}
	descriptions["pinf"] = Description{
		Title: "Ваше направление: ПИНФ (Прикладная информатика)",
		Text: "Бакалавр в области прикладной информатики - это профессионал, владеющий всем спектром современных информационных технологий, востребованный на рынке ИТ-сектора. " +
			"Выпускник может работать как в любых компаниях, занимающихся разработкой информационных систем, так и в ИТ-отделе любой компании (банке, торговой организации, государственном учреждении и т.д.), где такие системы эксплуатируются:" +
			"<ul class=\"text-grey\">" + "<li>Разработчиком программного обеспечения</li><li>Системным аналитиком</li><li>Веб-программистом</li><li>Разработчиком баз данных</li>" +
			"<li>Разработчиком приложений SAP</li><li>Разработчиком приложений виртуальной реальности</li><li>Разработчиком приложений Умного Дома и Интернет-вещей</li><li>IT-консультантом</li></ul> <br> Желаем успехов!",
	}
	descriptions["pinj"] = Description{
		Title: "Ваше направление: ПИНЖ (Программная инженерия)",
		Text: "Выпускники направления владеют методами и инструментальными средствами управления программными проектами, проектирования, разработки, тестирования и внедрения программных систем." +
			"Выпускник может работать в любой компании, занимающейся разработкой информационных систем: <ul class=\"text-grey\"><li>Руководителем программных проектов</li><li>Архитектором программных систем</li>" +
			"<li>Программистом</li><li>Администратором баз данных</li><li>Веб-программистом</li><li>Системным аналитиком</li><li>Специалистом по тестированию программных систем</li></ul><br> Желаем успехов!",
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
