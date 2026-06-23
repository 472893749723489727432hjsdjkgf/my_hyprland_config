import os
import shutil
from pathlib import Path


PATH_ROOT_PROJECT = Path(__file__).resolve().parent
CONF_DIR = Path("~/.config").expanduser()

class Directory:
    def __init__(self, dir_path: Path):

        self.dir_path = dir_path

    def move_dir(self):
        target_dir = CONF_DIR / self.dir_path.name
        
        try:
            
            shutil.move(str(self.dir_path), str(target_dir))
            print(f"Конфиг: {self.dir_path.name} успешно добавлен")
        except FileNotFoundError:
            print(f"Ошибка: файл или папка {self.dir_path.name} не найден")
        except shutil.Error as e:
            print(f"Ошибка перемещения: {e}")

def init():
  
    CONF_DIR.mkdir(parents=True, exist_ok=True)

def main():
    init()
    
  
    for item in PATH_ROOT_PROJECT.iterdir():
        

        if item == Path(__file__).resolve():
            continue
            

        directory_obj = Directory(item)
        directory_obj.move_dir()

if __name__ == "__main__":
    main()
