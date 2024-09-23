# Pokemon API Client 

This repository implements a client for pokemon APIs

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/codescalersinternships/PokeAPIHTTPClient-RawanMostafa.git
   ```

2. Install the dependencies
    ```bash
    go mod download
    ```

## Usage

  You can configure the endpoint for our client using different ways:

### 1. Using flags

   ```bash
    go run cmd/main.go [flags]
   ```
#### Flags:
   - endpoint=ENDPOINT

### 2. Using environment variables
This is used in case of no passed flags

   ```bash
    export VARNAME="my value"
    go run cmd/main.go 
   ```
#### Environment variables:
   - ENDPOINT


### 2. Using the default configurations
Our application provides default configurations in case no flags are provided and environment variables aren't set

   ```bash
    go run cmd/main.go 
   ```
#### Default configs:
  ```go
    const defaultEndpoint = "/machine"
  ```

Note: calling an API with known endpoint like Pokemon() will use its endpoint by default

#### Extra Utilities
  - Check this function to get environment variables : [`decideConfigs`](https://github.com/codescalersinternships/Datetime-client-RawanMostafa/blob/9c3cc7ecf671057648e10d07c550c171b51747a6/cmd/main.go#L43-L76)
  - Check this function to get the flags : [`getFlags`](https://github.com/codescalersinternships/Datetime-client-RawanMostafa/blob/9c3cc7ecf671057648e10d07c550c171b51747a6/cmd/main.go#L29-L41)
   

