input_file = 'day6_input.txt'


def load_input():
    with open(input_file) as file:
        input = file.readlines()
    return input


def part1(customs_input):
    answer_count_list = []
    current_group = []

    for line in customs_input:
        if line == '\n':
            answer_count_list.append(len(set(current_group)))
            current_group = []
            continue
        else:
            current_group = current_group + list(line.strip())
            # print(current_group)

    if current_group:
        answer_count_list.append(len(set(current_group)))

    return answer_count_list


def part2(customs_input):
    answer_count_list = []
    current_group = []

    for line in customs_input:
        if line == '\n':
            common_answers = current_group[0].intersection(*current_group[1:])
            answer_count_list.append(len(common_answers))
            current_group = []
            continue
        else:
            current_group.append(set(line.strip()))
            # print(current_group)

    if current_group:
        common_answers = current_group[0].intersection(*current_group[1:])
        answer_count_list.append(len(common_answers))

    return answer_count_list


def main():
    customs_input = load_input()

    answer_count_list = part1(customs_input)
    print(f'part1 answer list sum: {sum(answer_count_list)}')

    answer_count_list = part2(customs_input)
    print(f'part2 answer list sum: {sum(answer_count_list)}')


if __name__ == '__main__':
    main()
