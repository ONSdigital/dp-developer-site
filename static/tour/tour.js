import { toggleCodeExample } from './utils/CodeExamples'

// "Call To Action links" 
const jsExample = document.getElementById('jsExampleLink')
const tryItOut = document.getElementById('tryitout-get-observation')

// Results 
let results;
const datasetResults = document.getElementById('tryitout-list-datasets-results')
const jsonResults = document.getElementById('tryitout-list-datasets-json-result')

// Consts
const url = `https://api.beta.ons.gov.uk/v1${jsExample.dataset.endpoint}`
const currentPage = tryItOut.dataset.page

// Event listeners
jsExample.addEventListener('click', function () {
    toggleCodeExample(url)
})

tryItOut.addEventListener('click', function () {
    tryItOut.setAttribute('disabled', 'true')
    tryItOut.classList.add('btn--primary-disabled')
    fetch(url).then(response => {
        return response.json();
    }).then(parsedResponse => {
        results = parsedResponse;
        datasetResults.classList.remove('hidden')
    }).catch(err => {
        throw new Error(err)
    })
})

// Event Listener for toggling radio (results/json)

// Functionality for displaying JSON and building a table of results