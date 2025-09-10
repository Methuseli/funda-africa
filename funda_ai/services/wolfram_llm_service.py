# services/wolfram_llm_service.py
import httpx
from typing import Dict, Any
from .base_llm import BaseLLM

class WolframLLMService(BaseLLM):
    def __init__(self, api_url: str, appid: str, client: httpx.AsyncClient):
        self.api_url = api_url
        self.appid = appid
        self.client = client

    async def generate(self, prompt: str, **kwargs) -> Dict[str, Any]:
        params = {"appid": self.appid, "input": prompt, "output": "JSON"}
        resp = await self.client.get(self.api_url, params=params)
        resp.raise_for_status()
        return resp.json()
