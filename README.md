# GeoSync

![Project Banner](https://placehold.co/1200x300/333/FFF?text=GeoSync%20)

**GeoSync is a high-performance, polyglot backend system built on a microservice architecture designed to provide enriched geographical data. It takes a location query, resolves it to coordinates, and returns real-time weather data and local place information.**

This project serves as a practical demonstration of modern backend development practices, including microservice architecture, polyglot programming (Go & Python), containerization with Docker, and production-ready configuration.

---

## Features

*   **Location to Coordinates:** Converts any textual location query (e.g., "Eiffel Tower") into precise latitude and longitude.
*   **Data Enrichment:** Fetches real-time weather conditions for the given coordinates.
*   **Concurrent Processing:** Utilizes Go's powerful concurrency features (goroutines) to fetch data from multiple external APIs in parallel for maximum speed.
*   **Fully Containerized:** The entire application stack is defined in Docker and can be launched with a single command.
*   **Clean Architecture:** Services are decoupled, independently deployable, and communicate over a defined network.
*   **Interactive API Docs:** The FastAPI service provides automatic, interactive API documentation via Swagger UI.

---

## Architecture

The system is composed of two distinct microservices that run in separate Docker containers and communicate over a private network managed by Docker Compose.

```mermaid
graph TD
    subgraph "User's Machine"
        User(["👨‍💻 User / Client"])
    end

    subgraph "Docker Environment (Managed by Docker Compose)"
        F[🐍 FastAPI Service <br>(Port 8000)]
        G[🐹 Go Enrichment Service <br>(Port 8081)]
    end

    subgraph "External APIs"
        Geo[🗺️ Nominatim Geocoding API]
        Weather[☀️ OpenWeatherMap API]
        ReverseGeo[📍 Nominatim Reverse Geocoding]
    end

    User -- "HTTP POST /api/location" --> F
    F -- "Internal HTTP Call" --> G
    F -- "API Call" --> Geo
    G -- "Concurrent API Calls" --> Weather
    G -- "Concurrent API Calls" --> ReverseGeo

    style F fill:#009688,stroke:#333,stroke-width:2px,color:#fff
    style G fill:#007D9C,stroke:#333,stroke-width:2px,color:#fff

### 1. FastAPI Service (`fastapi-service`)
*   **Language:** Python (with FastAPI)
*   **Role:** The **Orchestrator** and public-facing entry point.
*   **Responsibilities:**
    *   Exposes the public REST API (`/api/location`).
    *   Validates user input using Pydantic.
    *   Calls an external geocoding service (Nominatim) to get coordinates.
    *   Calls the internal Go service to get enriched data.
    *   Aggregates the responses into a final JSON object for the user.

### 2. Go Enrichment Service (`go-service`)
*   **Language:** Go (with Gin)
*   **Role:** The high-performance, concurrent **Workhorse**.
*   **Responsibilities:**
    *   Exposes a single internal REST API (`/enrich`).
    *   Accepts latitude and longitude.
    *   Launches concurrent API calls (goroutines) to:
        *   **OpenWeatherMap** for current weather data.
        *   **Nominatim** for reverse geocoding to find the address.
    *   Aggregates the results and returns them to the FastAPI service.

---

## Tech Stack

*   **Backend:** Python 3.11 (FastAPI), Go 1.22 (Gin)
*   **Containerization:** Docker, Docker Compose
*   **Testing:** Pytest (Python), Go Test with Testify (Go)
*   **Linting/Formatting:** (Future enhancement: Black, Ruff, gofmt)

---

## Getting Started

### Prerequisites

*   [Docker](https://www.docker.com/products/docker-desktop/) installed and running.
*   [Docker Compose](https://docs.docker.com/compose/install/) (included with Docker Desktop).
*   A free API key from [OpenWeatherMap](https://openweathermap.org/api).
*   Git.

### Installation & Launch

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/MahadISL/GeoSync.git
    cd GeoSync
    ```

2.  **Create your environment file:**
    Create a `.env` file in the root of the project by copying the example.
    ```
    Now, open the `.env` file and add your OpenWeatherMap API key:
    ```env
    OPENWEATHERMAP_API_KEY=your_actual_api_key_here
    ```

3.  **Build and run with Docker Compose:**
    This single command builds the images for both services and starts the containers in the correct order.
    ```bash
    docker-compose up --build
    ```
    The API will be available at `http://localhost:8000`.

---

## Usage

Once the application is running, you can interact with it through the API.

### Get Enriched Location Data

Send a `POST` request to `/api/location`.

**Example using `curl`:**```bash
curl -X POST http://localhost:8000/api/location \
-H "Content-Type: application/json" \
-d '{"query": "CN Tower, Toronto"}'
```

**Expected Success Response:**
```json
{
  "latitude": 43.6425662,
  "longitude": -79.3870568,
  "weather": {
    "temperature": 15.1,
    "condition": "Clouds"
  },
  "places": [
    {
      "name": "CN Tower, 301, Front Street West, Financial District, Downtown, Toronto, Golden Horseshoe, Ontario, M5V 2T6, Canada",
      "category": "Location Address"
    }
  ]
}
```

### Interactive Documentation

For a full, interactive API explorer, navigate to the auto-generated Swagger UI in your browser:
**[http://localhost:8000/docs](http://localhost:8000/docs)**

### Running Tests

To run the automated tests for each service:

*   **Go Service:**
    ```bash
    cd geo-enrichment-service
    go test ./...
    ```

*   **FastAPI Service:**
    ```bash
    cd user-facing-real-time-service
    # Make sure to activate the virtual environment first
    source .\.venv\Scripts\Activate.ps1
    pytest
    ```

---

## Future Enhancements

*   Implement a real-time WebSocket endpoint for weather updates.
*   Add a database to cache results and reduce external API calls.
*   Deploy the application to a cloud provider (e.g., using AWS ECS or Google Cloud Run).
*   Integrate a CI/CD pipeline with GitHub Actions.