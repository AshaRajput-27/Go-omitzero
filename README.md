## 🚀 OmitZero in Go 1.24
#### This repository contains documentation and hands-on examples demonstrating the omitzero struct tag introduced in Go 1.24. The omitzero tag allows fields with zero values to be omitted when marshaling a struct into JSON.

<hr> 

### 📌 What is omitzero?
#### The omitzero tag is used with encoding/json to automatically omit fields that have zero values when converting a Go struct to JSON.

<hr>

### ✅ Zero Values in Go:
```
0 for numbers (int, float, etc.)
false for bool
"" (empty string) for string
nil for pointers, slices, maps, and interfaces
```
