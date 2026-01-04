# Weather Service

A simple Go application that provides weather information for specific locations.

## How to Build

To build the application, run the following command in the project root:

```bash
go build -o weather-service cmd/weather-service/main.go
```

## How to Run

To run the application, execute the built binary:

```bash
./weather-service
```

Or run it directly using `go run`:

```bash
go run cmd/weather-service/main.go
```

The server will start on port 8080.

## API Endpoints

### 1. Get Weather

Returns the weather for a specific location.

**Endpoint:** `/getWeather`
**Method:** `GET`
**Query Parameter:** `location` (required)

**Example:**

```bash
curl "http://localhost:8080/getWeather?location=Frauenfeld"
```

**Response:**

```json
{
  "location": "Frauenfeld",
  "condition": "cloudy",
  "temperature": "15°"
}
```

### 2. Get Supported Locations

Returns a list of all supported locations.

**Endpoint:** `/getSupportedLocations`
**Method:** `GET`

**Example:**

```bash
curl "http://localhost:8080/getSupportedLocations"
```

**Response:**

```json
[
  "Frauenfeld",
  "Miami",
  "Bangkok"
]
```

### 3. Get OpenAPI Specification

Downloads the OpenAPI specification file.

**Endpoint:** `/openapi.yaml`
**Method:** `GET`

**Example:**

```bash
curl "http://localhost:8080/openapi.yaml"
```

### 4. Get Metrics

Returns Go runtime metrics and custom application metrics in JSON format.

**Endpoint:** `/metrics`
**Method:** `GET`

**Example:**

```bash
curl "http://localhost:8080/metrics"
```

**Response (truncated):**

```json
{
  "weather_requests_total": 12,
  "supported_locations_requests_total": 5,
  "/gc/cycles/automatic:gc-cycles": 0,
  "/gc/heap/allocs:bytes": 1048576,
  "/sched/goroutines:goroutines": 5
}
```

### 5. Health Check

Returns the health status of the application.

**Endpoint:** `/health`
**Method:** `GET`

**Example:**

```bash
curl "http://localhost:8080/health"
```

**Response:**

```json
{
  "status": "up"
}
```

### 6. Profiling (pprof)

The application exposes pprof endpoints for profiling.

**Endpoint:** `/debug/pprof/`
**Method:** `GET`

**Example:**

View the pprof index in a browser:
`http://localhost:8080/debug/pprof/`

Download a heap profile:
```bash
curl -o heap.out http://localhost:8080/debug/pprof/heap
```

Analyze the profile with the Go tool:
```bash
go tool pprof heap.out
```

## Docker

### Build the Image

```bash
docker build -t weather-service .
```

### Run the Container

```bash
docker run -p 8080:8080 weather-service
```

### 6. Profiling (pprof)

The application exposes pprof endpoints for profiling.

**Endpoint:** `/debug/pprof/`
**Method:** `GET`

**Example:**

View the pprof index in a browser:
`http://localhost:8080/debug/pprof/`

Download a heap profile:
```bash
curl -o heap.out http://localhost:8080/debug/pprof/heap
```

Analyze the profile with the Go tool:
```bash
go tool pprof heap.out
```
