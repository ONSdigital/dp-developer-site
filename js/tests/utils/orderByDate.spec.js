/* eslint-disable no-extend-native */
import {orderByDate} from '../../utils/orderByDate';

describe('order by date tests', () => {
  test('that with the testData, 2019-05-01 will be the first element in the array', () => {
    const testData = [[0, 0, new Date('2019-12-01T11:00:00.135Z')], [0, 0, new Date('2019-05-01T11:00:00.135Z')]];
    const sorted = testData.sort(orderByDate);
    const firstEl = sorted[0][2];
    expect(firstEl.toISOString().includes('2019-05-01')).toBe(true);
  });
  test('that the orderByDate function sorts correctly when all array elements are of the same year and month, but different days', () => {
    const testData = [
      [0, 0, new Date('2020-01-25T11:00:00.135Z')],
      [0, 0, new Date('2020-01-18T11:00:00.135Z')],
      [0, 0, new Date('2020-01-29T11:00:00.135Z')],
      [0, 0, new Date('2020-01-07T11:00:00.135Z')],
    ];
    const expectedResults = [
      [0, 0, new Date('2020-01-07T11:00:00.135Z')],
      [0, 0, new Date('2020-01-18T11:00:00.135Z')],
      [0, 0, new Date('2020-01-25T11:00:00.135Z')],
      [0, 0, new Date('2020-01-29T11:00:00.135Z')],
    ];
    const sorted = testData.sort(orderByDate);
    expect(sorted).toEqual(expectedResults);
  });
});
