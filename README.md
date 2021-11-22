# URLCollector

# Description
[TODO: Description]

# Sample output
Happy path
```json
{
  "urls": [
    "https://apod.nasa.gov/apod/image/2001/BetelgeuseImagined_EsoCalcada_960.jpg",
    "https://apod.nasa.gov/apod/image/2001/OrionTrees123019Westlake1024.jpg"
  ]
}

# Error
{
  "error": "'from' date should be prior to 'to' date, 'to' date should be after 'from' date"
}
```

# TODO
* Implement context logger
* Implement request tracing
* Implement a cache to avoid retrieving the same URLs multiple times
* Implement liveness and readiness probe
* Implement a hypothetical (PoC) mechanism to make calls to the mediadownloader service (gRPC?)
* Polish code
* Add tests