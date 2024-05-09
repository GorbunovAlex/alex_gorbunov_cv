from backend.model.articles import ArticleModel

class ArticleController:
    def __init__(self):
        self.model = ArticleModel()

    def get_articles(self):
        return self.model.get_articles()