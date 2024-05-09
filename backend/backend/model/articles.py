from sqlalchemy.orm import Session
from backend.core.database import Database
from .models.articles import Article

class ArticleModel:
    def __init__(self):
        self.db: Session = Database().session


    def get_articles(self):
        return self.db.query(Article).all()