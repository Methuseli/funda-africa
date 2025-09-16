# services/orchestration.py
from typing import Dict, Any
from models.schemas import QueryRequest, QueryResponse
from services.wolfram_service import WolframService, WolframError
from services.wolfram_llm_service import WolframLLMService
from services.gemini_service import GeminiService, LLMError
from utils.phoenix_client import PhoenixClient, PhoenixClientError


class OrchestrationError(Exception):
    pass


class Orchestrator:
    """
    The main orchestration class that routes queries to the appropriate
    backend service (Wolfram Alpha, Wolfram LLM, or Gemini).
    """

    def __init__(self,
                 wolfram_alpha: WolframService,
                 wolfram_llm: WolframLLMService,
                 gemini: GeminiService,
                 phoenix: PhoenixClient = None):
        """
        Initializes the Orchestrator with instances of each service.
        
        :param wolfram_alpha: Service for the Wolfram Alpha Full Results API.
        :param wolfram_llm: Service for the Wolfram LLM API.
        :param gemini: Service for the Gemini API.
        :param phoenix: Client for a persistence service (optional).
        """
        self.wolfram_alpha = wolfram_alpha
        self.wolfram_llm = wolfram_llm
        self.gemini = gemini
        self.phoenix = phoenix

    async def handle_query(self, req: QueryRequest) -> QueryResponse:
        """
        Handles an incoming query by routing it to the correct service based on the mode.
        """
        metadata: Dict[str, Any] = {}
        model_used_parts = []
        answer = ""

        if req.mode == "wolfram":
            try:
                wf = await self.wolfram_alpha.query(req.prompt)
                metadata["wolfram_alpha"] = wf
                model_used_parts.append("wolfram_alpha")
                answer = self._extract_wolfram_plaintext(wf) or "No plaintext result from Wolfram Alpha."
            except WolframError as e:
                raise OrchestrationError(f"Wolfram Alpha request failed: {e}")

        elif req.mode == "llm":
            # Determine which LLM to use based on the user's preference or query type
            # For this example, we'll use a simple rule
            if self._is_stem_query(req.prompt):
                llm_service = self.wolfram_llm
                model_used_parts.append("wolfram_llm")
            else:
                llm_service = self.gemini
                model_used_parts.append("gemini")

            system = self._system_prompt(req.context)
            try:
                llm_res = await llm_service.generate(f"{system}\n{req.prompt}")
                metadata["llm_raw"] = llm_res["raw"]
                answer = llm_res["output"]
            except LLMError as e:
                raise OrchestrationError(f"LLM request failed: {e}")

        elif req.mode == "hybrid":
            # 1) Call Wolfram Alpha for factual data
            try:
                wf = await self.wolfram_alpha.query(req.prompt)
                metadata["wolfram_alpha"] = wf
                model_used_parts.append("wolfram_alpha")
            except WolframError as e:
                metadata["wolfram_error"] = str(e)
                wf = None

            # 2) Call the appropriate LLM with Wolfram Alpha evidence
            system = self._system_prompt(req.context)
            evidence = self._wolfram_brief(wf) if wf else "No Wolfram evidence available."

            # Determine which LLM to use for the explanation
            if self._is_stem_query(req.prompt):
                llm_service = self.wolfram_llm
                model_used_parts.append("wolfram_llm")
            else:
                llm_service = self.gemini
                model_used_parts.append("gemini")

            llm_input = f"Question:\n{req.prompt}\n\nWolfram evidence (if any):\n{evidence}\n\nPlease give a step-by-step solution in the specified style."
            try:
                llm_res = await llm_service.generate(f"{system}\n{llm_input}")
                metadata["llm_raw"] = llm_res["raw"]
                answer = llm_res["output"]
            except LLMError as e:
                raise OrchestrationError(f"LLM request failed: {e}")
        else:
            raise OrchestrationError(f"Unsupported mode: {req.mode}")

        model_used = " + ".join(model_used_parts) or "none"

        if self.phoenix:
            persist_payload = {
                "user_id": str(req.user_id) if req.user_id else None,
                "prompt": req.prompt,
                "response": answer,
                "model_used": model_used,
                "metadata": metadata
            }
            try:
                await self.phoenix.persist_interaction(persist_payload)
            except PhoenixClientError:
                pass

        return QueryResponse(answer=answer, model_used=model_used, metadata=metadata)

    def _is_stem_query(self, prompt: str) -> bool:
        """
        Simple heuristic to check if a query is STEM-related.
        This can be replaced with a more sophisticated NLP model.
        """
        stem_keywords = ["math", "physics", "chemistry", "biology", "engineering", "formula", "equation", "graph", "calculate", "solve"]
        return any(keyword in prompt.lower() for keyword in stem_keywords)

    def _system_prompt(self, context: Dict[str, Any]):
        if context and context.get("syllabus") == "ZIMSEC-4004":
            return ("You are an exam-focused math tutor for Zimbabwe O-Level (ZIMSEC 4004). "
                    "Provide short method steps and a clearly labelled final answer. Use language suitable for Form 3/4 students.")
        return "You are a helpful and clear tutor. Provide step-by-step answers."

    def _extract_wolfram_plaintext(self, wf_json: Dict[str, Any]) -> str:
        try:
            pods = wf_json.get("queryresult", {}).get("pods", [])
            texts = []
            for pod in pods:
                for sub in pod.get("subpods", []):
                    if "plaintext" in sub and sub["plaintext"]:
                        texts.append(f"{pod.get('title','')}: {sub['plaintext']}")
            return "\n".join(texts) if texts else ""
        except Exception:
            return ""

    def _wolfram_brief(self, wf_json: Dict[str, Any]) -> str:
        try:
            pods = wf_json.get("queryresult", {}).get("pods", [])
            brief = []
            for pod in pods[:4]:
                title = pod.get("title", "")
                text = ""
                for sub in pod.get("subpods", []):
                    if sub.get("plaintext"):
                        text = sub["plaintext"]
                        break
                if text:
                    brief.append(f"{title}: {text}")
            return "\n".join(brief)
        except Exception:
            return ""
