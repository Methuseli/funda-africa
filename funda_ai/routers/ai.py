# routers/ai.py
from fastapi import APIRouter, Depends, HTTPException
from models.schemas import QueryRequest, QueryResponse
from services.orchestration import Orchestrator
from fastapi import Request

router = APIRouter()

# we will use Request.app.state to fetch orchestrator instance
def get_orchestrator(request: Request) -> Orchestrator:
    orch = request.app.state.orchestrator
    if orch is None:
        raise HTTPException(status_code=500, detail="Orchestrator not initialized.")
    return orch

@router.post("/ai/query", response_model=QueryResponse)
async def ai_query(req: QueryRequest, orchestrator: Orchestrator = Depends(get_orchestrator)):
    try:
        return await orchestrator.handle_query(req)
    except Exception as exc:
        # map to HTTP 500 for now; refine with custom exceptions as you grow
        raise HTTPException(status_code=500, detail=str(exc))
