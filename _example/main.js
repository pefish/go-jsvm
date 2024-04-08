const m = require("./m.js");

function main() {
  m.test();

  console.log(isNaN(1))
  console.log(isNaN(NaN))

  const currencyConfigs = {
    "TRB": {
      BOLL: [72, 3]
    },
    "CYBER": {
      BOLL: [252, 3]
    }
  }

  for (const [currency, currencyConfig] of Object.entries(currencyConfigs)) {
    console.log(currency, JSON.stringify(currencyConfig))
  }

  test_module.test()

  try {
    test_module.testPanic()
  } catch (err) {
    console.log("error: ", err.message)
  }

  console.log(test_module.testPtr().a)

  console.log(test_module.testNull() === null)

  time.sleep(3)

  const func = () => {
    console.log(Math.average([1,2,3]))

    console.log(["a", "b", "c"].reverse())
  }
  func()

  // throw(new Error("test throw"))

  // try {
  //   test_module.testPanic()
  // } catch (err) {
  //   throw(new Error(err.message))
  // }

  throw({message: "err"})  // 如果 throw 一个 object，则 go 那边收到的 [object Object]
}

main()