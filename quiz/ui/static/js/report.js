let copyButtons = document.querySelectorAll(".copy-result-btn");
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
    // console.log(text);
    navigator.clipboard.writeText(`${text}`);
}));

