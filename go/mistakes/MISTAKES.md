# Table of content <sup><sup>[back](../README.md)</sup></sup>
- [Table of content back](#table-of-content-back)
- [Code \& project organize](#code--project-organize)
- [Data Type](#data-type)

# [Code & project organize](#table-of-content-back)
<details>
<summary> Unintended variable shadowing </summary>

- ~~Instead of~~
  ```go
  var client *http.Client
  if tracing {
      client, err := createClientWithTracing()
      if err != nil {
        return err
    }
    log.Println(client)
  } else {
    client, err := createDefaultClient()
    if err != nil {
    return err
    }
    log.Println(client)
  }
  // Use client
  ```

- Use this
  ```go
  var client *http.Client
  var err error
  if tracing {
      client, err = createClientWithTracing()
      if err != nil {
        return err
    }
  } else {
  // Same logic
  }
  // Use client
  ```

- Or this
  ```go
  if tracing {
    client, err = createClientWithTracing()
  } else {
    client, err = createDefaultClient()
  }
  if err != nil {
  // Common error handling
  }
  ```
</details>

---

<details>
<summary> Unnecessary nested code </summary>

- ~~Instead of this~~
  ```go
    // This join function concatenates two strings and returns a substring if the length is greater than max
    func join(s1, s2 string, max int) (string, error) {
      if s1 == "" {
        return "", errors.New("s1 is empty")
      } else {
        if s2 == "" {
          return "", errors.New("s2 is empty")
          } else {
            concat, err := concatenate(s1, s2)
            if err != nil {
              return "", err
            } else {
              if len(concat) > max {
                return concat[:max], nil
            } else {
                return concat, nil
            }
          }
        }
      }
    }
  ```
- [Use this](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88)
  ```go
  // Align the happy path to the left; you should quickly be able to scan down one column to see the expected execution flow.
  func join(s1, s2 string, max int) (string, error) {
    if s1 == "" {
      return "", errors.New("s1 is empty")
    }
    if s2 == "" {
      return "", errors.New("s2 is empty")
    }
    concat, err := concatenate(s1, s2)
    if err != nil {
      return "", err
    }
    if len(concat) > max {
      return concat[:max], nil
    }
    return concat, nil
  }
  ```
  ![fig1](./line-of-sight-in-code.png)
</details>

---

<details>
<summary> Create utility packages </summary>

> util is meaningless
- ~~Instead of this~~
  ```go
  package util
  func NewStringSet(...string) map[string]struct{} {
    // ...
  }
  func SortStringSet(map[string]struct{}) []string {
    // ...
  }
  set := util.NewStringSet("c", "a", "b")
  fmt.Println(util.SortStringSet(set))
  ```
- Use this
  ```go
  package stringset
  func New(...string) map[string]struct{} { ... }
  func Sort(map[string]struct{}) []string { ... }
  ```
</details>

---

# [Data Type](#table-of-content-back)

<details>
<summary> Slice length & capacity </summary>

  ```go
  s1 := make([]int, 3, 6) // 3-length, 6 capacity slice
  ```
![](./slice-length-capacity.png)
</details>

---

<details>
<summary> Inefficient slice initialization </summary>

- Instead of this
  ```go
  func convert(foos []Foo) []Bar {
    bars := make([]Bar, 0)
    for _, foo := range foos {
      bars = append(bars, fooToBar(foo))
    }
    return bars
  }
  ```
- Use this
  ```go
  func convert(foos []Foo) []Bar {
    n := len(foos)
    bars := make([]Bar, 0, n)
    for _, foo := range foos {
      bars = append(bars, fooToBar(foo))
    }
    return bars
  }
  ```
- Or this
  ```go
  func convert(foos []Foo) []Bar {
    n := len(foos)
    bars := make([]Bar, n)
    for i, foo := range foos {
      bars[i] = fooToBar(foo)
    }
    return bars
  }
  ```
</details>

