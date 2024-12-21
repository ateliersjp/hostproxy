# hostproxy

A request may have either an absolute or a relative URI below:

## Usage

```
  -p <port>
  	listening port (default 1080)
  -d <domain>
  	domain name (optional, if you need to use subdomains)
```

## Absolute URI 

```curl``` with the ```-x``` or ```--proxy``` option.

When you've deployed ```hostproxy -d mitm.jp -p 1080```, you do

```curl -x mitm.jp http://www.example.com/```

## Relative URI 

```curl``` without the ```-x``` or ```--proxy``` option or any other case where proxies are not available.

When you've deployed ```hostproxy -d mitm.jp -p 80```, you do

```curl http://mitm.jp/http://www.example.com/```
