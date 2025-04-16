# Parallel Data Processing

This project is a Go-based data processing application that efficiently handles and analyzes CSV data using parallel processing techniques.

## Project Structure

The project follows a clean architecture pattern with the following layers:

```ini
[Infrastructure] → [Repository] → [Use Case] → [Presenter] → [Output]

```

### Key Components

- **Infrastructure Layer**: Handles low-level operations like CSV file reading
- **Repository Layer**: Provides data access interfaces
- **Use Case Layer**: Contains business logic for data processing
- **Presenter Layer**: Handles data presentation and output formatting

## Features

- Parallel CSV file reading and processing
- Efficient data grouping and sampling
- Time-based data aggregation
- Master data integration
- Concurrent processing with worker pools

## Data Processing Flow

1. **Master Data Processing**:

   - Reads and groups master data from CSV
   - Provides reference data for raw data processing

2. **Raw Data Processing**:

   - Processes large CSV datasets in parallel
   - Implements time-based sampling
   - Calculates hourly averages
   - Uses worker pools for efficient processing

## Getting Started

### Prerequisites

- Go 1.21 or later
- CSV data files in the correct format

### Running the Application

1. Place your CSV files in the `assets` directory:

   - `GTG-1_MASTER_DATA.csv` for master data
   - `GTG-1.csv` for raw data

2. Run the application:

```bash
go run cmd/main.go

```

## Project Structure

```ini
.
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── infrastructure/      # Low-level implementations
│   │   └── csv_reader.go    # CSV file reading
│   ├── interfaces/          # Repository interfaces
│   │   ├── csv_presenter.go # Data presentation
│   │   └── csv_repository.go# Data access interfaces
│   ├── usecase/            # Business logic
│   │   └── process_raw.go   # Raw data processing
│   └── entities/           # Data models
│       └── master.go       # Master data structures
└── assets/                 # Data files

```

## Data Processing Features

- **Parallel Processing**: Uses goroutines and channels for efficient data processing
- **Time-based Aggregation**: Groups data by hour and calculates averages
- **Data Validation**: Includes error handling and data type conversion
- **Memory Efficient**: Processes data in chunks to handle large datasets

## Contributing

Feel free to submit issues and enhancement requests.