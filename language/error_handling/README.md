## Error Handling Design

Error handling is critical for making your programs reliable, trustworthy and respectful to those who depend on them. A proper error value is both specific and informative. It must allow the caller to make an informed decision about the error that has occurred. There are several ways in Go to create error values. This depends on the amount of context that needs to be provided.

## Notes

* Use the default error value for static and simple formatted messages.
* Create and return error variables to help the caller identify specific errors.
* Create custom error types when the context of the error is more complex.
* Error Values in Go aren't special, they are just values like any other, and so you have the entire language at your disposal.

Solution:
```

```
