import { Arithmetic } from "./Arithmetic.ts";

export class Divide implements Arithmetic {
  calculate(a: number, b: number): number {
    if (b === 0) {
      throw new Error("Can't be divide with 0");
    }
    return a / b;
  }
}
