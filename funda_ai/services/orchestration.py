# services/orchestration.py
from typing import Dict, Any
from models.schemas import QueryRequest, QueryResponse
from services.wolfram_service import WolframService, WolframError
from services.llm_service import LLMService, LLMError
from utils.phoenix_client import PhoenixClient, PhoenixClientError

class OrchestrationError(Exception):
    pass

class Orchestrator:
    def __init__(self, wolfram: WolframService, llm: LLMService, phoenix: PhoenixClient = None):
        self.wolfram = wolfram
        self.llm = llm
        self.phoenix = phoenix

    async def handle_query(self, req: QueryRequest) -> QueryResponse:
        metadata: Dict[str, Any] = {}
        model_used_parts = []

        # Decide execution path
        if req.mode == "wolfram":
            # Direct computation
            try:
                wf = await self.wolfram.query(req.prompt)
                metadata["wolfram"] = wf
                model_used_parts.append("wolfram")
                answer = self._extract_wolfram_plaintext(wf) or "No plaintext result from Wolfram."
            except WolframError as e:
                raise OrchestrationError(str(e))

        elif req.mode == "llm":
            # Pure LLM explanation/generation
            system = self._system_prompt(req.context)
            try:
                llm_res = await self.llm.chat(system, req.prompt)
                metadata["llm_raw"] = llm_res
                model_used_parts.append("llm")
                answer = self._extract_llm_text(llm_res)
            except LLMError as e:
                raise OrchestrationError(str(e))

        elif req.mode == "hybrid":
            # 1) call Wolfram for facts/calculation
            try:
                wf = await self.wolfram.query(req.prompt)
                metadata["wolfram"] = wf
                model_used_parts.append("wolfram")
            except WolframError as e:
                # still try LLM if wolfram fails (configurable)
                metadata["wolfram_error"] = str(e)
                wf = None

            # 2) call LLM with wolfram evidence
            system = self._system_prompt(req.context)
            evidence = self._wolfram_brief(wf) if wf else ""
            llm_input = f"Question:\n{req.prompt}\n\nWolfram evidence (if any):\n{evidence}\n\nPlease give step-by-step solution in ZIMSEC O-Level style."
            try:
                llm_res = await self.llm.chat(system, llm_input)
                metadata["llm_raw"] = llm_res
                model_used_parts.append("llm")
                answer = self._extract_llm_text(llm_res)
            except LLMError as e:
                raise OrchestrationError(str(e))
        else:
            raise OrchestrationError(f"Unsupported mode: {req.mode}")

        model_used = " + ".join(model_used_parts) or "none"

        # Persist to Phoenix (best-effort)
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
                # swallow — persistence shouldn't block the user response
                pass

        return QueryResponse(answer=answer, model_used=model_used, metadata=metadata)

    def _system_prompt(self, context: Dict[str, Any]):
        # Example system prompt tuned for ZIMSEC O-Level style if context contains syllabus
        if context and context.get("syllabus") == "ZIMSEC-4004":
            return ("You are an exam-focused math tutor for Zimbabwe O-Level (ZIMSEC 4004). "
                    "Provide short method steps and a clearly labelled final answer. Use language suitable for Form 3/4 students.")
        return "You are a helpful math tutor. Provide step-by-step answers."

    def _extract_llm_text(self, llm_response: Dict[str, Any]) -> str:
        # Expect OpenAI-style response
        try:
            return llm_response["choices"][0]["message"]["content"].strip()
        except Exception:
            return str(llm_response)

    def _extract_wolfram_plaintext(self, wf_json: Dict[str, Any]) -> str:
        # Try to pull a plain-text worth of answer from pods
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
        # Return a compact representation of key pods for LLM
        try:
            pods = wf_json.get("queryresult", {}).get("pods", [])
            brief = []
            for pod in pods[:4]:  # limit size
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
