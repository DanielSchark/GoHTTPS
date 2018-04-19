# GoHttpsGraphy
Golang web API using TSL and GraphQL

---

## Generate certificate

To actually have a HTTPS connection you need a private and a public key, this example of creating a public key is **only for testing** and **mustn't be used in production.** [More about...](https://www.globalsign.com/en/ssl-information-center/what-is-an-ssl-certificate/)

To generate a public key, run this in your console and in the same directory as your project is in, then type for **`[host]`** your address (without brackets).

For example just use **`--host=localhost`**
```
go run $GOROOT/src/crypto/tls/generate_cert.go --host=[host]
```

Remember the file **`key.pem`** is the private key and **`cert.pem`** is your public key.

## GraphQL
To use GraphQL, use a GraphQL playground ([like this one](https://github.com/graphcool/graphql-playground)) and feel free to query.

```
query {
	users {
		id
		name {first last}
		email
	}
	user(id: 1) {
		name {first last}
		birthdate {month day year}
		gender
		phone
		jobs
		created
		edited
	}
}
```

URL requests like `<host>/graph?query={<query>}` will be supported soon...

### ToGetDone list:
* [x] Implement GraphQL
* [ ] Support HTTPS _(Currently not working cause of issues with certificate)_
* [ ] Finish GraphQL examples
* [ ] Support GraphQL URL query

_Description follows..._
