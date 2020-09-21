import {chart} from 'highcharts';
import {orderByDate} from '../utils/orderByDate';

const buildJSONView = (container, response) => {
  container.innerHTML = `<code>${JSON.stringify(response, null, 2)}</code>`;
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
      buildChart(el, response);
      break;
    case 'jsonOnly':
      return;
  }
  window.location.hash = 'results-section';
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
    paragraph.innerHTML = `The latest release can be found at <code>${text}</code>`;
  } else if (resultType === 'singlePoint') {
    paragraph.innerHTML = `The value of CPIH for the United Kingdom in August 2016 was <code>${text}</code>`;
  }
};

const buildChart = (chartContainer, data) => {
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

  chart(chartContainer, {
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
