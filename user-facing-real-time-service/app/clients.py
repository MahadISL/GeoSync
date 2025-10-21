import httpx
from app import config import settings

NOMINATIM_BASE_URL = "https://nominatim.openstreetmap.org/search"
# ENRICHMENT_SERVICE_URL = "http://geo-enrichment-service:8081/enrich"
USER_AGENT = "GeoSync-API-App"

async def geocode_location(query: str) -> tuple[float, float] | None:
    """
    Calls the Nominatim API to convert a location query into coordinates.
    Returns a tuple of (latitude, longitude) or None if not found.
    """
    async with httpx.AsyncClient() as client:
        params = {"q": query, "format": "json", "limit": 1}
        headers = {"User-Agent": USER_AGENT}

        try:
            response = await client.get(NOMINATIM_BASE_URL, params=params, headers=headers)
            response.raise_for_status()  # Raise an exception for 4XX/5XX responses

            data = response.json()
            if not data:
                return None

            latitude = float(data[0]["lat"])
            longitude = float(data[0]["lon"])
            return latitude, longitude

        except (httpx.RequestError, KeyError, IndexError, ValueError) as e:
            print(f"An error occurred during geocoding: {e}")
            return None


async def get_enriched_data(lat: float, lon: float) -> dict | None:
    #Calls the Go geo-enrichment-service to get weather and places data.
    async with httpx.AsyncClient() as client:
        try:
            # The Go service expects a POST request with a JSON body
            response = await client.post(settings.enrichment_service_url, json={"latitude": lat, "longitude": lon})
            response.raise_for_status()
            return response.json()
        except httpx.RequestError as e:
            print(f"An error occurred while calling the enrichment service: {e}")
            return None