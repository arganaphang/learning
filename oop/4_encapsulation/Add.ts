import { Arithmetic } from "./Arithmetic.ts";

export class Add implements Arithmetic {
  calculate(a: number, b: number): number {
    return a + b;
  }
}
