from pydantic_settings import BaseSettings, SettingsConfigDict

class Settings(BaseSettings):
   
    enrichment_service_url: str = "http://go-service:8081/enrich"

    model_config = SettingsConfigDict(env_file=".env", env_file_encoding='utf-8')

settings = Settings()