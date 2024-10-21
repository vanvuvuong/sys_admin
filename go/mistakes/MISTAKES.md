Table of content <sup><sup>[back](../README.md)</sup></sup>


# Code & project organize
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

<details>
<summary> Unnecessary nested code </summary>

<details>