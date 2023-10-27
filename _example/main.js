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

}

main()