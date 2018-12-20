input_file = 'day6_input.txt'
# input_file = 'day6_input_test.txt'


# return manhattan distance between two points
def get_manhattan_distance(source, destination):
    x_distance = abs(destination[0] - source[0])
    y_distance = abs(destination[1] - source[1])
    return x_distance + y_distance


def get_dist_to_all_coords(source, coord_list):
    total_dist = 0
    for coord in coord_list:
        total_dist += get_manhattan_distance(source, coord)
    return total_dist


def print_grid(manhattan_grid):
    # print the grid for debugging
    for row in range(len(manhattan_grid[0])):
        for col in range(len(manhattan_grid)):
            if manhattan_grid[col][row]['tied_for_distance']:
                print(' . ', end='')
            else:
                print(' %s ' % manhattan_grid[col][row]['closest_node'], end='')
        print()
    print()

def main():
    # read input file and create a list of tuples containing coordinates
    coordinate_list = []
    with open(input_file, 'r') as fd:
        for coord in fd:
            coord_tuple = tuple(map(int, coord.rstrip('\n').replace(' ', '').split(',')))
            coordinate_list.append(coord_tuple)

    # build the grid
    manhattan_grid = [[{'distance_list': [],
                        'closest_node': None,
                        'tied_for_distance': False,
                        'dist_to_all': None} for _ in range(400)] for _ in range(400)]
    # manhattan_grid = [[{'distance_list': [],
    #                     'closest_node': None,
    #                     'tied_for_distance': False,
    #                     'dist_to_all': None} for _ in range(10)] for _ in range(10)]
    # build a dict containing the number of closest points for a particular coordinate
    area_count_by_point = {}
    part_b_total = 0

    # loop through each position in the grid and determine how far it is from each point in the coordinate list
    for col in range(len(manhattan_grid)):
        for row in range(len(manhattan_grid[col])):
            shortest_distance = None
            closest_node = None
            for i, coordinate in enumerate(coordinate_list):
                current_distance = get_manhattan_distance((col, row), coordinate)
                manhattan_grid[col][row]['distance_list'].append(current_distance)

                if (col, row) == coordinate:
                    closest_node = i
                    break

                if shortest_distance:
                    if current_distance < shortest_distance:
                        shortest_distance = current_distance
                        closest_node = i
                else:
                    shortest_distance = current_distance
                    closest_node = i

            # determine of this shortest distance is tied with other distances, if so mark it as tied
            if manhattan_grid[col][row]['distance_list'].count(shortest_distance) > 1:
                manhattan_grid[col][row]['tied_for_distance'] = True

            # set the closest node, then increment the area counter if not tied
            if not manhattan_grid[col][row]['tied_for_distance']:
                manhattan_grid[col][row]['closest_node'] = closest_node
                try:
                    area_count_by_point[closest_node] += 1
                except KeyError:
                    area_count_by_point[closest_node] = 1

            # capture distance to all coordinates for Part B
            # print((col, row))
            dist_to_all = get_dist_to_all_coords((col, row), coordinate_list)
            manhattan_grid[col][row]['dist_to_all'] = dist_to_all
            # print(dist_to_all)
            if dist_to_all < 10000:
                part_b_total += 1

    print_grid(manhattan_grid)

    print('area_count_by_point', area_count_by_point)

    # build list containing points that all have an infinite area
    infinite_area_points = set()
    for row in manhattan_grid[0]:
        infinite_area_points.add(row['closest_node'])
    for row in manhattan_grid[-1]:
        infinite_area_points.add(row['closest_node'])
    for i in range(len(manhattan_grid)):
        infinite_area_points.add(manhattan_grid[i][0]['closest_node'])
        infinite_area_points.add(manhattan_grid[i][len(manhattan_grid[i])-1]['closest_node'])

    print('infinite_area_points', infinite_area_points)

    max_area = None
    for point, count in area_count_by_point.items():
        if point not in infinite_area_points:
            if max_area:
                if count > max_area:
                    max_area = count
            else:
                max_area = count
    print('part a:', max_area)
    print('part b:', part_b_total)

if __name__ == "__main__":
    main()