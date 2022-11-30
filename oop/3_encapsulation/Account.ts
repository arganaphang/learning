export class Account {
  name: string; // Default public modiier
  public email: string; // we define it explicitly
  private password: string; // we close the ousider for accessing this attribute, and can be only use inside this class

  constructor(name: string, email: string, password: string) {
    this.name = name;
    this.email = email;
    this.password = password;
  }

}