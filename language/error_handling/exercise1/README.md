## Exercise 1

Create two error variables, one called ErrInvalidValue and the other called ErrAmountTooLarge. Provide the static message for each variable. Then write a function called checkAmount that accepts a float64 type value and returns an error value. Check the value for zero and if it is, return the ErrInvalidValue. Check the value for greater than $1,000 and if it is, return the ErrAmountTooLarge. Write a main function to call the checkAmount function and check the return error value. Display a proper message to the screen.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/Ltxl8Hkrkl)) | [Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/WHmYkHwYjf))
