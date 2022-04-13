export function quick_sort(arr: number[]): number[] {
  if (arr.length < 2) {
    return arr;
  }

  const pivot = arr[0];
  const less = [];
  const more = [];

  for (let i = 1; i < arr.length; i++) {
    const element = arr[i];

    if (element <= pivot) {
      less.push(element);
    } else {
      more.push(element);
    }
  }

  return [...quick_sort(less), pivot, ...quick_sort(more)];
}
