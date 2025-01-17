export function sub(a: number, b: number) {
  if (a < b) {
    throw new Error("a must be greater than b");
  }
  return a - b;
}
