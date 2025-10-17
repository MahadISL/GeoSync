from fastapi import FastAPI, HTTPException
from typing import Dict
from app import schemas
from app import clients

app = FastAPI(title="GeoSync API")


@app.get("/health")
def health_check() -> Dict[str, str]:
    # Health check endpoint to confirm the service is running.
    
    return {"status": "UP"}

@app.post("/api/location", response_model=schemas.CoordinatesResponse)
async def get_location_coordinates(request: schemas.LocationQueryRequest):
    #Accepts a location query and returns its geographical coordinates.
   
    coordinates = await clients.geocode_location(request.query)

    if coordinates is None:
        raise HTTPException(
            status_code=404,
            detail=f"Could not find coordinates for the query: '{request.query}'"
        )

    lat, lon = coordinates
    return schemas.CoordinatesResponse(latitude=lat, longitude=lon)