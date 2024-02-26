async function fetchOpenExplorer(e) {
    try {
        url = e.target.getAttribute("data-url")
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        const data = await response.json();
        console.log(data);
    } catch (error) {
        console.error(error);
    }
}
const openExplorerBtn = document.getElementById("open-explorer");
openExplorerBtn.addEventListener("click", fetchOpenExplorer);
