# Weather Tracker App

A weather tracker application made with Go, Templ, HTMX and AlpineJS. The application uses the www.meteosource.com public API to generate the weather data for a given city. 

## Features 
- Search box
- 5 day weather forecast

## Getting Started

### Installation 


### Usage
```Shell
cp .env.example .env
# populate the API_KEY with your own

# using make 
make run

# using Go 
go run ./cmd/api/main.go
```


## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
