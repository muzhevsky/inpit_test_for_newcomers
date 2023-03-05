window.addEventListener('DOMContentLoaded', (event) => {
    coderButton.addEventListener("click", onCoderButtonClick)
    designerButton.addEventListener("click", onDesignerButtonClick)

    function onCoderButtonClick(){
        window.location = "/test/?preferences=coding"
    }

    function onDesignerButtonClick(){
        window.location = "/test/?preferences=design"
    }
});