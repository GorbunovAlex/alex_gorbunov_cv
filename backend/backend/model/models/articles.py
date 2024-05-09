from sqlalchemy import Boolean, Column, Integer, String

from backend.core.database import Base


class Article(Base):
    __tablename__ = "articles"

    id = Column(Integer, primary_key=True, index=True)
    title = Column(String)
    content = Column(String)
    published = Column(Boolean)
    created_at = Column(String)
    updated_at = Column(String)