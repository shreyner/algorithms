import { selection_sort } from './selection_sort';

describe('Selection sort', () => {
  test('should sorting [4,9,5,3,8]', () => {
    expect(selection_sort([4, 9, 5, 3, 8])).toEqual([3, 4, 5, 8, 9]);
  });

  test('should sorting [5,-1,-3,15]', () => {
    expect(selection_sort([5, -1, -3, 15])).toEqual([-3, -1, 5, 15]);
  });

  test('should sorting with one element [-1]', () => {
    expect(selection_sort([-1])).toEqual([-1]);
  });

  test('should sorting empty []', () => {
    expect(selection_sort([])).toEqual([]);
  });
});
