# utils/phoenix_client.py
import httpx
from typing import Dict, Any

class PhoenixClientError(Exception):
    pass

class PhoenixClient:
    def __init__(self, endpoint: str, token: str, client: httpx.AsyncClient):
        self.endpoint = endpoint
        self.token = token
        self.client = client

    async def persist_interaction(self, payload: Dict[str, Any]) -> Dict[str, Any]:
        headers = {"Authorization": f"Bearer {self.token}", "Content-Type": "application/json"}
        try:
            resp = await self.client.post(self.endpoint, json=payload, headers=headers, timeout=5.0)
            resp.raise_for_status()
            return resp.json()
        except httpx.HTTPError as exc:
            # do not raise (persist failure shouldn't block core answer) — log or return error info
            raise PhoenixClientError(f"Failed to persist to Phoenix: {exc}") from exc
