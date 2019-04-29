# webhole - dummy http server for testing

Dumps out incoming requests to stdout.

## Get/compile

Requires a working [go](https://golang.org) setup.

    $ cd webhole
    $ go build
    $ ./webhole

Runs on port 8080 by default. Use `-a` to run on another address, eg:

    $ ./webhole -a :1234

Try sending it a gzipped request:

    $ echo fookbarwibble | gzip >wibble.gz
    $ curl -X POST -H "Content-Encoding: gzip" --data-binary @wibble.gz http://localhost:8080/foo

You should see output something like:
```
==============================
HTTP/1.1 POST /foo

Accept: */*
Content-Encoding: gzip
Content-Length: 34
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.64.0

fookbarwibble

(14 bytes in body)
```

