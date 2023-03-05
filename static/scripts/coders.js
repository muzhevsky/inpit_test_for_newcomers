window.addEventListener('DOMContentLoaded', (event) => {
    let container = document.getElementById("container");

    function userAnswer() {
        this.ifst = 0;
        this.ivcht = 0;
        this.pinf = 0;
        this.pinj = 0;
    }

    let usersAnswers = [];
    let testData;
    let request = new XMLHttpRequest();
    request.open("GET", "http://localhost:8081/test/testData.json");
    request.send();
    request.onload = main;

    function main(){
        if (request.status != 200) {
            console.log(`error ${request.status}: ${request.statusText}`);
        }
        else {
            testData = JSON.parse(request.responseText);
            testData.currentQuestion = 0;
            testData.length = testData.CoderQuestions.length;
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
        let selectedAnswer = testData.CoderQuestions[testData.currentQuestion].Answers[selected];

        clearContainer();

        usersAnswers[usersAnswers.length-1].ifst += selectedAnswer.Ifst;
        usersAnswers[usersAnswers.length-1].ivcht += selectedAnswer.Ivcht;
        usersAnswers[usersAnswers.length-1].pinf += selectedAnswer.Pinf;
        usersAnswers[usersAnswers.length-1].pinj += selectedAnswer.Pinj;

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
        let currentQuestion = testData.CoderQuestions[testData.currentQuestion];

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
        result.set("ifst", 0);
        result.set("ivcht", 0);
        result.set("pinf", 0);
        result.set("pinj", 0);


        for (let i = 0; i < usersAnswers.length; i++){
            result.set("ifst", result.get("ifst") + usersAnswers[i].ifst);
            result.set("ivcht", result.get("ivcht") + usersAnswers[i].ivcht);
            result.set("pinf", result.get("pinf") + usersAnswers[i].pinf);
            result.set("pinj", result.get("pinj") + usersAnswers[i].pinj);
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

        switch (maxKey){
            case "ifst":
                description = testData.Descriptions.ifst;
                break
            case "pinf":
                description = testData.Descriptions.pinf;
                break
            case "ivcht":
                description = testData.Descriptions.ivcht;
                break
            case "pinj":
                description = testData.Descriptions.pinj;
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