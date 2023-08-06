# Api for interacting with the Open Broadcaster Software (OBS)'s' API
written in Go using Fiber
https://docs.gofiber.io

## Initialize project
```
go mod tidy
go run main.go
cp config.sample.toml config.toml

```
Update the config.toml file with your creds for OBS 
Test the API by hitting the endpoint with

```
http://127.0.0.1:3000/api/v1/admin/<endpoint mentioned in main.go>

```
