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
    console.log(currency, currencyConfig)
  }

  test_module.test()

  try {
    test_module.testPanic()
  } catch (err) {
    console.log("error: ", err.message)
  }

  console.log(test_module.testPtr().a)

  console.log(test_module.testNull() === null)

  const func = () => {
    console.log(Math.average([1,2,3]))

    console.log(["a", "b", "c"].reverse())
  }

  func()
}

main()