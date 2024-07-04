# KeyValueStore

KeyValueStore is a side project aimed at enhancing Go programming skills and gaining deeper knowledge of persistence
structures such as key/value stores and append-only logs combined with in-memory hash indexes.

## Introduction

KeyValueStore is a simple implementation of a key/value store written in Go. It utilizes an append-only log for
persistent storage and in-memory hash indexes for fast lookups. This project serves as a practical exercise to improve
Go programming skills and understanding of basic data persistence mechanisms.

## Features

- **Key/Value Store**: Store and retrieve values based on keys.
- **Append-Only Log**: Persistent storage using an append-only log for durability.
- **In-Memory Indexes**: Fast lookups using in-memory hash indexes.
- **Graceful Shutdown**: Proper shutdown handling to ensure data consistency.
- **Logging**: Structured logging for better debugging and monitoring.

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Git

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/keyvaluestore.git
    cd keyvaluestore
    ```

2. Initialize and tidy up Go modules:
    ```sh
    go mod tidy
    ```

3. Build the project:
    ```sh
    go build -o keyvaluestore
    ```

## Usage

1. Run the key/value store server:
    ```sh
    ./keyvaluestore
    ```

2. Interact with the store using HTTP requests. Examples:
    - **Set a value**:
        ```sh
        curl -X POST -d '{"key":"foo", "value":"bar"}' http://localhost:8080/key-value
        ```

    - **Get a value**:
        ```sh
        curl -X GET 'http://localhost:8080/key-value/foo'
        ```
