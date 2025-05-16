from fastapi import FastAPI, Header, Cookie, Request
from typing import Optional
from pydantic import BaseModel
import json

app = FastAPI()

class ObjectModel(BaseModel):
    integer: int
    string: str
    boolean: Optional[bool] = False


@app.post("/{pathInteger}/{pathString}")
async def get_item(
        # Request body
        request: ObjectModel,
        # Path parameters
        pathInteger: int,
        pathString: str,
        # Header && Cookie parameters
        x_request_id: Optional[str] = Header(None, convert_underscores=False),
        session_id: Optional[str] = Cookie(None),
        # Parameters
        q: Optional[str] = None,
) -> ObjectModel:
    response = ObjectModel(
        integer=pathInteger,
        string=pathString,
        boolean=request.boolean
    )
    return response

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)