from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles

app = FastAPI()

app.mount("/", StaticFiles(directory="public", html=True), name="public")


@app.get("/api/v1/hello")
def hello():
    return {"Hello": "World"}
