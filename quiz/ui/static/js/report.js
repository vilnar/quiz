const copyButtons = document.querySelectorAll(".copy-result-btn");
const alertCopy = document.querySelector("#alert-copy");
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
    alertCopy.classList.remove("d-none");
    alertCopy.classList.add("show");
    setTimeout(function() {
        alertCopy.classList.remove("show");
        alertCopy.classList.add("d-none");
    }, 2000);
}));




