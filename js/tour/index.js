
import {toggleCodeExample} from './CodeExamples';
import {fetchData} from './FetchData';
import {buildJSONView, buildResultsView} from './Results';
import {initialiseTabFunctionality} from '../utils/tabs';

export default function tourInit() {
  // Build API URL
  const endpoint = document.querySelector('[data-tour-endpoint]').dataset.tourEndpoint;
  const url = `https://api.beta.ons.gov.uk/v1${endpoint}`;

  // Toggle Code Example
  const exampleContainer = document.querySelector('[data-tour-example-block]');
  const jsExample = document.querySelector('[data-tour-example]');

  jsExample.addEventListener('click', function() {
    toggleCodeExample(exampleContainer, url);
  });

  // Results Views
  const tryItOut = document.querySelector('[data-tour-tryitout=\'observation\']');
  const resultsContainer = document.querySelector('[data-tour-tryitout=\'results\']');
  const resultsOutputContainer = document.querySelector('[data-tour-results-type]');
  const jsonContainer = document.querySelector('[data-tour-results-view=\'json\'] > .markdown > pre');

  tryItOut.addEventListener('click', async function() {
    initialiseTabFunctionality(resultsContainer);
    // Fetch data; disable Try It Out button while waiting for response/error
    tryItOut.setAttribute('disabled', 'true');
    tryItOut.classList.add('btn--primary-disabled');
    const resultsData = await fetchData(url);
    resultsContainer.classList.remove('hidden');
    tryItOut.classList.remove('btn--primary-disabled');

    // Get results type string. If none present, assume current page of tour does not have a Results section and default to 'jsonOnly'
    const resultsType = resultsOutputContainer ? resultsOutputContainer.dataset.tourResultsType : 'jsonOnly';

    buildJSONView(jsonContainer, resultsData);
    buildResultsView(resultsOutputContainer, resultsData, resultsType);
  });
}
