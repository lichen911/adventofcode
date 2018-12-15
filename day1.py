input_file = 'day1_input.txt'
# input_file = 'day1_input_test.txt'

total = 0
freq_counts = {}
stop_loop = False

while not stop_loop:
    with open(input_file, 'r') as fd:
        for number in fd:
            print(number)
            number = int(number)
            total += number
            try:
                freq_counts[total] += 1
            except KeyError:
                freq_counts[total] = 0

            if freq_counts[total] >= 1:
                print('first dup found! - %s' % total)
                stop_loop = True
                break

print(total)
