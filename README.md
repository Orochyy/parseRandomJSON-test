# ParseRandomJSON-test

## The Task
Create a production ready web service which combines two existing web services.
Fetch a random name from https://names.mcquay.me/api/v0
Fetch a random Chuck Norris joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy
Combine the results and return them to the user.
# Example
## Fetching a name

$ curl "https://names.mcquay.me/api/v0/"
{"first_name":"Hasina","last_name":"Tanweer"}
Fetching a joke
$ curl "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy"
{ "type": "success", "value": { "id": 181, "joke": "John Doe's OSI network model has only one layer - Physical.", "categories": ["nerdy"] } }

## Using the new web service

$ curl "http://localhost:5000"
Hasina Tanweer's OSI network model has only one layer - Physical..

## Run
```go
go run main.go
```