import {setVisibility} from '../../utils/setVisibility';

describe('set visibility tests', () => {
  test('that the element is set to visible', () => {
    const el = document.createElement('div');
    setVisibility(el, true);
    expect(el.classList.contains('hidden')).toBe(false);
  });
  test('that the element is hidden', () => {
    const el = document.createElement('div');
    setVisibility(el, false);
    expect(el.classList.contains('hidden')).toBe(true);
  });
});
