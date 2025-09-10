# services/llm_service.py
import httpx
from typing import Dict, Any
from google.genai import types


class LLMError(Exception):
    pass

class LLMService:
    """Simple OpenAI-compatible async wrapper. Replace or extend for other providers."""
    def __init__(self, default_model: str, thinking_budget: int):
        self.default_model = default_model
        self.thinking_budget = thinking_budget

    def chat(self, system_prompt: str, user_prompt: str = None, model: str = None) -> Dict[str, Any]:
        model = model or self.default_model
        if user_prompt == None:
            prompts = system_prompt
        else:
            prompts = user_prompt
        try:
            response = self.client.models.generate_content(
                model=self.default_model,
                contents=prompts,
                config=types.ThinkingConfig(thinking_budget=self.thinking_budget)
            )
            return response.text
        except httpx.HTTPError as exc:
            raise LLMError(f"LLM request failed: {exc}") from exc
