# GoGet / Validator

This library provides an easy way to validate form inputs in go. All inputs are validated as strings so reflection is not needed. This library takes the approach of validation prior to adding values to a struct as opposed to after.

### WORK IN PROGRESS

This is an early stage project. The test suite is not yet complete. Please consider this a pre-alpha release.

### How It Works

```
// setup your rules
rules := make(map[string]string)
rules["name"] = "required|alpha"
rules["email"] = "required|email"

// parse the form
r.ParseForm()

// call the validator
success, messages := validator.Validate(r.Form, rules)
```

### Available Rules

Docs coming soon.

### Credits and Inspiration

The validation rules in this library were inspired by the Laravel Framework for PHP.

### Warranty and License

No warranty provided. Use at your own risk. :)

Licensed under the MIT license.