const jsonView = document.querySelector("[data-tour-results-view='json'] > .markdown > pre");
const resultsView = document.querySelector("[data-tour-results-view='results']");

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
    jsonView.innerHTML = `<code>${JSON.stringify(response, null, 2)}</code>`
}

const buildResultsView = (el, response, resultType) => {
    switch(resultType) {
        case 'table':
           buildTableBody(el, response);
           break;
        case 'latestRelease':
            buildText(el, response.links.latest_version.href)
            break;
        case 'singlePoint':
            buildText(el, response.observations[0].observation)
            break;
        case 'chart':
            buildChart(el, response)
            break;
        case 'jsonOnly':
            return;
    }
}

const buildTableBody = (resultsContainer, data) => {
    let table = resultsContainer;
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

const buildText = (resultsContainer, text) => {
    let paragraph = resultsContainer;
    let resultType = resultsContainer.dataset.tourResultsType

    if (resultType === 'latestRelease') {
        paragraph.innerHTML = `The latest release can be found at <tt>${text}</tt>`
    } else if (resultType === 'singlePoint') {
        paragraph.innerHTML = `The value of CPIH for the United Kingdom in August 2016 was <tt>${text}</tt>`
    }
}

const buildChart = (resultsContainer, data) => {

}

export { toggleResultsView, buildJSONView, buildResultsView }