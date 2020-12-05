input_file = 'day3_input.txt'
slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]


def print_hillside(hillside):
    for y in range(len(hillside)):
        print(''.join(hillside[y]))
    print()


def go_downhill(hillside, slope):
    slope_right = slope[0]
    slope_down = slope[1]
    tree_count = 0
    x_pos = 0
    for y in range(slope_down, len(hillside), slope_down):
        x_pos = x_pos + slope_right
        x_idx = x_pos % len(hillside[y])
        if hillside[y][x_idx] == '#':
            tree_count = tree_count + 1
    return tree_count


def main():
    hillside = []

    with open(input_file) as file:
        for line in file:
            hillside.append(list(line.strip()))

    tree_multiple = 1
    for slope in slopes:
        tree_count = go_downhill(hillside, slope)
        print(f'Number of trees for slope {slope}: {tree_count}')
        tree_multiple = tree_multiple * tree_count
    print(f'Tree counts multiplied: {tree_multiple}')


if __name__ == '__main__':
    main()
