
function main() {
  const result = http_go.getJson({
    url: "http://127.0.0.1:8000/api/v1/user/info",
    params: {},
    headers: {
      "Json-Web-Token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTg3OTIxMDYsImlhdCI6MTcxODc5NTcwNiwicGF5bG9hZCI6eyJ1c2VyX2lkIjozfX0-IU6sRq8jrzHV9sdSXievCyBC4rp4gPQtU1kjIEH-pl2GstZGJutB2nrKSb0jWpOsn_Q1Wk7K2V2RX007lwMwZM8BY0jrUKliosIi8CQXPgE64-JnBnZi_i9HCst37oJnDhe1aOXOd7M7kReGnvbMseWaINnEcr0MoyNFpiVn02RE1rBNbjKrgqMAoTIDbcKLexz8yT5b_fbeGhH-TXOOTORwCGQK3CW4FcS2Rh1UVkj4vnn3DM4CMUw2yJBND370UPMEU8sYn_F7b9d_GOh91mV92KdqJg28oyNKQCuD-jvaYaRagsIDt4gsbZYxg"
    }
  })
  console.log(result)
}

main()
