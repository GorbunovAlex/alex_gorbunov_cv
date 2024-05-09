from fastapi import FastAPI

from backend.views.articles import ArticlesView

def init_app():
    app = FastAPI()
    app.include_router(ArticlesView().router)
    return app


app = init_app()