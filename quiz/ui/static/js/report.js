const copyButtons = document.querySelectorAll(".copy-result-btn");
copyButtons.forEach(el => el.addEventListener("click", event => {
    let quizWrap = event.target.parentNode.parentNode.parentNode.nextElementSibling;
    if (quizWrap.classList.contains("copy-wrap") == false) {
        return;
    }

    let resultText = quizWrap.getElementsByClassName("result-text");
    if (resultText.length < 1) {
        return;
    }

    let text = "";
    for (let i = 0; i < resultText.length; i++) {
        let t = resultText[i].innerText.replace(/^\s*$(?:\r\n?|\n)/gm, "");
        text = text + t + "\n";
    }

    navigator.clipboard.writeText(`${text}`);

    let alertCopy = document.querySelector("#alert-copy");
    if (!alertCopy) {
        alertCopy = event.target.parentNode.parentNode.parentNode.previousElementSibling;
        if (!alertCopy) {
            return;
        }
        if (alertCopy.classList.contains("alert") == false) {
            return;
        }
    }

    alertCopy.classList.remove("d-none");
    alertCopy.classList.add("show");
    setTimeout(function() {
        alertCopy.classList.remove("show");
        alertCopy.classList.add("d-none");
    }, 2000);
}));




