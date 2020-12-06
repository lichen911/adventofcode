input_file = 'day5_input.txt'


def load_input():
    seat_list = []
    with open(input_file) as file:
        for line in file:
            seat_list.append(line.strip())
    return seat_list


def get_seat_values(seat):
    high_value = 127
    low_value = 0
    for letter in seat[0:7]:
        if letter == 'F':
            high_value = ((high_value - low_value) // 2) + low_value
        elif letter == 'B':
            low_value = ((high_value - low_value) // 2) + low_value + 1
    row = low_value

    high_value = 7
    low_value = 0
    for letter in seat[7:]:
        if letter == 'L':
            high_value = ((high_value - low_value) // 2) + low_value
        elif letter == 'R':
            low_value = ((high_value - low_value) // 2) + low_value + 1
    column = low_value

    return row, column, row * 8 + column


def print_seating_chart(seating_chart):
    for row in seating_chart:
        print(row)


def find_empty_seat(id_list):
    id_list.sort()
    for index in range(len(id_list)):
        if index == 0 or index == len(id_list) - 1:
            continue
        if id_list[index] - id_list[index-1] != 1:
            return id_list[index] - 1
    return None


def main():
    seat_list = load_input()
    # print(seat_list)
    seating_chart = [['O' for _ in range(8)] for _ in range(128)]
    id_list = []

    highest_seat_id = 0
    for seat in seat_list:
        row, column, id = get_seat_values(seat)
        if id > highest_seat_id:
            highest_seat_id = id
        # print(f'seat: {seat}, row: {row}, column: {column}, id: {id}')
        seating_chart[row][column] = id
        id_list.append(id)

    # print_seating_chart(seating_chart)
    print(f'highest seat id: {highest_seat_id}')

    empty_seat = find_empty_seat(id_list)
    print(f'empty seat: {empty_seat}')


if __name__ == '__main__':
    main()