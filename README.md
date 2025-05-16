# OAS Proxy

## Overview
OAS Proxy is HTTP proxy that sits in front of your application (typically backend) and 
verifies traffic against OpenAPI Specification (OAS) 3.0. Currently we can validate Request and Response payloads.

## Supported request formats
 - ✅ JSON Request validation
 - ❌ Form Request validation

## Modes of the proxy
 - ✅ Proxy mode - All requests even tho they are invalid, will pass through to the backend. When invalid request is detected, proper log is printed to STDOUT.
 - ❌ Strict mode - All requests that are invalid will be rejected with 400 Bad Request. Valid requests will be passed to the backend and proper message is included in the response body.

## Road map
 - ✅ JSON Request/Response validation
 - ✅ JSON Response validation
 - ❌ Strict mode
 - ❌ Proper testing coverage
 - ❌ Form Request/Response validation
 - ❌ Rework CLI / ENV variables handling
 - ❌ OpenTelemetry support
 - ❌ Proper CI/CD + well built Docker image
 - ❌ Proper versioning
 - ❌ Verify support for OAS 3.0->3.1



### Other docs
    - [Testing](docs/testing.md)
