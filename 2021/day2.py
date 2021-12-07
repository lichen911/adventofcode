from utils import get_input


def follow_directions(direction_list):
    position = {
        "horizontal": 0,
        "depth": 0
    }
    for step in direction_list:
        step_list = step.split()
        # print(f"{step_list = }")
        if len(step_list) == 2:
            direction = step_list[0]
            count = int(step_list[1])

            if direction == "up":
                position["depth"] -= count
            elif direction == "down":
                position["depth"] += count
            elif direction == "forward":
                position["horizontal"] += count
            elif direction == "backward":
                position["horizontal"] -= count
            else:
                print("Unknown direction")

    return position


def follow_directions_with_aim(direction_list):
    position = {
        "horizontal": 0,
        "depth": 0,
        "aim": 0
    }

    for step in direction_list:
        step_list = step.split()
        # print(f"{step_list = }")
        if len(step_list) == 2:
            direction = step_list[0]
            count = int(step_list[1])

            if direction == "up":
                position["aim"] -= count
            elif direction == "down":
                position["aim"] += count
            elif direction == "forward":
                position["horizontal"] += count
                position["depth"] += position["aim"] * count
            # elif direction == "backward":
            #     position["horizontal"] -= count
            else:
                print("Unknown direction")

    return position


def main():
    # puzzle_input = get_input(day=2, sample=True)
    puzzle_input = get_input(day=2, sample=False)
    print(puzzle_input)

    result = follow_directions(puzzle_input)
    print(result)
    print(f"final result: {result['depth'] * result['horizontal']}")

    result = follow_directions_with_aim(puzzle_input)
    print(result)
    print(f"final result: {result['depth'] * result['horizontal']}")


if __name__ == "__main__":
    main()
