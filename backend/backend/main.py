from fastapi import FastAPI

from backend.views.articles import ArticlesView

app = FastAPI()

ArticlesView()