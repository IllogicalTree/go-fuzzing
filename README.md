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
We were to and decided to fix it.

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

## Disclaimer

This is 100% satire, relying on a third party service to perform such a trivial task is absurd.

