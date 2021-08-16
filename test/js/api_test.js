const axios = require('axios')

// ANanotheractual.email@real.company.org
// newnewamazingpassword

axios
    .post('http://localhost:8081/signin', {
        Type: "login",
        Email: "ANanotheractual.email@real.company.org",
        Password: "newnewamazingpassword"
    })
    .then((res) => {
        console.log(res)
    })
    .catch((error) => {
        console.error(error)
    })