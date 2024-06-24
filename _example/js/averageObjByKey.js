
function main() {
  const objs = [
    {
      "value": 1
    },
    {
      "value": 2.0
    },
    {
      "value": "a"
    }
  ]
  // console.log(Math.averageObjByKey(objs, "value"))
  try {
    console.log(Math.averageObjByKey(objs, "value"))
  } catch (err) {
    console.log(err)
  }
}

main()
