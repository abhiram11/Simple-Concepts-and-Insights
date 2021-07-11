// const assert = require("assert"); // this is default nodejs library
const assert = require("chai").assert;
const sayHello = require("../app").sayHello;

// const app = require("../app");

describe("App", function () {
  it("app should return hello", function () {
    let result = sayHello();
    assert.equal(result, "hello");
  });
});
