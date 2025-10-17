import httpx


NOMINATIM_BASE_URL = "https://nominatim.openstreetmap.org/search"
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