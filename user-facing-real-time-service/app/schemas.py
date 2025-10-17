from pydantic import BaseModel, Field
from typing import List

class LocationQueryRequest(BaseModel):
    query: str = Field(..., min_length=1, description="The location to search for, e.g., 'Eiffel Tower' or 'Berlin, Germany'")


class WeatherData(BaseModel):
    temperature: float
    condition: str

class PlaceData(BaseModel):
    name: str
    category: str

class EnrichedLocationResponse(BaseModel):
    #The final response model
    latitude: float
    longitude: float
    weather: WeatherData | None # None if the weather API fails
    places: List[PlaceData] = [] # Defaults to an empty list