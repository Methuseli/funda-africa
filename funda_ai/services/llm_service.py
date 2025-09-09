# services/llm_service.py
import httpx
from typing import Dict, Any

class LLMError(Exception):
    pass

class LLMService:
    """Simple OpenAI-compatible async wrapper. Replace or extend for other providers."""
    def __init__(self, api_key: str, client: httpx.AsyncClient, api_url: str, default_model: str):
        self.api_key = api_key
        self.client = client
        self.api_url = api_url
        self.default_model = default_model

    async def chat(self, system_prompt: str, user_prompt: str, model: str = None) -> Dict[str, Any]:
        model = model or self.default_model
        payload = {
            "model": model,
            "messages": [
                {"role": "system", "content": system_prompt},
                {"role": "user", "content": user_prompt}
            ],
            "max_tokens": 1200
        }
        headers = {"Authorization": f"Bearer {self.api_key}"}
        try:
            resp = await self.client.post(self.api_url, json=payload, headers=headers, timeout=30.0)
            resp.raise_for_status()
            return resp.json()
        except httpx.HTTPError as exc:
            raise LLMError(f"LLM request failed: {exc}") from exc
