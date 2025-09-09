# main.py
import httpx
from fastapi import FastAPI
from contextlib import asynccontextmanager

from core.config import Settings
from services.wolfram_service import WolframService
from services.llm_service import LLMService
from services.orchestration import Orchestrator
from utils.phoenix_client import PhoenixClient
from routers import ai as ai_router


settings = Settings()


@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup: create shared client and services
    http_client = httpx.AsyncClient(timeout=30.0)

    wolfram = WolframService(appid=settings.WOLFRAM_APPID, client=http_client)
    llm = LLMService(
        api_key=settings.LLM_API_KEY,
        client=http_client,
        api_url=settings.LLM_API_URL,
        default_model=settings.DEFAULT_LLM_MODEL,
    )
    phoenix = PhoenixClient(
        endpoint=str(settings.PHOENIX_AI_INTERACTIONS_URL),
        token=settings.PHOENIX_API_TOKEN,
        client=http_client,
    )

    app.state.http_client = http_client
    app.state.orchestrator = Orchestrator(
        wolfram=wolfram,
        llm=llm,
        phoenix=phoenix,
    )

    # yield to let the app run
    yield

    # Shutdown: cleanup
    await http_client.aclose()


def create_app() -> FastAPI:
    app = FastAPI(title="FundaAfrica AI Gateway", lifespan=lifespan)
    app.include_router(ai_router.router)
    return app


app = create_app()
