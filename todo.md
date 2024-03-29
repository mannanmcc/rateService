- [&check;] Add container for wiremock to mock remote api call (rest GET api)
    Follow this page: https://www.alexhyett.com/mock-api-calls-wiremock/
- [&check;] Custom error type to wrap the dependency downstream error
- [&check;] add OZZO validation to the request in the service layer to validate the request
- [&cross;] Clean up a bit on the bdd test suite setup such as move few things to the test_setuo.go file
- [&cross;] introduce mongo or caching to store the rate for 24 hours
- [&cross;] Add timeout for the remote api call
- [&cross;] Add another GRPC service as dependencies and setup proxy for that service mocking for BDD test
- [&check;] add zap logger with context to log - 
- [&check;] Add transport layer and call handler from TL
- [&cross;] Add linter configuration
- [&cross;] Add interface to log the request and response of the api
- [&cross;] Check how to use dependancy injection provider instead of putting everything manually in main.go file
- [&cross;] Read about secure code in golang, https://securego.io/docs/rules/g102.html


Reading:
- https://slscan.io/en/latest/secure-development/go/