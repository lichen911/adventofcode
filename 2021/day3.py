from utils import get_input


def get_occurrences(diagnostic_report):
    occurrences = [{0: 0, 1: 0} for _ in range(len(diagnostic_report[0]))]

    for row in diagnostic_report:
        # print(f"{row = }")
        for index, bit in enumerate(row):
            bit = int(bit)
            # print(f"{index = }, {bit = }")
            occurrences[index][bit] += 1

    return occurrences


def get_rates(occurrences):
    gamma_rate = ""
    epsilon_rate = ""

    for bit in occurrences:
        if bit[0] > bit[1]:
            most_common = 0
            least_common = 1
        else:
            most_common = 1
            least_common = 0

        gamma_rate += str(most_common)
        epsilon_rate += str(least_common)

    return gamma_rate, epsilon_rate


def main():
    # puzzle_input = get_input(day=3, sample=True)
    puzzle_input = get_input(day=3, sample=False)
    print(f"{puzzle_input = }")

    occurrences = get_occurrences(puzzle_input)
    print(f"{occurrences = }")

    gamma_rate, epsilon_rate = get_rates(occurrences)
    print(f"{gamma_rate = }, {epsilon_rate = }")

    print(f"final result: {int(gamma_rate, 2) * int(epsilon_rate, 2)}")


if __name__ == "__main__":
    main()
