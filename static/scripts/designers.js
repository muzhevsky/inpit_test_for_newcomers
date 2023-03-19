window.addEventListener('DOMContentLoaded', (event) => {
    let container = document.getElementById("container");

    function userAnswer() {
        this.dizn = 0;
        this.tlvd = 0;
        this.rklm = 0;
    }

    let usersAnswers = [];
    let testData;
    let request = new XMLHttpRequest();
    request.open("GET", "http://localhost:8081/test/designersData.json");
    request.send();
    request.onload = main;

    function main(){
        if (request.status != 200) {
            console.log(`error ${request.status}: ${request.statusText}`);
        }
        else {
            testData = JSON.parse(request.responseText);
            testData.currentQuestion = 0;
            testData.length = testData.DesignerQuestions.length;
            console.log(testData);
            startTest();
        }
    }

    function startTest(){
        showQuestion();
        usersAnswers.push(new userAnswer());

        document.querySelector(".progressBarValue").width = 0;
    }

    function nextQuestion(){
        let selected = document.querySelector("input[name = \"question\"]:checked").value;
        let selectedAnswer = testData.DesignerQuestions[testData.currentQuestion].Answers[selected];

        clearContainer();

        usersAnswers[usersAnswers.length-1].dizn += selectedAnswer.Dizn;
        usersAnswers[usersAnswers.length-1].tlvd += selectedAnswer.Tlvd;
        usersAnswers[usersAnswers.length-1].rklm += selectedAnswer.Rklm;

        testData.currentQuestion++;
        //TODO: прогресс бар
        //progressBar.width = progressBar.parentElement.width * questionData.currentQuestion / questionData.length;
        if (testData.currentQuestion < testData.length) {
            showQuestion();
            usersAnswers.push(new userAnswer());
        }
        else showResult();
    }

    function showQuestion(){
        let answers = "";
        let currentQuestion = testData.DesignerQuestions[testData.currentQuestion];

        for (let i = 0; i < currentQuestion.Answers.length; i++){
            answers += "<div class=\"radioContainer\">\n" +
                "    <input id=\"answer"+ i + "\" class=\"radioButton\" type=\"radio\" name=\"question\" value=\""+i+"\">\n" +
                "    <label for=\"answer"+ i + "\">" + currentQuestion.Answers[i].Text + "</label>\n" +
                "    <br>\n" +
                "</div>\n";
        }

        container.innerHTML =
            "        <div class=\"questionHeader\">" + currentQuestion.Text + "</div>\n" +
            "        <div class=\"separator\"></div>" +
            "        <form class=\"answersForm\">\n" +

            answers + "\n" +

            "        <input id=\"nextQuestion\" type=\"button\" value=\"Далее\">\n" +
            "        </form>\n" +
            "        <div class=\"progressBar\">\n" +
            "            <div class=\"progressBarValue\" width=\"0px\"></div>\n" +
            "        </div>"

        document.querySelector("#nextQuestion").addEventListener("click", (ev) =>
        {
            let selected = document.querySelector("input[name = \"question\"]:checked");
            if (selected != null) nextQuestion();
        });
    }

    function showResult(){
        let result = new Map()
        result.set("dizn", 0);
        result.set("rklm", 0);
        result.set("tlvd", 0);


        for (let i = 0; i < usersAnswers.length; i++){
            result.set("dizn", result.get("dizn") + usersAnswers[i].dizn);
            result.set("rklm", result.get("rklm") + usersAnswers[i].rklm);
            result.set("tlvd", result.get("tlvd") + usersAnswers[i].tlvd);
        }

        let maxValue = 0;
        let maxKey = "";
        result.forEach((value, key, map) => {
            if (maxValue < value){
                maxValue = value;
                maxKey = key;
            }
        });

        let description;

        console.log(maxKey);

        switch (maxKey){
            case "dizn":
                description = testData.Descriptions.dizn;
                break
            case "tlvd":
                description = testData.Descriptions.tlvd;
                break
            case "rklm":
                description = testData.Descriptions.rklm;
                break
        }

        clearContainer();

        container.innerHTML =
            "        <div class=\"questionHeader\">" + description.Title + "</div>\n" +
            "        <div class=\"separator\"></div>" +
            "        <div class=\"description\">" + description.Text + "</div>";
    }

    function clearContainer(){
        while (container.firstChild) {
            container.removeChild(container.firstChild);
        }
    }
});