const jsonView = document.querySelector("[data-tour-results-view='json'] > .markdown > pre");
const resultsView = document.querySelector("[data-tour-results-view='results']");
const table = document.getElementById("tryitout-list-datasets-table");

const toggleResultsView = (viewType) => {
    switch(viewType) {
        case 'json':
            jsonView.classList.remove('hidden');
            resultsView.classList.add('hidden');
            break;
        case 'results':
            jsonView.classList.add('hidden');
            resultsView.classList.remove('hidden');
            break;
    }
}

const buildJSONView = (response) => {
    console.log(jsonView)
    jsonView.innerHTML = `<code>${JSON.stringify(response, null, 2)}</code>`
}

const buildResultsView = (response, resultType) => {
    switch(resultType) {
        case 'table':
           buildTableBody(response);
        case 'jsonOnly':
            return;
    }
}

const buildTableBody = (data) => {
    data.items.forEach(dataset => {
        let row = document.createElement("tr")
        let datasetID = document.createElement("td");
        datasetID.innerText = dataset.id;
        row.appendChild(datasetID);
        let datasetName = document.createElement("td");
        datasetName.innerText = dataset.title;
        row.appendChild(datasetName);
        table.appendChild(row);
    })
}

export { toggleResultsView, buildJSONView, buildResultsView }