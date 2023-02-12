window.addEventListener('DOMContentLoaded', (event) => {
    let questionData;

    let request = new XMLHttpRequest();
    request.open("GET", "http://localhost:8081/coderTestData.json");
    request.send();
    request.onload = main;

    function main(){
        if (request.status != 200) {
            console.log(`error ${request.status}: ${request.statusText}`);
        } else {
            questionData = JSON.parse(request.responseText)
        }

        fillTest();

        function fillTest(){
            document.body.innerHTML += "<p>"+ questionData.Text + "</p>"
        }
    }




});