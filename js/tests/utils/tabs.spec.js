import {initialiseTabFunctionality} from '../../utils/tabs';

const container = document.createElement('section');
const tabMarkup = `
<div role="tablist">
<a type="button" role="tab" tabindex="0" aria-controls="json" aria-selected="true">1</a>
<a type="button" role="tab" tabindex="0" aria-controls="results" aria-selected="false">2</a>
</div>
<div id="json" tabindex="0" role="tabpanel" aria-labelledby="json-tab" aria-hidden="false">Content 1</div>
<div id="results" tabindex="0" role="tabpanel" aria-labelledby="results-tab" aria-hidden="true">Content 2</div>
`;

const resetState = () => {
  container.innerHTML = tabMarkup;
};

const rightArrowPress = new KeyboardEvent('keyup', {key: 'rightarrow', keyCode: 39});
const leftArrowPress = new KeyboardEvent('keyup', {key: 'leftarrow', keyCode: 37});

describe('tab functionality', () => {
  describe('default state', () => {
    resetState();
    initialiseTabFunctionality(container);
    test('that tab 1 is the default selected tab', () => {
      expect(container.querySelector('a[aria-selected=true]').innerHTML === '1').toBe(true);
    });

    test('that tab 2 is not selected by default', () => {
      expect(container.querySelector('a[aria-selected=false]').innerHTML === '2').toBe(true);
    });

    test('that Content 1 is the default tab panel displayed', () => {
      expect(container.querySelector('div[aria-hidden=false]').innerHTML === 'Content 1').toBe(true);
    });

    test('that Content 2 is hidden by default', () => {
      expect(container.querySelector('div[aria-hidden=true]').innerHTML === 'Content 2').toBe(true);
    });
  });

  describe('keyboard accessibility', () => {
    test('that on right arrow keyup, the selected tab is 2', () => {
      resetState();
      initialiseTabFunctionality(container);
      const tabList = container.querySelector('[role=tablist]');
      tabList.dispatchEvent(rightArrowPress);
      expect(container.querySelector('a[aria-selected=true]').innerHTML === '2').toBe(true);
    });
    test('that on right arrow keyup, Content 2 is displayed', () => {
      resetState();
      initialiseTabFunctionality(container);
      const tabList = container.querySelector('[role=tablist]');
      tabList.dispatchEvent(rightArrowPress);
      expect(container.querySelector('[aria-hidden=false]').innerHTML === 'Content 2').toBe(true);
    });
    test('that on two right arrow presses, tab 1 is selected', () => {
      resetState();
      initialiseTabFunctionality(container);
      const tabList = container.querySelector('[role=tablist]');
      tabList.dispatchEvent(rightArrowPress);
      tabList.dispatchEvent(rightArrowPress);
      expect(container.querySelector('[aria-selected=true]').innerHTML === '1').toBe(true);
    });
    test('that on left arrow keyup, tab 2 is selected', () => {
      resetState();
      initialiseTabFunctionality(container);
      const tabList = container.querySelector('[role=tablist]');
      tabList.dispatchEvent(leftArrowPress);
      expect(container.querySelector('[aria-selected=true]').innerHTML === '2').toBe(true);
    });
    test('that on two left arrow presses, tab 1 is selected', () => {
      resetState();
      initialiseTabFunctionality(container);
      const tabList = container.querySelector('[role=tablist]');
      tabList.dispatchEvent(leftArrowPress);
      expect(container.querySelector('[aria-selected=true]').innerHTML === '2').toBe(true);
    });
  });
});

