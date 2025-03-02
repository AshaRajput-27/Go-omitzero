# 🚀 'omitzero' - A New Way to Handle Zero Values in JSON (Go 1.24)

## WHAT:
#### This feature helps automatically remove fields with zero values when converting a struct to JSON, making the output cleaner and more efficient. </br>
#### Earlier, we had omitempty, but omitzero is more flexible and works better for different data types. </br>

<hr>

## WHY:
#### When converting a Go struct to JSON, fields with zero values (like 0, "", false, nil, etc.) are still included unless we explicitly use omitempty. </br>
#### With omitzero, Go automatically detects zero values for all types and removes them from the JSON output. </br>
#### This reduces the size of the JSON and makes the API responses cleaner. </br>

<hr>

### 🔴 Before (Without omitzero)
```go
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Active bool  `json:"active"`
}
```
[View Code File 'withoutOmitzero.go'](withoutOmitzero.go)

```json
{"name":"omit zero","age":0,"email":"","active":false}
```

### 🟢 Using 'omitzero' (Go 1.24)
```go
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age,omitzero"`
	Email string `json:"email,omitzero"`
	Active bool  `json:"active,omitzero"`
}
```
[View Code File 'withOmitzero.go'](withOmitzero.go)

```json
{"name":"omit zero"}
```

### 🚀 Boom! 
#### The zero-value fields (age, email, active) are automatically removed, making the JSON response cleaner.

<hr>

