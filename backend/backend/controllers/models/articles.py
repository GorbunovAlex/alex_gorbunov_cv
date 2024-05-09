from pydantic import BaseModel

class ArticleBase(BaseModel):
    id: int
    title: str
    content: str
    published: bool
    created_at: str
    updated_at: str

    class Config:
        orm_mode = True