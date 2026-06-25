import os
from pathlib import Path


def create_zshrc():
    Path("~.zshrc").touch()
    print("Файл .zshrc создан!")


