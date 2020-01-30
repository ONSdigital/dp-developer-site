import {toggleCodeExample} from '../../tour/CodeExamples';

describe('displaying the code example', () => {
  const codeContainer = document.createElement('div');
  const linkContainer = document.createElement('p');
  const url = 'https://localhost:8080/';
  codeContainer.classList.add('hidden');
  toggleCodeExample(codeContainer, linkContainer, url);

  test('that the codeContainer is no longer hidden', () => {
    expect(codeContainer.classList.contains('hidden')).toBe(false);
  });
  test('that the linkContainer displays Hide JavaScript example', () => {
    expect(linkContainer.innerText).toBe('Hide JavaScript example');
  });
  test('that the code block has rendered', () => {
    expect(codeContainer.innerHTML.indexOf('code') !== -1).toBe(false);
  });
});

describe('hiding the code example', () => {
  const container = document.createElement('div');
  const text = document.createElement('p');
  const url = 'https://localhost:8080/';
  toggleCodeExample(container, text, url);

  test('that the codeContainer has the hidden class', () => {
    expect(container.classList.contains('hidden')).toBe(true);
  });
  test('that the linkContainer displays Show JavaScript example', () => {
    expect(text.innerText).toBe('Show JavaScript example');
  });
});
