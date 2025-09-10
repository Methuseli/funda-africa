# services/base_llm.py
from abc import ABC, abstractmethod
from typing import Dict, Any, Optional


class BaseLLM(ABC):
    """
    Abstract base class for all Large Language Model (LLM) services.

    This class defines the standard interface that all LLM services
    (e.g., OpenAI, Anthropic, Wolfram LLM) must implement so they can
    be used interchangeably by the Orchestrator.
    """

    def __init__(self, model: Optional[str] = None):
        """
        Optionally specify a model identifier for the LLM.

        :param model: Optional model name or ID.
        """
        self.model = model

    @abstractmethod
    async def generate(self, prompt: str, **kwargs) -> Dict[str, Any]:
        """
        Generate a response from the LLM.

        :param prompt: Input text or question for the LLM.
        :param kwargs: Extra provider-specific parameters (e.g., temperature, max_tokens).
        :return: A dictionary containing at least:
                 {
                     "output": str,   # the main text result
                     "raw": dict      # full raw provider response (optional, for debugging)
                 }
        """
        ...
