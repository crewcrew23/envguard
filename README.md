# envguard

`envguard` — This is a simple, readable and flexible library for inspecting `.env` files in Go.

📌 Supports validation `int`, `string`, `URL`, `IP`  
✅ Strict rules: `Min`, `Max`, `Regex`, `UUID`, `Email`, `IP v4/v6` etc.  
💡 Convenient API: readable call chain

---

## 🔧 Installation

```bash
go get github.com/crewcrew23/envguard
```

## 🚀 Quick start

```go
package main

import (
    "fmt"
    "log"

    "github.com/crewcrew23/envguard"
)

func main() {
    file, err := envguard.ParseFile(".env")
    if err != nil {
        log.Fatal(err)
    }

    err = file.Get("UUID").String().UUID().Validate()
    if err != nil {
        fmt.Println("Error UUID:", err)
    }

    err = file.Get("PORT").Integer().Min(1024).Max(65535).Validate()
    if err != nil {
        fmt.Println("Error PORT:", err)
    }

    err := file.Get("IP").IP().V4().Validate()
    if err != nil {
        fmt.Println("Error IP:", err)
    }

    //Custom validate
    err := file.Get("CUSTOM_VALUE").Integer().Custom(func(i int) bool {
		//other logic
	}, "err msg")

}
```

## ✅ Supported Validators

---

### 🔢 Integer

- `Min(int)`
- `Max(int)`
- `Between(min, max int)`
- `Even()`
- `Odd()`
- `Positive()`, `Negative()`, `NonZero()`
- `Contains(...int)`
- `NotContains(...int)`
- `DivisibleBy(int)`
- `Custom(func(int) bool, errMessage string)`

---

### 🔤 String

- `Min(int)`
- `Max(int)`
- `Length(min, max int)`
- `NotEmpty()`
- `NotBlank()`
- `IsAlpha()`
- `IsAlphanumeric()`
- `HasPrefix(string)`
- `HasSuffix(string)`
- `Contains(string)`
- `NotContains(string)`
- `Email()`
- `UUID()`
- `MatchRegex(pattern string)`
- `Custom(func(string) bool, errMessage string)`

---

### 🌐 URL

- `Scheme(...string)`
- `Host(...string)`
- `Port(...string)`
- `Custom(func(string) bool, errMessage string)`

---

### 🌍 IP

- `V4()`
- `V6()`
- `Custom(func(string) bool, errMessage string)`
