import {chart} from 'highcharts';
import {orderByDate} from '../utils/orderByDate';

const buildJSONView = (response) => {
  const jsonTextContainer = document.querySelector('[data-tour-results-view=\'json\'] > .markdown > pre');
  jsonTextContainer.innerHTML = `<code>${JSON.stringify(response, null, 2)}</code>`;
};

const buildResultsView = (el, response, resultType) => {
  switch (resultType) {
    case 'table':
      buildTableBody(el, response);
      break;
    case 'latestRelease':
      buildText(el, response.links.latest_version.href);
      break;
    case 'singlePoint':
      buildText(el, response.observations[0].observation);
      break;
    case 'chart':
      buildChart(response);
      break;
    case 'jsonOnly':
      return;
  }
};

const buildTableBody = (resultsContainer, data) => {
  const table = resultsContainer;
  data.items.forEach((dataset) => {
    const row = document.createElement('tr');
    row.classList.add('tour-table__body-row');
    const datasetID = document.createElement('td');
    datasetID.classList.add('tour-table__body-cell');
    datasetID.innerText = dataset.id;
    row.appendChild(datasetID);
    const datasetName = document.createElement('td');
    datasetName.classList.add('tour-table__body-cell');
    datasetName.innerText = dataset.title;
    row.appendChild(datasetName);
    table.appendChild(row);
  });
};

const buildText = (resultsContainer, text) => {
  const paragraph = resultsContainer;
  const resultType = resultsContainer.dataset.tourResultsType;

  if (resultType === 'latestRelease') {
    paragraph.innerHTML = `The latest release can be found at <tt>${text}</tt>`;
  } else if (resultType === 'singlePoint') {
    paragraph.innerHTML = `The value of CPIH for the United Kingdom in August 2016 was <tt>${text}</tt>`;
  }
};

const buildChart = (data) => {
  // The relevant element where the chart needs to be added in the html should have an id of 'chart'
  let timeseries = [];
  const maxNumberOfPointsOnChart = 10;

  data.observations
      .map(function(data) {
        // create a new field from the date to allow it to be ordered
        const sort = new Date('1-' + data.dimensions.time.label.replace('-', ' '));
        // build an array with the 3 values we need and convert the value to a number
        const chartdata = [
          data.dimensions.time.label,
          parseFloat(data.observation),
          sort,
        ];
        // add these arrays to the timeseries array
        timeseries.push(chartdata);
      });

  timeseries = timeseries.sort(orderByDate).splice(0, maxNumberOfPointsOnChart);

  chart('chart', {
    series: [
      {
        data: timeseries,
      },
    ],
    navigation: {
      buttonOptions: {
        enabled: false,
      },
    },
    title: {
      text: null,
    },
    yAxis: {
      labels: {
        enabled: false,
      },
      title: {
        enabled: false,
      },
    },
    legend: {
      enabled: false,
    },
    xAxis: {
      type: 'category',
      crosshair: true,
    },
  });
};

export {buildJSONView, buildResultsView};
