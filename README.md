# Coupon Issuance System

## Features

- Campaign creation
- Coupon issuance
- Distributed environment support (Redis-based)
- Concurrency control (Redis distributed lock)

## Installation

1. Install and start Redis:

```bash
# Install Redis (macOS)
brew install redis

# Start Redis server
brew services start redis
```

2. Install dependencies:

```bash
go mod download
```

## Running Server

```bash
go run cmd/server/main.go
```

## Running Tests

```bash
# Run tests
go test ./...
```

## Project Structure

```
.
├── cmd/
│   └── server/          # Server entry point
│       └── main.go
├── internal/
│   ├── campaign/        # Campaign service
│   │   ├── service.go   # Campaign service implementation
│   │   └── service_test.go
│   └── coupon/          # Coupon generator
│       └── generator.go
├── gen/                 # Protocol Buffer generated files
│   └── campaign/
│       └── v1/
├── go.mod
└── README.md
```

## Data Structure

#### Campaign

- `campaign`
  - `id`: Campaign ID
  - `available_coupons`: Number of available coupons
  - `start_time`: Campaign start time (Unix timestamp)

#### Coupon

- `coupon`
  - `id`: Coupon ID
  - `campaign_id`: Associated campaign ID
  - `coupon_code`: Generated coupon code

## API Endpoints

### CreateCampaign

- Create new campaign

### GetCampaign

- Get campaign information
- Get list of issued coupon codes

### IssueCoupon

- Issues coupon
- Uses distributed lock for concurrency control

## Concurrency Control

- Distributed lock for concurrency control
- Atomic operations for coupon count management