INPUT_FILE = "day1_input.txt"
# INPUT_FILE = "day1_input_sample.txt"


def get_input(file_name):
    with open(file_name, "r") as fd:
        puzzle_input = fd.read().split()
    return list(map(int, puzzle_input))


def get_depth_increases(puzzle_input):
    count = 0
    for i in range(1, len(puzzle_input)):
        current = puzzle_input[i]
        previous = puzzle_input[i-1]
        if current > previous:
            count += 1
    return count


def get_depth_sliding_window(puzzle_input):
    count = 0
    for i in range(1, len(puzzle_input)):
        current = puzzle_input[i:i+3]
        previous = puzzle_input[i-1:i+2]
        if sum(current) > sum(previous):
            count += 1
    return count


def main():
    puzzle_input = get_input(INPUT_FILE)
    # print(puzzle_input)

    depth_increases = get_depth_increases(puzzle_input)
    print(f"{depth_increases = }")

    depth_sliding_window = get_depth_sliding_window(puzzle_input)
    print(f"{depth_sliding_window = }")


if __name__ == "__main__":
    main()
