const m = require("./_example/js/m.js");

function main() {
  m.test();

  console.log(isNaN(1))
  console.log(isNaN(NaN))

  console.log(Math.max(3,4,1))

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

  time_go.sleep(1)

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