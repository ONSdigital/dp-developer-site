
import { toggleCodeExample } from './CodeExamples'
import { fetchData } from './FetchData';
import { toggleResultsView } from './Results';

export default function tourInit () {

// "Call To Action links" 
const jsExample = document.querySelector('[data-tour-example]');
const tryItOut = document.getElementById('tryitout-get-observation');
const endpoint = document.querySelector('[data-tour-endpoint]').dataset.tourEndpoint

// Results 
let results;

const resultsTabs = document.querySelectorAll('[data-tour-results-tab]')
const datasetResults = document.getElementById('tryitout-list-datasets-results');
const jsonResults = document.getElementById('tryitout-list-datasets-json-result');

console.log(resultsTabs)

// State
const url = `https://api.beta.ons.gov.uk/v1${endpoint}`;
const currentPage = tryItOut.dataset.tourPage;

// Event Listeners
jsExample.addEventListener('click', function () {
    toggleCodeExample(url)
})

tryItOut.addEventListener('click', async function () {
    tryItOut.setAttribute('disabled', 'true');
    tryItOut.classList.add('btn--primary-disabled');
    results = await fetchData(url)
    datasetResults.classList.remove('hidden');
    tryItOut.classList.remove('btn--primary-disabled');
})

// Event Listener for toggling radio (results/json) - loop through buttons and apply class depending on which radio button is clicked
resultsTabs.forEach((tab) => {
    tab.addEventListener('click', (e) => {
        toggleResultsView(e.target.dataset.tourResultsTab)
    })
})

// Functionality for displaying JSON and building a table of results

}
