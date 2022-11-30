import { Account } from "./Account.ts";

const john = new Account("John Doe", "johndoe@mail.com", "we_set_this_password");

console.log(`Name ${john.name}`)
// Gonna ignore 🫣, Don't do this in your project
// deno-lint-ignore ban-ts-comment
// @ts-ignore
console.log(`Password ${john.password}`) // This gonna be error, because we can't access password, since password attribute is private