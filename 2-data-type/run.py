import os

dir_path = os.path.dirname(__file__)
os.chdir(dir_path)

os.system("go run main.go")
