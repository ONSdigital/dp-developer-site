
import { toggleCodeExample } from './CodeExamples'

export default function tourInit () {

// "Call To Action links" 
const jsExample = document.querySelector('[data-tour-example]');
const tryItOut = document.getElementById('tryitout-get-observation');
const endpoint = document.querySelector('[data-tour-endpoint]').dataset.tourEndpoint

console.log('endpoint', endpoint)

// Results 
let results;
const datasetResults = document.getElementById('tryitout-list-datasets-results');
const jsonResults = document.getElementById('tryitout-list-datasets-json-result');

// State
const url = `https://api.beta.ons.gov.uk/v1${endpoint}`;
const currentPage = tryItOut.dataset.page;
let resultsView = 'json';

// Event listeners
jsExample.addEventListener('click', function () {
    toggleCodeExample(url)
})

tryItOut.addEventListener('click', function () {
    tryItOut.setAttribute('disabled', 'true');
    tryItOut.classList.add('btn--primary-disabled');
    fetch(url).then(response => {
        return response.json();
    }).then(parsedResponse => {
        results = parsedResponse;
        console.log('results', results)
        datasetResults.classList.remove('hidden');
        tryItOut.classList.remove('btn--primary-disabled');
    }).catch(err => {
        throw new Error(err);
    })
})

// Event Listener for toggling radio (results/json) - loop through radio buttons and apply class depending on which radio button is clicked

// Functionality for displaying JSON and building a table of results
}
