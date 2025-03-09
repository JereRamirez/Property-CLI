# Property-CLI Tool

## Overview
This command-line tool allows users to filter, sort, and paginate a list of properties based on different criteria. The tool supports loading data from both local JSON files and remote URLs.

## Pre-Requisites
- Go 1.18+ (go run approach)
- Internet connection (if loading data from a URL)

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/JereRamirez/Property-CLI.git
   cd <repo-folder>
   ```
2. Build the application:
   ```sh
   go build -o property-filter main.go
   ```

## Usage

- The repository has a dummy properties.json file for basic usage; replace with desire data if necessary. 
- You can use this tool in two ways:

### 1. Running Directly with `go run`
This approach requires Go to be installed on your system.
```sh
 go run main.go <source> <filterKey> <filterType> <filterValue> [-sortBy key] [-desc true/false] [-page N] [-pageSize N]
```
Example:
```sh
 go run main.go properties.json price equal 250000
```

### 2. Building and Running the Executable
This method allows you to compile the tool into a standalone binary, which can be run without requiring Go.

#### Step 1: Build the tool
```sh
 go build -o property-filter
```

#### Step 2: Run the compiled binary
```sh
 ./property-filter <source> <filterKey> <filterType> <filterValue> [-sortBy key] [-desc true/false] [-page N] [-pageSize N]
```
Example:
```sh
 ./property-filter properties.json price equal 250000
```

### Parameters
- `<source>`: Path to a local JSON file or a URL containing property data.
- `<filterKey>`: The property attribute to filter by (e.g., `price`, `rooms`).
- `<filterType>`: Type of filter (`equal`, `lessThan`, `greaterThan`).
- `<filterValue>`: The value to compare against.
- `-sortBy key`: (Optional) Sort results by a specified key (`price`, `squareFootage`).
- `-desc true/false`: (Optional) Sort in descending order (`true`) or ascending (`false`, default).
- `-page N`: (Optional) Specify page number (default: `1`).
- `-pageSize N`: (Optional) Specify results per page (default: `10`).

### Examples (replace command with selected approach)
#### Filtering
#### Equal - LessThan - GreaterThan for Numeric fields (Price - Square Footage - Rooms - Bathrooms)
```sh
go run main.go properties.json price equal 250000 
go run main.go properties.json rooms equal 3

go run main.go properties.json price lessThan 200000
go run main.go properties.json squareFootage greaterThan 1400
```

#### Inclusion for Amenities
```sh
go run main.go properties.json amenities include garage
go run main.go properties.json amenities include pool
```
#### Matching a word in a Description
```sh
go run main.go properties.json description match luxury
```
#### Distance radius in km from a provided coordinate {latitude, longitude, kms}
```sh
go run main.go properties.json location distance 40.7128,-74.0060,5
go run main.go properties.json location distance 40.7306,-73.9352,10
```
#### Sorting
```sh
go run main.go properties.json price greaterThan 200000 -sortBy price
go run main.go properties.json rooms equal 4 -sortBy squareFootage -desc true
```
#### Pagination
```sh
go run main.go properties.json price greaterThan 100000 -page 2 -pageSize 20
go run main.go properties.json rooms greaterThan 2 -sortBy price -page 3 -pageSize 50
```
#### Combination
```sh
go run main.go properties.json squareFootage greaterThan 1500 -sortBy price -desc true -page 1 -pageSize 10
go run main.go properties.json price lessThan 500000 -sortBy squareFootage -page 2 -pageSize 15
```

#### Load from a URL, filter by rooms, sort by price, and paginate results
```sh
./property-filter "https://example.com/properties.json" rooms equal 3 -sortBy price -desc false -page 2 -pageSize 5
```

## Additional Features
- Supports large datasets with pagination
- Works with both local files and remote sources
- Extensible filtering and sorting capabilities

## Future Enhancements
- Support for additional data formats (e.g., CSV, XML)
- More advanced filtering (e.g., multiple conditions, range filters)
- Improved performance with indexing and caching