def get_input(day, sample=False):
    if sample:
        file_name = f"day{day}_input_sample.txt"
    else:
        file_name = f"day{day}_input.txt"

    with open(file_name, "r") as fd:
        puzzle_input = fd.read().split('\n')

    return puzzle_input
