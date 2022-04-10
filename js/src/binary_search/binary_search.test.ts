import { binary_search } from './binary_search';

describe('Binary search', () => {
  test('should find 7 in [1, 2, 3, 4, 5, 6, 7, 8, 9]', () => {
    expect(binary_search([1, 2, 3, 4, 5, 6, 7, 8, 9], 7)).toBe(7);
  });

  test('should find 2 in [1, 2, 3, 4, 5, 6, 7, 8, 9]', () => {
    expect(binary_search([1, 2, 3, 4, 5, 6, 7, 8, 9], 2)).toBe(2);
  });

  test('should find last element in [1, 2, 3, 4, 5, 6, 7, 8, 9]', () => {
    expect(binary_search([1, 2, 3, 4, 5, 6, 7, 8, 9], 9)).toBe(9);
  });

  test('should find first element in [1, 2, 3, 4, 5, 6, 7, 8, 9]', () => {
    expect(binary_search([1, 2, 3, 4, 5, 6, 7, 8, 9], 1)).toBe(1);
  });

  test('should not fount high in [1,2,3,4,5,6,7,8,9]', () => {
    expect(binary_search([1, 2, 3, 4, 5, 6, 7, 8, 9], 15)).toBe(null);
  });

  test('should not fount low in [2,3,4,5,6,7,8,9]', () => {
    expect(binary_search([2, 3, 4, 5, 6, 7, 8, 9], 1)).toBe(null);
  });
});
