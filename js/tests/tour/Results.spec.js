import {buildJSONView, buildResultsView} from '../../tour/Results';
import chartResponse from './chartResponse';

describe('building JSON view', () => {
  const jsonResponse = '{ Test JSON response }';
  const el = document.createElement('div');
  buildJSONView(el, jsonResponse);
  test('that the JSON response is built', () => {
    expect(el.innerHTML).toBe(`<code>"{ Test JSON response }"</code>`);
  });
});

describe('building the results view', () => {
  describe('building the table results view', () => {
    const el = document.createElement('div');
    const resultType = 'table';
    const response = {
      items: [
        {
          id: 1,
          title: 'Test',
        },
      ],
    };
    test('that the table results is rendered', () => {
      buildResultsView(el, response, resultType);
      expect(el.innerHTML.includes('table')).toBe(true);
    });
  });

  describe('building the latest release results view', () => {
    const el = document.createElement('div');
    el.setAttribute('data-tour-results-type', 'latestRelease');
    const resultType = 'latestRelease';
    const response = {
      links: {
        latest_version: {
          href: 'testURL',
        },
      },
    };
    test('that the latest release text is rendered', () => {
      buildResultsView(el, response, resultType);
      expect(el.innerHTML).toBe(`The latest release can be found at <tt>${response.links.latest_version.href}</tt>`);
    });
  });
  describe('building the single point results view', () => {
    const el = document.createElement('div');
    el.setAttribute('data-tour-results-type', 'singlePoint');
    const resultType = 'singlePoint';
    const response = {
      observations: [
        {
          observation: 'single point test',
        },
      ],
    };
    test('that the single point text is rendered', () => {
      buildResultsView(el, response, resultType);
      expect(el.innerHTML)
          .toBe(`The value of CPIH for the United Kingdom in August 2016 was <tt>${response.observations[0].observation}</tt>`);
    });
  });
  describe('building the chart results view', () => {
    const el = document.createElement('div');
    const resultType = 'chart';
    test('that the chart is rendered', () => {
      buildResultsView(el, chartResponse, resultType);
      expect(el.innerHTML.includes('highcharts')).toBe(true);
    });
  });
});
