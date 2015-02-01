# Status

Demonstrate how to setup a `/.status` endpoint

Inspired by memcached "stats", this optional feature can be enabled to help monitoring the service.
This example shows how to enable the stats, and how to setup the `/.status` route.

curl demo:
```
curl -i http://127.0.0.1:8080/.status
curl -i http://127.0.0.1:8080/.status
...
```

Output example:
```
{
  "Pid": 21732,
  "UpTime": "1m15.926272s",
  "UpTimeSec": 75.926272,
  "Time": "2013-03-04 08:00:27.152986 +0000 UTC",
  "TimeUnix": 1362384027,
  "StatusCodeCount": {
        "200": 53,
        "404": 11
  },
  "TotalCount": 64,
  "TotalResponseTime": "16.777ms",
  "TotalResponseTimeSec": 0.016777,
  "AverageResponseTime": "262.14us",
  "AverageResponseTimeSec": 0.00026214
}
```
