from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

from backend.core.config import Config

Base = declarative_base()


class DatabaseMeta(type):
    _instances = {}
    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            instance = super().__call__(*args, **kwargs)
            cls._instances[cls] = instance
        return cls._instances[cls]


class Database(metaclass=DatabaseMeta):
    def __init__(self):
        self.config = Config()
        SQLALCHEMY_DATABASE_URL = f"postgresql://{self.config.get_config('Database', 'db_user')}:{self.config.get_config('Database', 'db_password')}@{self.config.get_config('Database', 'host')}:{self.config.get_config('Database', 'port')}/{self.config.get_config('Database', 'db_name')}"
        engine = create_engine(
            SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False}
        )
        self.session = sessionmaker(autocommit=False, autoflush=False, bind=engine)