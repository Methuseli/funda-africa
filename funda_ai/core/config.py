# core/config.py
from pydantic_settings import BaseSettings
from pydantic import AnyHttpUrl


class Settings(BaseSettings):
    # External APIs
    WOLFRAM_APPID: str 
    GEMINI_API_KEY: str  # e.g. OpenAI
    # Phoenix (internal) endpoint to persist interactions
    PHOENIX_AI_INTERACTIONS_URL: AnyHttpUrl
    PHOENIX_API_TOKEN: str

    # Defaults
    DEFAULT_LLM_MODEL: str = "gemini-2.5-flash"
    THINKING_BUDGET: int = 0
    class Config:
        env_file = ".env"
