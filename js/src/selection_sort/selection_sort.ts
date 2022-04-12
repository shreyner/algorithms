function findMinElement(arr: number[]): number {
  let indexMinElement = 0;

  for (let i = 1; i < arr.length; i++) {
    indexMinElement = arr[indexMinElement] > arr[i] ? i : indexMinElement;
  }

  return indexMinElement;
}

export function selection_sort(arr: number[]): number[] {
  const sortingArray = [...arr];
  const sortedArray = [];
  const length = sortingArray.length;

  for (let i = 0; i < length; i++) {
    const indexMinElement = findMinElement(sortingArray);

    sortedArray.push(...sortingArray.splice(indexMinElement, 1));
  }

  return sortedArray;
}
