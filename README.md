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

## Additional Info

I added an additional route at `/receipts/{id}` to easily retrieve a given receipt. This was used for testing and I decided to leave it. In addition, I added the `/receipts` endpoint to return all receipts and a `/` endpoint to test that the server is running. 

I also created a basic `Item` model, which is not really used but I left it because it does not impact performance. 