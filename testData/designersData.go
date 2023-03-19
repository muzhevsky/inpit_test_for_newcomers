package testData

import (
	"encoding/json"
	"fmt"
)

type designersData struct {
	DesignerQuestions []*DesignerQuestion
	Descriptions      map[string]Description
}

func GetDesignersData() []byte {
	designerQuestions := make([]*DesignerQuestion, 0)
	answers := make([]*DesignerAnswer, 0)

	// вопрос 1

	answers = append(answers,
		&DesignerAnswer{
			Text: "Сочинением рассказов",
			Dizn: 1,
			Tlvd: 2,
			Rklm: 3,
		},
		&DesignerAnswer{
			Text: "Съемкой и обработкой фото и видео",
			Dizn: 3,
			Tlvd: 3,
			Rklm: 1,
		},
	)
	designerQuestions = append(designerQuestions, &DesignerQuestion{
		Text:    "Чем вам больше нравится заниматься?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*DesignerAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 2

	answers = append(answers,
		&DesignerAnswer{
			Text: "Умею убеждать людей",
			Dizn: 0,
			Tlvd: 0,
			Rklm: 2,
		},
		&DesignerAnswer{
			Text: "У меня есть чувство стиля",
			Dizn: 2,
			Tlvd: 0,
			Rklm: 0,
		},
		&DesignerAnswer{
			Text: "Считаю, что главное в продукте - это внешний вид",
			Dizn: 1,
			Tlvd: 1,
			Rklm: 0,
		},
	)

	designerQuestions = append(designerQuestions, &DesignerQuestion{
		Text:    "Выберите утверждение, луче всего характеризующее вас",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*DesignerAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 3

	answers = append(answers,
		&DesignerAnswer{
			Text: "Занимаюсь профессионально",
			Dizn: 2,
			Tlvd: 1,
			Rklm: 1,
		},
		&DesignerAnswer{
			Text: "Иногда рисую для души",
			Dizn: 1,
			Tlvd: 1,
			Rklm: 1,
		},
		&DesignerAnswer{
			Text: "Рисование не интересует",
			Dizn: 0,
			Tlvd: 2,
			Rklm: 2,
		},
	)
	designerQuestions = append(designerQuestions, &DesignerQuestion{
		Text:    "Как вы относитесь к рисованию",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*DesignerAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 4

	answers = append(answers,
		&DesignerAnswer{
			Text: "Есть. Телевидение будет расти и развиваться.",
			Dizn: 1,
			Tlvd: 2,
			Rklm: 1,
		},
		&DesignerAnswer{
			Text: "Нет. Интернет стремительно забирает аудиторию у телевидения.",
			Dizn: 2,
			Tlvd: 1,
			Rklm: 2,
		},
	)
	designerQuestions = append(designerQuestions, &DesignerQuestion{
		Text:    "Есть ли у телевидения будущее, по вашему мнению?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*DesignerAnswer, 0)
	//////////////////////////////////////////////////////////////////////////////////////

	// вопрос 5

	answers = append(answers,
		&DesignerAnswer{
			Text: "Мощная рекламная кампания, подчеркивающая сильные стороны и скрывающая слабые",
			Dizn: 1,
			Tlvd: 1,
			Rklm: 2,
		},
		&DesignerAnswer{
			Text: "Привлекательная визуальная состовляющая - красивый и удобный интерфейс, переходы и т.д.",
			Dizn: 2,
			Tlvd: 2,
			Rklm: 0,
		},
	)
	designerQuestions = append(designerQuestions, &DesignerQuestion{
		Text:    "Что по вашему больше влияет на коммерческий успех продукта в интернете",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*DesignerAnswer, 0)

	// вопрос 6

	answers = append(answers,
		&DesignerAnswer{
			Text: "Мне нравится не только красота продукта, но и удобство использования",
			Dizn: 2,
			Tlvd: 1,
			Rklm: 1,
		},
		&DesignerAnswer{
			Text: "Для меня главное - красота и содержание продукта",
			Dizn: 1,
			Tlvd: 2,
			Rklm: 2,
		},
	)
	designerQuestions = append(designerQuestions, &DesignerQuestion{
		Text:    "Какое утверждение про вас правдиво?",
		Answers: answers,
		//ImageLink: ""
	})
	answers = make([]*DesignerAnswer, 0)

	// вопрос 7

	answers = append(answers,
		&DesignerAnswer{
			Text: "Отлично",
			Dizn: 2,
			Tlvd: 1,
			Rklm: 1,
		},
		&DesignerAnswer{
			Text: "Хорошо",
			Dizn: 0,
			Tlvd: 2,
			Rklm: 2,
		},
		&DesignerAnswer{
			Text: "Удовлетворительно",
			Dizn: 0,
			Tlvd: 1,
			Rklm: 1,
		},
	)
	designerQuestions = append(designerQuestions, &DesignerQuestion{
		Text:    "Хорошо ли вы учитесь?",
		Answers: answers,
		//ImageLink: ""
	})

	//////////////////////////////////////////////////////////////////////////////////////

	descriptions := make(map[string]Description)
	descriptions["dizn"] = Description{
		Title: "Ваше направление: ДИЗН (Дизайн).",
		Text: "Дизайнер, специализирующийся в сфере графического дизайна, - это профессионал, умеющий разрабатывать и создавать графическую продукцию самого разнообразного направления: " +
			"оформление книг, журналов, газет, компакт-дисков и пр., подготовка рекламных баннеров, вывесок, буклетов, листовок, а также логотипов, визиток, фирменных упаковок и т.д.," +
			" разработка Интернет-страниц, телевизионных и мультимедийных заставок с использованием 2-D и 3-D-графики. " +
			" <br><br>Выпускник может работать по следующим специальностям:" +
			"<ul class=\"text-grey\"><li>креативный директор (арт-директор);</li><li>  дизайнер рекламной продукции;</li><li>дизайнер компьютерной графики;</li><li>специалист по компьютерной анимации;</li>" +
			"<li>  дизайнер трехмерной графики и анимации;</li><li>дизайнер интерьера;</li><li>  дизайнер мультимедиа программ;</li><li>веб-дизайнер;</li><li>специалист по дизайну видеопродукции.</li></ul>" +
			"<br>Желаем успехов!",
	}
	descriptions["rklm"] = Description{
		Title: "Ваше направление: РКЛМ (Реклама)",
		Text: "Выпускники работают в ведущих редакциях СМИ, пресс-службах, государственных и частных компаниях, исследовательских центрах, учреждениях, подразделениях, занятые изучением функционирования коммуникационного пространства, " +
			"департаментах рекламы и связей с общественностью российских и зарубежных компаний. " +
			"<br>Деятельность бакалавров рекламы и связей с общественностью может быть реализована в следующих профессиональных специализациях:" +
			"<ul class=\"text-grey\"><li>менеджер по рекламе и связям с общественностью (PR); </li><li> креативный директор и арт-директор;</li>" +
			"<li>дизайнер рекламной продукции;</li><li>дизайнер компьютерной графики;</li><li>специалист по компьютерной анимации;</li><li>видеомонтажер;</li>" +
			"<li>PR-советник руководителя;</li><li>бренд-менеджер;</li><li>копирайтер;</li><li>медиапланер;</li><li>арт-байер и медиа-байер;</li><li>режиссер видео- и аудиорекламы и др.</li></ul><br> Желаем успехов!",
	}
	descriptions["tlvd"] = Description{
		Title: "Ваше направление: ТЛВД (Телевидение)",
		Text: "Особым спросом у работодателей пользуются универсальные сотрудники, способные разрабатывать творческую основу видеоматериалов и при этом владеющие техническими средствами видеосъемки и компьютерной обработки видеоматериалов. " +
			"Выпускники направления «Телевидение» могут работать в качестве:" +
			"<ul class=\"text-grey\"><li>  тележурналиста;</li><li>корреспондента (репортера);</li><li>  телеведущего;</li><li> теле- и кинорежиссера;</li><li>  редактора теле- и кинопрограмм;</li>" +
			"<li>  видеооператора;</li><li>режиссера видеомонтажа;</li><li>  звукорежиссера;</li><li>продюсера;</li><li>    комментатора и обозревателя;</li><li>    менеджера проектов в сфере кино и телевидения;</li>" +
			"<li>  руководителя телестудии и телекомпании.</li> </ul><br>Желаем успехов!",
	}

	designerData := designersData{
		DesignerQuestions: designerQuestions,
		Descriptions:      descriptions,
	}
	result, err := json.Marshal(designerData)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(string(result))

	return result
}
