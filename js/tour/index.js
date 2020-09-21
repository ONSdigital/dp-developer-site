import {toggleCodeExample} from './CodeExamples';
import {fetchData} from './FetchData';
import {buildJSONView, buildResultsView} from './Results';
import {initialiseTabFunctionality} from '../utils/tabs';
import {setVisibility} from '../utils/setVisibility';

export default function tourInit() {
  // Build API URL
  const endpoint = document.querySelector('[data-tour-endpoint]').dataset.tourEndpoint;
  const url = `https://api.beta.ons.gov.uk/v1${endpoint}`;

  // Toggle Code Example
  const exampleDetails = document.querySelector('[data-tour-example]');
  const exampleLabel = document.querySelector('[data-tour-example-label]');
  const exampleContainer = document.querySelector('[data-tour-example-block]');

  exampleDetails.addEventListener('toggle', function() {
    toggleCodeExample(exampleDetails, exampleLabel, exampleContainer, url);
  });

  // Results Views
  const tryItOut = document.querySelector('[data-tour-tryitout=\'button\']');
  const resultsSection = document.querySelector('[data-tour-tryitout-results]');
  const errorContainer = document.querySelector('[data-tour=\'error\'');
  const resultsContainer = document.querySelector('[data-tour-tryitout=\'results\']');
  const resultsOutputContainer = document.querySelector('[data-tour-results-type]');
  const jsonContainer = document.querySelector('[data-tour-results-view=\'json\'] > .markdown > pre');
  let resultsData;

  tryItOut.addEventListener('click', async function() {
    setVisibility(errorContainer, false);
    initialiseTabFunctionality(resultsContainer);
    // Fetch data; disable Try It Out button while waiting for response/error
    tryItOut.setAttribute('disabled', 'true');
    tryItOut.classList.add('btn--primary-disabled');

    try {
      resultsData = await fetchData(url);
      setVisibility(resultsSection, true);

      // Get results type string. If none present, assume current page of tour does not have a Results section and default to 'jsonOnly'
      const resultsType = resultsOutputContainer ? resultsOutputContainer.dataset.tourResultsType : 'jsonOnly';
      buildJSONView(jsonContainer, resultsData);
      buildResultsView(resultsOutputContainer, resultsData, resultsType);
      setVisibility(resultsContainer, true);
    } catch {
      tryItOut.removeAttribute('disabled');
      tryItOut.classList.remove('btn--primary-disabled');
      setVisibility(errorContainer, true);
    }
  });
}
