# services/gemini_service.py
from typing import Dict, Any, Optional
import google.generativeai as genai
from .base_llm import BaseLLM


class LLMError(Exception):
    """Custom exception for LLM errors."""
    pass


class GeminiService(BaseLLM):
    """
    Service class to interact with the Gemini API, implementing the BaseLLM interface.
    """

    def __init__(self, api_key: str, model: Optional[str] = 'gemini-pro'):
        """
        Initializes the Gemini service with an API key and an optional model.

        :param api_key: The API key for the Gemini service.
        :param model: The name of the Gemini model to use.
        """
        super().__init__(model)
        genai.configure(api_key=api_key)

    async def generate(self, prompt: str, **kwargs) -> Dict[str, Any]:
        """
        Generates a response from the Gemini model.

        :param prompt: The input prompt for the model.
        :param kwargs: Optional keyword arguments, e.g., 'temperature'.
        :return: A dictionary containing the output and the raw response.
        :raises LLMError: If the API request fails.
        """
        try:
            model_instance = genai.GenerativeModel(self.model)
            response = await model_instance.generate_content_async(
                prompt,
                generation_config=kwargs
            )

            # Extract the text from the response
            output_text = response.text

            return {
                "output": output_text,
                "raw": response.to_dict()
            }
        except Exception as e:
            raise LLMError(f"Gemini API request failed: {e}")
