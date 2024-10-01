# print with variables, separation, ending
greeting = "Hello"
name = "Bob"

print(f"{greeting} {name}")
print("Hello", name, end=" | ")
print("Hello", name, sep=":")


# rang loop
range_1 = range(10, -10, -2)
range_2 = range(10, 20, 2)


# map: apply a function in a iterable object (string, set, tuple, list)
strings = ["this", "is", "example", "for", "a", "map", "function"]
length = map(len, strings)

# map: map with args function
def test_map(string: str):
    return "custom string message:", string
map_results = map(lambda string: test_map(string), strings)

# filter: run a function in a iterable object & remove those object that has False return value
strings = ["this", "is", "example", "for", "a", "map", "function"]
def longer_than_4(string: str) -> bool:
    return len(strings) > 4
length = filter(longer_than_4, strings)

# filter: filter with 2 args function
strings = ["this", "is", "example", "for", "a", "map", "function"]
compare_string = ""
def compare_data(arg_1: str, args_2: str) -> bool:
    return arg_1 == args_2
filterd_string = filter(lambda string: compare_data(string, compare_string), strings)


# sorted: sort iterable object ascending or descending
people = [
    {"name": "Alice", "age": 30},
    {"name": "Bob", "age": 25},
    {"name": "Zi", "age": 27},
    {"name": "Long", "age": 55},
    {"name": "Paul", "age": 20},
]
sorted_nums = sorted(people, key=lambda person: person["age"], reverse=True)


# enumerate: get index from list
strings = ["this", "is", "example", "for", "a", "map", "function"]
for index, value in enumerate(strings):
    print(index, value, sep=":")

# zip: mapping 2 iterable to a combine iterable object
names = [ "Alice", "Bob", "Zi", "Long", "Paul", ]
ages = [ 30, 25, 27, 55, 20, ]
people = list(zip(names, ages))
