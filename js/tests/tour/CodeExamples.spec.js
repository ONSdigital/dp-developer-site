import { toggleCodeExample } from '../../tour/CodeExamples';

describe('displaying the code example', () => {
  const exampleContainer = document.createElement('div');
  const label = document.createElement('span');
  const details = document.createElement('details');
  const url = 'https://localhost:8080/';
  details.setAttribute('open', 'open');
  toggleCodeExample(details, label, exampleContainer, url);

  test('that the label displays Hide JavaScript example', () => {
    expect(label.innerText).toBe('Hide JavaScript example');
  });
  test('that the code block has rendered', () => {
    expect(exampleContainer.innerHTML.indexOf('code') !== -1).toBe(false);
  });
});

describe('hiding the code example', () => {
  const exampleContainer = document.createElement('div');
  const label = document.createElement('span');
  const details = document.createElement('details');
  const url = 'https://localhost:8080/';
  details.removeAttribute('open');
  toggleCodeExample(details, label, exampleContainer, url);

  test('that the label displays Show JavaScript example', () => {
    expect(label.innerText).toBe('Show JavaScript example');
  });
});
