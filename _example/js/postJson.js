
function main() {
  const result = http_go.postJson({
    url: "http://127.0.0.1:8000/api/v1/login",
    params: {
      "username": "lai",
      "password": ""
    }
  })
  console.log(result)
}

main()
