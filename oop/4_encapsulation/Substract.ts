import { Arithmetic } from "./Arithmetic.ts";

export class Substract implements Arithmetic {
  calculate(a: number, b: number): number {
    return a - b;
  }
}
