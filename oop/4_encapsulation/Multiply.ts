import { Arithmetic } from "./Arithmetic.ts";

export class Multiply implements Arithmetic {
  calculate(a: number, b: number): number {
    return a * b;
  }
}
