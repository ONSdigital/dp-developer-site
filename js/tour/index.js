
import { toggleCodeExample } from './CodeExamples'
import { fetchData } from './FetchData';
import { toggleResultsView, buildJSONView, buildResultsView } from './Results';

export default function tourInit() {

    // Build API URL
    const endpoint = document.querySelector('[data-tour-endpoint]').dataset.tourEndpoint
    const url = `https://api.beta.ons.gov.uk/v1${endpoint}`;


    // Toggle Code Example
    const jsExample = document.querySelector('[data-tour-example]');

    jsExample.addEventListener('click', function () {
        toggleCodeExample(url)
    })


    // Results Views
    let results;
    const tryItOut = document.getElementById('tryitout-get-observation');
    const resultsTabs = document.querySelectorAll('[data-tour-results-tab]')
    const datasetResults = document.getElementById('tryitout-list-datasets-results')
    const resultsContainer = document.querySelector('[data-tour-results-type]')

    tryItOut.addEventListener('click', async function () {
        // Fetch data; disable Try It Out button while waiting for response/error
        tryItOut.setAttribute('disabled', 'true');
        tryItOut.classList.add('btn--primary-disabled');
        results = await fetchData(url)
        datasetResults.classList.remove('hidden');
        tryItOut.classList.remove('btn--primary-disabled');

        // Get results type string. If none present, assume current page of tour does not have a Results section and default to 'jsonOnly'
        const resultsType = resultsContainer ? resultsContainer.dataset.tourResultsType : 'jsonOnly';

        buildJSONView(results)
        buildResultsView(results, resultsType)
    })

    // Toggle between Results and JSON
    resultsTabs.forEach((tab) => {
        tab.addEventListener('click', (e) => {
            toggleResultsView(e.target.dataset.tourResultsTab)
        })
    })
}
