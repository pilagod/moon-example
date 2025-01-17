import { add } from "packages/add";

describe("add", () => {
  it("1 + 1 = 2", () => {
    expect(add(1, 1)).toBe(2);
  });
});
