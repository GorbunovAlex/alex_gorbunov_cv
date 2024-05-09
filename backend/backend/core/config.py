import configparser

from backend.definitions import CONFIG_PATH

class ConfigMeta(type):
    _instances = {}
    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            instance = super().__call__(*args, **kwargs)
            cls._instances[cls] = instance
        return cls._instances[cls]
    

class Config(metaclass=ConfigMeta):
    def __init__(self):
         self.config = configparser.ConfigParser()
         self.file = open(CONFIG_PATH)
         self.config = self.config.read_file(self.file)

    def get_config(self, section, option):
        return self.config.get(section, option)