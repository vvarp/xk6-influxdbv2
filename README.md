# xk6-influxdbv2

This is a influxdb(v2) output library for [k6](https://github.com/loadimpact/k6),
implemented as an extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

The extension in this
current repo served as an example for an xk6 output tutorial,
but using one or the other is up to the user. :)

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  go get -u github.com/k6io/xk6/cmd/xk6
  ```

2. Build the binary:
  ```shell
  xk6 build v0.31.0 --with github.com/li-zhixin/xk6-influxdbv2
  ```

## Run

To run a `k6` case with this extension, first ensure you have  environment variables:

- influxDBv2Url
- influxDBv2Token
- influxDBv2Organization
- influxDBv2Bucket

Then run case:

```shell
k6 run ./test.js --out influxdbv2
```
