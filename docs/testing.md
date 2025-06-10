# Testing & Local Development
> This document describes how to run the tests for the project.

## Run the test app
> This section describes how to run the test app, which is proxied by the OAS Proxy.
> All tests are running in strict mode, which means that they should end up with 400 Bad Request if the request is invalid, this is used to verify that the proxy is working correctly, since logs cant be checked within test framework.

### Dependencies
To run the test app, you need to have the following dependencies installed:
- Python 3.13 (tested on it) or higher
- Installed dependencies from `tests/python/requirements.txt` file

### Prepare the environment
0. cd into the `tests/python` directory:
```bash
cd tests/python
```
1. Create a virtual environment:
```bash
python3 -m venv venv
```
2. Activate the virtual environment:
```bash
source venv/bin/activate
```
3. Install the dependencies:
```bash
pip install -r requirements.txt
```
4. Run the test app:
```bash
python3 app.py
```
or, but make sure that the virtual environment is activated:
```bash
python3 tests/python/app.py
```

## Start OAS Proxy
> This section describes how to start the OAS Proxy, which is used to test the test app.

### Run OAS Proxy
```bash
TARGET_URL=http://localhost:8000 \
STRICT_MODE=true \
go run cmd/app/main.go
```


## Verify the Proxy and Test App
> Test the apps with curl requests...

### Verify Test App
```bash
curl -d '{ 
    "integer": 2,
    "string": "string"
    }' -X POST --header "Content-Type: application/json" --silent localhost:8000/1/foo | jq .
```

### Verify OAS Proxy
```bash
curl -d '{ 
    "integer": 2,
    "string": "string"
    }' -X POST --header "Content-Type: application/json" --silent localhost:8080/1/foo | jq .
```