input_file = 'day1_input.txt'

number_list = []
with open(input_file) as file:
    for line in file:
        number_list.append(int(line.strip()))

for first_num in number_list:
    for second_num in number_list:
        if first_num + second_num == 2020:
            print(first_num, second_num, first_num * second_num)

for first_num in number_list:
    for second_num in number_list:
        for third_num in number_list:
            if first_num + second_num + third_num == 2020:
                print(first_num, second_num, third_num, first_num * second_num * third_num)