# Worldline Acquiring Go SDK

## Introduction

The Go SDK helps you to communicate with the Worldline Acquiring API. Its primary features are:

* convenient go wrapper around the API calls and responses
    * marshalls Go request structs to HTTP requests
    * unmarshalls HTTP responses to Go response structs or Go exceptions
* handling of all the details concerning authentication
* handling of required metadata

See the [Worldline Acquiring Documentation](https://docs.acquiring.worldline-solutions.com/Developer-Tools/sdk/go) for more information on how to use the SDK.

## Requirements

Go version 1.26 or higher is supported. No additional packages are required.

## Examples

This repository also contains some example code. This code is contained in the examples folder.

## Installation

### Source

To install the latest version of this repository, run the following command from a terminal:

    go get github.com/Worldline-Acquiring/acquiring-sdk-go

### Release

You can add `@version` to the go get command to get a specific version. For example, `go get github.com/Worldline-Acquiring/acquiring-sdk-go@2.9.0` will download version 2.9.0 of the SDK. See the releases page for an overview of available releases.

Note that new major versions may introduce breaking changes. We therefore recommend pinning the version of the SDK in your [Go module](https://go.dev/ref/mod).

## Running tests

There are two types of tests: unit tests and integration tests. The unit tests will work out-of-the-box; for the integration tests some configuration is required. First, some environment variables need to be set:

* `acquiring.api.oauth2.clientId` for the OAUth2 client id to use.
* `acquiring.api.clientSecret` for the OAuth2 client secret to use.
* `acquiring.api.merchantId` for your merchant ID.
* `acquiring.api.proxyUrl` for the URL to the proxy to use (optional). If set, it should be in the form `scheme://[userinfo@]host[:port]`. Examples: `http://proxy.example.org`, `http://user:pass@proxy.example.org`, `http://proxy.example.org:3128`.

The following commands can now be executed from the root directory of the SDK folder to execute the tests:

* Unit tests:
    
    ```
    go test ./...
    ```
*  Both unit and integration tests:
    
    ```
    go test -tags=integration  ./...
    ```
