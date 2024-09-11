import json
from typing import Union


def read_json(filename: str):
    with open(filename, "r+") as file:
        content = file.read()
    return json.loads(content)


def write_json(filename: str, content: Union[dict, list]):
    with open(filename, "w+") as file:
        file.write(json.dumps(content))
