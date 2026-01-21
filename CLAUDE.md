# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

tm-go is a Go API client library for the Ticketmatic ticket management platform. It provides REST API bindings with HMAC-SHA256 authentication.

## Build and Test Commands

```bash
# Build all packages
go build ./...

# Run all tests (requires API credentials)
go test ./...

# Run tests for a specific package
go test ./ticketmatic/...
go test ./ticketmatic/orders/...

# Run a single test
go test ./ticketmatic -run TestFunctionName
```

**Test environment variables required:**
- `TM_TEST_ACCOUNTCODE` - Account code
- `TM_TEST_ACCESSKEY` - API access key
- `TM_TEST_SECRETKEY` - API secret key
- `TM_TEST_SERVER` (optional) - Defaults to https://apps.ticketmatic.com

## Architecture

### Core Package: `ticketmatic/`

- **client.go** - HTTP client with HMAC-SHA256 authentication, request building, error handling
- **types.go** - Generated data models (~280KB) for all Ticketmatic API entities
- **time.go** - Custom time type with flexible JSON parsing for multiple timestamp formats
- **stream.go** - Streaming support for newline-delimited JSON responses
- **widgets.go** - Widget URL generation and signature verification

### API Endpoint Packages

Each endpoint package follows the same pattern:
- `doc.go` - Package documentation
- `operations.go` - API functions (Getlist, Get, Create, Update, Delete)
- `operations_test.go` - Tests

**Main endpoints:** `contacts/`, `events/`, `orders/`, `diagnostics/`, `jobs/`, `eventstream/`, `subscribers/`, `tools/`, `sales/`

### Settings Packages (`settings/`)

Deeply nested configuration packages for:
- `communicationanddesign/` - Documents, email templates, ticket layouts, web skins
- `ticketsales/` - Delivery scenarios, lock types, fees, payment methods, sales channels
- `pricing/` - Price lists, price types, fees
- `seatingplans/` - Seating plans, seat ranks
- `system/` - Contact fields, custom fields, filters, reports, views (18+ sub-packages)
- `products/`, `productcategories/`, `vouchers/`

### Request Flow

1. Create `Client` with AccountCode, AccessKey, SecretKey
2. Call operation function (e.g., `orders.Get(client, id, params)`)
3. Request signed with HMAC-SHA256: `TM-HMAC-SHA256 key={accesskey} ts={timestamp} sign={hash}`
4. Response unmarshaled into typed structs

## Key Types

- `Client` - Main API client struct
- `Request` - HTTP request builder with fluent API
- `Time` - Custom time wrapper handling multiple timestamp formats
- `Stream` - For polling event streams
- `Widgets` - Widget URL generation and verification
- `RateLimitError`, `RequestError` - Error types

## Dependencies

Standard library only - no external dependencies.
