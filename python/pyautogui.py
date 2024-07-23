import pyautogui
from time import sleep


def move_and_click(position: tuple):
    pyautogui.moveTo(position)
    pyautogui.click()


def drag_and_drop_file(from_position: tuple, to_position: tuple):
    sleep(1)
    pyautogui.moveTo(from_position)
    pyautogui.dragTo(to_position, duration=2, button="left")
    sleep(1)


def search_file(table_name: str):
    sleep(1)
    with pyautogui.hold("ctrl"):
        pyautogui.press("f")
    with pyautogui.hold("ctrl"):
        pyautogui.press("a")
    pyautogui.write(table_name)
    sleep(1)
    pyautogui.press("enter")
