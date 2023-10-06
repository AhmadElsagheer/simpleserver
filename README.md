# simpleserver
A simple http server that receives any request and logs it to stdout.

## Installation
```bash
go install github.com/ahmadelsagheer/simpleserver@latest
```

## Usage
1. Start server
```bash
simpleserver -p 8000
```
2. Send a request
```bash
# GET request
curl http://localhost:8000
# POST request
curl -X POST --data '{"hello": "world"}' http://localhost:8000
```

You can use it with [ngrok](https://ngrok.com/) as follows:

1. Start both simple and ngrok servers
```bash
# terminal 1
simpleserver -p 8000
# terminal 2
ngrok http 8000
```
2. Send a request (from any machine)
```bash
curl https://<NGROK_FORWARDING_URL>
```
