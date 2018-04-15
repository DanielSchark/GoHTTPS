# GoHTTPS
Go web API using TSL and GraphQL
___

## Public key generator
Generate a public key in the same directory and type for **[host]** your address (without brackets).

For example just use **"localhost"**
```
go run $GOROOT/src/crypto/tls/generate_cert.go --host=[host]
```

Remember the file **key.pem is the private** key and **cert.pem is your public** key.

## Request examples
```
curl --get --url "https://localhost:443" --insecure
```
Using `--insecure` prevents SSL certificate checking, otherwise it wouldn't work, cause this self created SSL certificate is not registered by a certificate Authority.

Now to use GraphQL use the request with this URL param (Just attach it to the URL).
```
?query={User{Id,Name{First,Last}}}
```

_Still working on GraphQL. Description follows..._
