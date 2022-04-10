export function binary_search(arr: number[], target: number): number | null {
  let low = 0;
  let high = arr.length - 1;

  while (low <= high) {
    const mid = Math.floor((low + high) / 2);
    const item = arr[mid];

    if (item === target) {
      return target;
    }

    if (target < item) {
      high = mid - 1;
    } else {
      low = mid + 1;
    }
  }

  return null;
}
