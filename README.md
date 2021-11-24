# URLCollector

## Description
A NASA picture URLs collector API.

## Launch
```shell
$ make make docker-run-dev
```

# Stop
```shell
$ make docker-stop-dev
```

# Restart
```shell
$ make docker-start-dev
```

# Sample output
Happy path
```json
{
  "urls": [
    "https://apod.nasa.gov/apod/image/2001/BetelgeuseImagined_EsoCalcada_960.jpg",
    "https://apod.nasa.gov/apod/image/2001/OrionTrees123019Westlake1024.jpg"
  ]
}
```

Error
```json
{
  "error": "'from' date should be prior to 'to' date, 'to' date should be after 'from' date"
}
```

## Notes
From NASA API docs
```text
In documentation examples, the special DEMO_KEY api key is used.
This API key can be used for initially exploring APIs prior to signing up, but it has much lower rate limits, so you’re encouraged to signup for your own API key if you plan to use the API (signup is quick and easy). The rate limits for the DEMO_KEY are:

  Hourly Limit: 30 requests per IP address per hour
  Daily Limit: 50 requests per IP address per day
```


# TODO
* Implement context logger
* Implement request tracing
* Implement a cache to avoid retrieving the same URLs multiple times
* Implement liveness and readiness probe
* Implement a hypothetical (PoC) mechanism to make calls to the mediadownloader service (gRPC?)
* Add tests
