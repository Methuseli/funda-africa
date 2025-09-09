# core/config.py
from pydantic import BaseSettings, AnyHttpUrl

class Settings(BaseSettings):
    # External APIs
    WOLFRAM_APPID: str 
    LLM_API_KEY: str  # e.g. OpenAI
    LLM_API_URL: str = "https://api.openai.com/v1/chat/completions"
    # Phoenix (internal) endpoint to persist interactions
    PHOENIX_AI_INTERACTIONS_URL: AnyHttpUrl
    PHOENIX_API_TOKEN: str

    # Defaults
    DEFAULT_LLM_MODEL: str = "gemini-2.5-flash"

    class Config:
        env_file = ".env"
