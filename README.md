# Proxy-server
Proxy server for 3rd party services


## Example
**POST** *HOST/proxy-server/v1/proxy/send*
```
curl --location --request POST 'localhost:8080/proxy-server/v1/proxy/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "method":"GET",
    "url":"https://catfact.ninja/fact",
    "header":{
        "User-Agent":"proxy-server"
    }

}'
```

```curl
curl --location --request POST 'localhost:8080/proxy-server/v1/proxy/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "method": "POST",
    "url": "https://httpbin.org/post",
    "body": {
        "name": "Alias",
        "message": "Hello"
    },
    "header": {
        "User-Agent": "proxy-server",
        "Content-Type": "application/json"
    }
}'
```

**GET** *HOST/proxy-server/v1/proxy/get-by-id/{id}*

```
curl --location --request GET 'localhost:8080/proxy-server/v1/proxy/get-by-id/603619bc-8f52-4a5a-9e43-33433cf85300'
```

**GET** *HOST/proxy-server/v1/proxy/all*

```
curl --location 'localhost:8080/proxy-server/v1/proxy/all'
```

## Build and run
```
docker build -t proxy-server .
```
```
docker run proxy-server
```
