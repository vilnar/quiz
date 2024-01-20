async function fetchOpenExplorerDbDumpDir() {
    try {
        const response = await fetch('/admin/open-explorer-dbdumpdir');
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        console.log(data);
    } catch (error) {
        console.error(error);
    }
}

var openExplorerBtn = document.getElementById('open-explorer-dbdumpdir');
openExplorerBtn.addEventListener('click', fetchOpenExplorerDbDumpDir);
