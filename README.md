# Flight Path Tracker Microservice

This microservice provides an API endpoint to track the flight path based on a list of flights provided as input. The API accepts a JSON payload containing the flights and total flight paths starting and ending airports.

## API Endpoint

`POST /calculate`

### Request Payload

The request payload should be a JSON array containing the list of flights in any order. Each flight should be represented as a JSON array with two elements: the source airport code and the destination airport code.

Example Request Payload:

```json
[
  ["IND", "EWR"],
  ["SFO", "ATL"],
  ["GSO", "IND"],
  ["ATL", "GSO"]
]
```

### Response

The API response will be a JSON array representing the sorted flight path. It contains the source and destination airports of the calculated flight path.

Example Response:

```json
["SFO", "EWR"]
```

If an error occurs during the calculation, the API will return an appropriate HTTP status code and an error message in the response body.

## Getting Started

### Prerequisites

- Go programming language (version 1.15 or higher) installed

### Running the Server

1. Clone the repository:

   ```bash
   git clone <repository-url>
   ```

2. Navigate to the server directory:

   ```bash
   cd flight-path-tracker-server
   ```

3. Start the server:

   ```bash
   go run main.go
   ```

   The server will start running on port 8080.

### Making API Requests

You can make API requests to the flight path tracker using a tool like `curl` or an API testing tool like Postman.

Example curl command:

```bash
curl -X POST -H "Content-Type: application/json" -d '[[ "IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]' http://localhost:8080/calculate
```

### Error Handling

The server code includes basic error handling for invalid input and internal errors. However, it's recommended to enhance the error handling as per your specific requirements and edge cases.

## License

This project is licensed under the [MIT License](LICENSE).
