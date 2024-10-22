# Table of content <sup><sup>[back](../README.md)</sup></sup>

- [Table of content back](#table-of-content-back)
- [String formatting map](#string-formatting-map)
- [Golang Keywords](#golang-keywords)
- [Golang Built-ins](#golang-built-ins)
- [Code structure](#code-structure)
- [Golang mistakes](#golang-mistakes)

# [String formatting map](#table-of-content-back)

| Type name      | Abbr         | Type name                 | Abbr         |
| -------------- | ------------ | ------------------------- | ------------ |
| Type           | %T           | Unicode                   | %U           |
| Bool           | %t           | Float                     | %f / %e / %E |
| Int            | %d           | String                    | %s / %q      |
| Binary / Octal | %b / %o / %O | Value in a default format | %v / %#v     |
| Character      | %c           | Pointer                   | %p           |
| Error          | %w           | Number & Width            | %6d          |
| Heximal        | %x / %X      | Float & Width             | %6.3f        |
| String & Width | %10s / %-10s |                           |              |

# [Golang Keywords](#table-of-content-back)

| -        | -           | -      | -         | -      |
| -------- | ----------- | ------ | --------- | ------ |
| break    | default     | func   | interface | select |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

# [Golang Built-ins](#table-of-content-back)

| -       | -          | -         | -      | -     |
| ------- | ---------- | --------- | ------ | ----- |
| append  | bool       | byte      | cap    | close |
| complex | complex128 | complex64 | copy   | false |
| float32 | float64    | imag      | int    | int16 |
| int32   | int64      | int8      | iota   | len   |
| make    | new        | nil       | panic  | print |
| println | real       | recover   | string | true  |
| uint    | uint16     | uint32    | uint64 | uint8 |
| uintptr |            |           |        |       |

# [Code structure](#table-of-content-back)

<details>
<summary> 1. Common structure </summary>

```go
project-root/
    ├── cmd/
    │   ├── your-app-name/
    │   │   ├── main.go         # Application entry point
    │   │   └── ...             # Other application-specific files
    │   └── another-app/
    │       ├── main.go         # Another application entry point
    │       └── ...
    ├── internal/                # Private application and package code
    │   ├── config/
    │   │   ├── config.go       # Configuration logic
    │   │   └── ...
    │   ├── database/
    │   │   ├── database.go     # Database setup and access
    │   │   └── ...
    │   └── ...
    ├── pkg/                     # Public, reusable packages
    │   ├── mypackage/
    │   │   ├── mypackage.go    # Public package code
    │   │   └── ...
    │   └── ...
    ├── api/                     # API-related code (e.g., REST or gRPC)
    │   ├── handler/
    │   │   ├── handler.go      # HTTP request handlers
    │   │   └── ...
    │   ├── middleware/
    │   │   ├── middleware.go  # Middleware for HTTP requests
    │   │   └── ...
    │   └── ...
    ├── web/                     # Front-end web application assets
    │   ├── static/
    │   │   ├── css/
    │   │   ├── js/
    │   │   └── ...
    │   └── templates/
    │       ├── index.html
    │       └── ...
    ├── scripts/                 # Build, deployment, and maintenance scripts
    │   ├── build.sh
    │   ├── deploy.sh
    │   └── ...
    ├── configs/                 # Configuration files for different environments
    │   ├── development.yaml
    │   ├── production.yaml
    │   └── ...
    ├── tests/                   # Unit and integration tests
    │   ├── unit/
    │   │   ├── ...
    │   └── integration/
    │       ├── ...
    ├── docs/                    # Project documentation
    ├── .gitignore               # Gitignore file
    ├── go.mod                   # Go module file
    ├── go.sum                   # Go module dependencies file
    └── README.md                # Project README
```

Here's a brief explanation of the key directories:

- `cmd/`: This directory contains application-specific entry points (usually one per application or service). It's where you start your application.

- `internal/`: This directory holds private application and package code. Code in this directory is not meant to be used by other projects. It's a way to enforce access control within your project.

- `pkg/`: This directory contains public, reusable packages that can be used by other projects. Code in this directory is meant to be imported by external projects.

- `api/`: This directory typically holds HTTP or RPC API-related code, including request handlers and middleware.

- `web/`: If your project includes a front-end web application, this is where you'd put your assets (CSS, JavaScript, templates, etc.).

- `scripts/`: Contains scripts for building, deploying, or maintaining the project.

- `configs/`: Configuration files for different environments (e.g., development, production) reside here.

- `tests/`: Holds unit and integration tests for your code.

- `docs/`: Project documentation, such as design documents or API documentation.

The folder structure for a Go project can vary depending on the size and complexity of the project, as well as personal or team preferences. Here are some alternative folder structures for Go projects:

</details>

---

<details>
<summary> 2. Flat Structure </summary>
In smaller projects, you might opt for a flat structure where all your Go source files reside in the project root directory. This approach is simple but may become hard to manage as the project grows.

```go
project-root/
    ├── main.go
    ├── handler.go
    ├── config.go
    ├── database.go
    ├── ...
    ├── static/
    ├── templates/
    ├── scripts/
    ├── configs/
    ├── tests/
    └── docs/
```

</details>

---

<details>
<summary> 3. Layered Structure </summary>
Organize your code into layers, such as "web," "api," and "data." This approach helps separate concerns.

```go
project-root/
    ├── main.go
    ├── web/
    │   ├── handler.go
    │   ├── static/
    │   ├── templates/
    ├── api/
    │   ├── routes.go
    │   ├── middleware/
    ├── data/
    │   ├── database.go
    │   ├── repository.go
    ├── configs/
    ├── tests/
    ├── docs/
```

</details>

---

<details>
<summary> 4. Domain-Driven Design (DDD) </summary>
In larger applications, consider structuring your project based on domain-driven design principles. Each domain has its own directory.

```go
project-root/
    ├── cmd/
    │   ├── app1/
    │   ├── app2/
    ├── internal/
    │   ├── auth/
    │   │   ├── handler.go
    │   │   ├── service.go
    │   ├── orders/
    │   │   ├── handler.go
    │   │   ├── service.go
    │   ├── ...
    ├── pkg/
    │   ├── utility/
    │   │   ├── ...
    │   ├── ...
    ├── api/
    │   ├── app1/
    │   │   ├── ...
    │   ├── app2/
    │   │   ├── ...
    ├── web/
    │   ├── app1/
    │   │   ├── ...
    │   ├── app2/
    │   │   ├── ...
    ├── scripts/
    ├── configs/
    ├── tests/
    └── docs/
```

</details>

---

<details>
<summary> 5. Clean Architecture </summary>
You can adopt a clean architecture approach, which emphasizes a separation of concerns between different layers of your application.

```go
project-root/
   ├── cmd/
   │   ├── your-app/
   │   │   ├── main.go
   ├── internal/
   │   ├── app/
   │   │   ├── handler.go
   │   │   ├── service.go
   │   ├── domain/
   │   │   ├── model.go
   │   │   ├── repository.go
   ├── pkg/
   │   ├── utility/
   │   │   ├── ...
   ├── api/
   │   ├── ...
   ├── web/
   │   ├── ...
   ├── scripts/
   ├── configs/
   ├── tests/s
   └── docs/
```

</details>

---

<details>
<summary> 6. Modular Structure </summary>
Organize your code into separate modules, each with its own directory structure. This approach can be useful when developing multiple independent components within a single project.

```go
project-root/
    ├── module1/
    │   ├── cmd/
    │   ├── internal/
    │   ├── pkg/
    │   ├── api/
    │   ├── web/
    │   ├── scripts/
    │   ├── configs/
    │   ├── tests/
    │   └── docs/
    ├── module2/
    │   ├── ...
```

Remember that the right folder structure depends on the specific needs of your project and your team's development practices. Choose a structure that helps maintain code organization, readability, and collaboration as your project evolves.

</details>
  
---

# [Golang mistakes](#table-of-content-back)

Please refer to [here](./mistakes/MISTAKES.md)
