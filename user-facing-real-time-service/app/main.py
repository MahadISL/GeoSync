from fastapi import FastAPI
from typing import Dict

app = FastAPI(title="GeoSync API")


@app.get("/health")
def health_check() -> Dict[str, str]:
    # Health check endpoint to confirm the service is running.
    
    return {"status": "UP"}