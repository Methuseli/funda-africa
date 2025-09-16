# services/wolfram_llm_service.py
import httpx
from typing import Dict, Any, Optional
from .base_llm import BaseLLM


class LLMError(Exception):
    """Custom exception for LLM errors."""
    pass


class WolframLLMService(BaseLLM):
    """
    Service class to interact with the Wolfram LLM API, implementing the BaseLLM interface.
    """

    def __init__(self, appid: str, client: httpx.AsyncClient, model: Optional[str] = "wolfram-llm"):
        """
        Initializes the Wolfram LLM service.

        :param appid: The AppID for the Wolfram LLM API.
        :param client: An httpx.AsyncClient instance.
        :param model: The name of the Wolfram LLM model.
        """
        super().__init__(model)
        self.appid = appid
        self.client = client
        self.base_url = "https://www.wolframalpha.com/api/v1/llm-api"

    async def generate(self, prompt: str, **kwargs) -> Dict[str, Any]:
        """
        Generates a response from the Wolfram LLM API.

        :param prompt: The input prompt for the model.
        :param kwargs: Optional keyword arguments, e.g., 'maxchars'.
        :return: A dictionary containing the output and the raw response.
        :raises LLMError: If the API request fails.
        """
        params = {
            "appid": self.appid,
            "input": prompt,
        }
        if "maxchars" in kwargs:
            params["maxchars"] = kwargs["maxchars"]
        
        try:
            resp = await self.client.get(self.base_url, params=params)
            resp.raise_for_status()

            # The Wolfram LLM API returns plain text, so we handle it directly
            output_text = resp.text

            return {
                "output": output_text,
                "raw": output_text
            }
        except httpx.HTTPError as exc:
            raise LLMError(f"Wolfram LLM request failed: {exc}") from exc
        except Exception as exc:
            raise LLMError(f"An unexpected error occurred: {exc}") from exc
