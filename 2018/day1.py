input_file = 'day1_input.txt'
<<<<<<< HEAD
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
=======

mass_list = []
with open(input_file, 'r') as fd:
    for mass in fd:
        mass_list.append(int(mass))

# Part 1
fuel_sum = 0
for mass in mass_list:
    fuel_sum += int(mass/3) - 2

print(fuel_sum)


# Part 2
def get_fuel_req(input_mass):
    def _fuel_req(current_mass, fuel_total):
        # print('current mass {}, fuel total {}'.format(current_mass, fuel_total))
        if current_mass <= 0:
            return fuel_total
        else:
            fuel_req = int(current_mass / 3) - 2
            return _fuel_req(fuel_req, fuel_total + max(0, fuel_req))

    return _fuel_req(input_mass, 0)


fuel_sum = 0
for mass in mass_list:
    fuel_sum += get_fuel_req(mass)

print(fuel_sum)
>>>>>>> ad2019
