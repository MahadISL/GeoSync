from pydantic import BaseModel, Field


class LocationQueryRequest(BaseModel):
    query: str = Field(..., min_length=1, description="The location to search for, e.g., 'Eiffel Tower' or 'Berlin, Germany'")


class CoordinatesResponse(BaseModel):
    latitude: float
    longitude: float