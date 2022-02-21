# String reversal as a service

We had the revolutionary idea to create the next big SaaS product, introducing "String Reversal" as a service. We completely eliminate the 4 lines of javascript required to reverse a string in the place of a single fetch call, 100% guaranteed to load faster than all of your ads and analytics.

## Why do I want this?

Clean and performant code of course, our purpose built and optimised algorithms reverse your string in record time. Reduce the number of lines required to reverse a string by 75%;

```
var str = "Hello World!";
var splitString = str.split("");
var reverseArray = splitString.reverse();
var joinArray = reverseArray.join("");
```

Are you fed up with writing all of this just to reverse a string in javascript?
We were too and decided to fix it.

## Getting started

Get the server up and running using

```
go run .
```

Send a post request with your string that you need reversed:

```
curl --location --request POST 'localhost:8090' \
--header 'Content-Type: application/json' \
--data-raw '{
    "string": "Hello world!"
}'
```

Be amazed at the speedy and useful response
```
{
    "string": "!dlrow olleH"
}
```

## Testing

To ensure we deliver on our promise of reversing strings accurately our developers grudgingly added a very comprehensive suite of tests.

These tests can be run using the following command:

```
go test
```

### Further testing

After some ex-customers reported that some strings weren't being reversed correctly, management forced us to investigate and said our 3 tests didn't provide enough coverage.. our senior developer was about to manually add more tests but a new hire suggested using fuzz testing. 

We uncovered some issues with our string reversal algorithm but have since rectified these, these fuzz tests can again be run using:

```
go test -fuzz=Fuzz
```

## Disclaimer

This is 100% satire, relying on a third party service to perform such a trivial task is absurd.

