# hostproxy

A request may have either an absolute or a relative URI below:

## Usage

```
  -p <port>
  	listening port (default 1080)
  -d <domain>
  	domain name (optional, if you need requests using subdomains)
```

## Absolute URI 

```curl``` with the ```-x``` or ```--proxy``` option.

When you've deployed ```hostproxy``` on ```mitm.jp:8080```, you
do

```curl -x mitm.jp:8080 http://www.example.com/```

Then, the request would be sent.

```http
GET http://www.example.com/ HTTP/1.1
Host: www.example.com
User-Agent: curl/8.0.0
Accept: */*
Proxy-Connection: Keep-Alive

```

## Relative URI 

```curl``` without the ```-x``` or ```--proxy``` option or any other case where proxies are not available.

When you've deployed ```hostproxy``` on ```mitm.jp:8080```, you
do

```curl http://mitm.jp:8080/http://www.example.com/```

Then, the request would be sent.

```http
GET /http://inet-ip.info HTTP/1.1
Host: mitm.jp:8080
User-Agent: curl/8.0.0
Accept: */*

```

## Subdomain-based URI
