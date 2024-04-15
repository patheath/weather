# Weather
Practice Go by building a simple command line program for hourly weather

Hard coded location: 02481

# Weather providers:
1. weather.gov
2. weatherapi.com
- Register for an account and set the token into WA_TOKEN `export WA_TOKEN="<token>"`

# Basic flow
1. run via command line './bin/weather' after building `make build`
2. makes concurrent calls to weather providers
3. returns data and displays pretty version of output
4. exits

# Run tests
`make test` - run all

`make test -s` - run unit tests only

`make integ-test` - run integration tests
