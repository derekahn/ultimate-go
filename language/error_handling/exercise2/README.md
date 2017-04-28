## Exercise 2
Create a custom error type called appError that contains three fields, err error, message string and code int. Implement the error interface providing your own message using these three fields. Implement a second method named temporary that returns false when the value of the code field is 9. Write a function called checkFlag that accepts a bool value. If the value is false, return a pointer of your custom error type initialized as you like. If the value is true, return a default error. Write a main function to call the checkFlag function and check the error using the temporary interface.

[Template](exercises/template2/template2.go) ([Go Playground](http://play.golang.org/p/9nEdNSMa_j)) |
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](http://play.golang.org/p/7iX9wZX6WP))

Solution:
```

```
