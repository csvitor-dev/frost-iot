# `frost-iot` | IOT Pseudo-Application with _Golang_

## Development Flow

Whenever you are going to perform a task associated with an **Issue**, create a new _branch_:

```bash
git checkout -b your-branch
```

When finished, create a **Pull Request** for the `main` _branch_!

## Project Structure

```text
├───cmd
│   ├───client
│   ├───device
│   └───server
├───examples
├───internal
│   ├───messages
│   │   ├───constants
│   │   ├───requests
│   │   └───responses
│   ├───owtp
│   ├───socket
│   │   ├───process
│   │   └───transporters
│   └───types
├───pkg
│   └───types
└───src
    ├───actuators
    ├───client
    ├───device
    ├───manager
    ├───sensors
    └───view
```
