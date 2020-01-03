import { toggleCodeExample } from './utils/CodeExamples'
import { Fetcher } from './utils/FetchData'

// Event listeners
const jsExample = document.getElementById('jsExampleLink')
const callToAction = document.getElementById('tryitout-get-observation')
const datasetResults = document.getElementById('tryitout-list-datasets-results')
const jsonResults = document.getElementById('tryitout-list-datasets-json-result')

// Consts
const endpoint = `https://api.beta.ons.gov.uk/v1${jsExample.dataset.endpoint}`
const currentPage = callToAction.dataset.page

// Event listeners
jsExample.addEventListener('click', function() {
    toggleCodeExample(endpoint)
})

callToAction.addEventListener('click', function () {
    Fetcher.get(endpoint)
        .then(response => {
            console.log(response)
        })
        .catch(err => {
            console.error(err)
        })
})
