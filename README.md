# HTTP(s) middleware go

http middleware handler

## Install

```shell
go get github.com/yosgo-opensource/http-middleware-go
```

## Method

- HttpCORS

```
Access-Control-Allow-Origin: * [default]
Access-Control-Allow-Headers: Content-Type [default]
```

- HttpJSONResponse

```
Content-Type: text/html [GET][default]
Content-Type: application/json [POST][default]
```

## Useage

### HttpCORS

```golang
http.Handle("/route",
  middleware.HttpCORS(
    http.HandlerFunc(watermarkHandler)
  )
)
```

### HttpJSONReqponse

```golang
http.Handle("/route",
  middleware.HttpJSONResponse(
    http.HandlerFunc(watermarkHandler)
  )
)
```
