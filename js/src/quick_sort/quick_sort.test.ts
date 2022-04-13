import { quick_sort } from './quick_sort';

describe('Quick Sort', () => {
  test("shoud don't sort base case []", () => {
    expect(quick_sort([])).toEqual([]);
  });

  test('shoud sort with one element', () => {
    expect(quick_sort([1])).toEqual([1]);
  });

  test('shoud sort with two elements [2,1]', () => {
    expect(quick_sort([2, 1])).toEqual([1, 2]);
  });

  test('shoud sort with elements [3,15,-1,4,5,6,10,8,11]', () => {
    expect(quick_sort([3, 15, -1, 4, 5, 6, 10, 8, 11])).toEqual([
      -1, 3, 4, 5, 6, 8, 10, 11, 15,
    ]);
  });
});
