from fastapi import FastAPI, HTTPException
from typing import Dict
from app import schemas
from app import clients

app = FastAPI(title="GeoSync API")


@app.get("/health")
def health_check() -> Dict[str, str]:
    # Health check endpoint to confirm the service is running.
    
    return {"status": "UP"}

@app.post("/api/location", response_model=schemas.EnrichedLocationResponse)
async def get_location_coordinates(request: schemas.LocationQueryRequest):
    """
    Full process:
    1. Geocodes the user's query to get coordinates.
    2. Calls the enrichment service to get weather and places.
    3. Returns the aggregated data.
    """
   
    coordinates = await clients.geocode_location(request.query)
    if coordinates is None:
        raise HTTPException(
            status_code=404,
            detail=f"Could not find coordinates for the query: '{request.query}'"
        )
    lat, lon = coordinates

    # Step 2: Call the Go service for enrichment
    enriched_data = await clients.get_enriched_data(lat, lon)

    # Step 3: Combine and return the final response
    response_data = {
        "latitude": lat,
        "longitude": lon,
        "weather": enriched_data.get("weather") if enriched_data else None,
        "places": enriched_data.get("places") if enriched_data else []
    }

    return schemas.EnrichedLocationResponse(**response_data)