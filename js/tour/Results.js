const jsonView = document.querySelector("[data-tour-results-view='json']");
const resultsView = document.querySelector("[data-tour-results-view='results']");


export const toggleResultsView = (viewType) => {
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