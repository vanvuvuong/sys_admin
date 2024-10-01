import csv
import glob

files = [
    "data/pattern*.csv",
]


def get_csv_file_path(pattern_str: str):
    return glob.glob(pattern_str)


def read_csv_file(filename: str):
    rows = []
    with open(file=filename, mode="r+", encoding="utf-8") as csv_file:
        content = csv.reader(
            csv_file,
            delimiter=",",
            quotechar='"',
        )
        # run_times = 0
        for row in content:
            # if run_times >= 2:
            #     continue
            rows.append(row)
            # run_times += 1
    return rows


def write_csv_file(filename: str, rows: list):
    with open(file=filename, mode="w+", encoding="utf-8", newline="") as csv_file:
        writer = csv.writer(csv_file, delimiter=",", quotechar='"')
        writer.writerows(rows)
