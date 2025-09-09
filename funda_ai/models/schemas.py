# models/schemas.py
from typing import Optional, Dict, Any, Literal
from pydantic import BaseModel
from uuid import UUID

ModeLiteral = Literal["llm", "wolfram", "hybrid"]

class QueryRequest(BaseModel):
    user_id: Optional[UUID]  # optional - your Phoenix IDs are binary_id
    prompt: str
    mode: ModeLiteral = "hybrid"
    context: Optional[Dict[str, Any]] = None

class QueryResponse(BaseModel):
    answer: str
    model_used: str
    metadata: Optional[Dict[str, Any]] = {}
