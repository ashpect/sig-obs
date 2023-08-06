# Api for interacting with the Open Broadcaster Software (OBS)'s' API
written in Go using Fiber
https://docs.gofiber.io

## Initialize project
```
cp config.sample.toml config.toml
go mod tidy

```
Update the config.toml file with your creds for OBS 
Run the project with

```
go run main.go

```
Test the API by hitting the endpoint with

```
http://127.0.0.1:3000/api/v1/admin/<endpoint mentioned in main.go>

```
