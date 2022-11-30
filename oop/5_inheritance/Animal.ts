export abstract class Animal {
  type: string;

  constructor(type: string) {
    this.type = type;
  }

  hello(): void {
    console.log(`Type ${this.type}`);
  }

  abstract eat(): void;
}
