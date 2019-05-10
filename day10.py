class Point(object):
    def __init__(self, position, velocity):
        self.position = position
        self.velocity = velocity

    def __str__(self):
        return('position=<{}, {}> velocity=<{}, {}>'.format(self.position[0], self.position[1],
                                                            self.velocity[0], self.velocity[1]))


def load_point_list():
    point_list = []
    # input_file = 'day10_test_input.txt'
    input_file = 'day10_input.txt'
    with open(input_file, 'r') as fd:
        for line in fd:
            pos_x = int(line[10:].split(',')[0])
            pos_y = int(line.split(',')[1].strip().split('>')[0])
            vel_x = int(line.split('<')[-1].strip().split(',')[0])
            vel_y = int(line.split('<')[-1].strip().split(',')[1].strip().split('>')[0])

            new_point = Point((pos_x, pos_y), (vel_x, vel_y))

            point_list.append(new_point)
    return point_list


def print_coord_list(point_list):
    for point in point_list:
        print(point)


def print_point_plot(point_plot):
    for y in point_plot:
        print(''.join(y))


def get_plot_size(point_list):
    max_x = 0
    min_x = 0
    for point in point_list:
        if point.position[0] > max_x:
            max_x = point.position[0]
        if point.position[0] < min_x:
            min_x = point.position[0]

    max_y = 0
    min_y = 0
    for point in point_list:
        if point.position[1] > max_y:
            max_y = point.position[1]
        if point.position[1] < min_y:
            min_y = point.position[1]

    total_size_x = abs(max_x) + abs(min_x)
    total_size_y = abs(max_y) + abs(min_y)
    return total_size_x, total_size_y, min_x, min_y


def plot_points(point_list):
    size_x, size_y, zero_ref_x, zero_ref_y = get_plot_size(point_list)
    point_plot = [['.' for _ in range(size_x+1)] for _ in range(size_y+1)]

    zero_ref_x = abs(zero_ref_x)
    zero_ref_y = abs(zero_ref_y)

    for point in point_list:
        point_plot[point.position[1]+zero_ref_y][point.position[0]+zero_ref_x] = '#'

    print_point_plot(point_plot)
    print('\n')


def advance_position(point_list):
    for point in point_list:
        new_x = point.position[0] + point.velocity[0]
        new_y = point.position[1] + point.velocity[1]
        point.position = (new_x, new_y)


def main():
    point_list = load_point_list()
    # print_coord_list(point_list)

    # for second in range(5):
    #     print('second: {}'.format(second))
    #     print(get_plot_size(point_list))
    #     plot_points(point_list)
    #     advance_position(point_list)

    for second in range(15000):
        # print(get_plot_size(point_list))
        if get_plot_size(point_list)[0] == 180:
            plot_points(point_list)
            print('second: {}'.format(second))
            break
        advance_position(point_list)


if __name__ == "__main__":
    main()
