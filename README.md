# fetch-receipt-processor-challenge
https://github.com/fetch-rewards/receipt-processor-challenge

## Description

I created this repo to meet the coding challenge for Fetch. 

## Startup

`go run main.go`

## Interacting with the API

Submit requests to `localhost:8080`

## Testing

I have written some very basic tests in `test.sh`. With the server running, run `./test.sh` and the points output should match the hardcoded "expected" values.

In addition, I am using the `testing` library to write more extensive tests. To run tests: `go test ./src/models -v`

## Additional Info

I added an additional route at `/receipts/{id}` to easily retrieve a given receipt. This was used for testing and I decided to leave it. In addition, I added the `/receipts` endpoint to return all receipts and a `/` endpoint to test that the server is running. 

I used this exercise to learn basic Go coding techniques. I attempted to apply best practices but this is my first time exploring the language.