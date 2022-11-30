import { Arithmetic } from "./Arithmetic.ts";
import { Add } from "./Add.ts";
import { Substract } from "./Substract.ts";
import { Multiply } from "./Multiply.ts";
import { Divide } from "./Divide.ts";

let operation: Arithmetic;
operation = new Add();
console.log(`Add ${operation.calculate(1, 2)}`);
operation = new Substract();
console.log(`Substract ${operation.calculate(5, 3)}`);
operation = new Multiply();
console.log(`Multiply ${operation.calculate(4, 3)}`);
operation = new Divide();
console.log(`Divide ${operation.calculate(12, 2)}`);
