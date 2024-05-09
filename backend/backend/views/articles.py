from fastapi import APIRouter, HTTPException

from backend.controllers.articles import ArticleController
from backend.model.articles import ArticleModel


class ArticlesView:

    def __init__(self):
        self.controller = ArticleController()
        self.router = APIRouter()
        self.router.add_api_route('/articles', self.get_articles, methods=['GET'])


    def get_articles(self):
        articles = self.controller.get_articles()
        if not articles:
            raise HTTPException(status_code=404, detail="Articles not found")
        return {
            "data": articles
        }