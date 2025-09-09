# services/wolfram_service.py
import httpx
from typing import Any, Dict

class WolframError(Exception):
    pass

class WolframService:
    def __init__(
        self,
        appid: str,
        client: httpx.AsyncClient,
        base_url: str = "https://api.wolframalpha.com/v2",
    ):
        self.appid = appid
        self.client = client
        self.base_url = base_url

    async def query(self, input_text: str) -> Dict[str, Any]:
        """Query WolframAlpha and return JSON response."""
        params = {
            "appid": self.appid,
            "input": input_text,
            "output": "JSON",
            "format": "plaintext,mathml",
        }
        try:
            # Use client's configured timeout (set at startup)
            resp = await self.client.get(f"{self.base_url}/query", params=params)
            resp.raise_for_status()
            return resp.json()
        except httpx.HTTPError as exc:
            raise WolframError(f"Wolfram request failed: {exc}") from exc
