def load_input_file():
    input_file = 'day12_input.txt'
    # input_file = 'day12_input_test.txt'
    plant_rules = {}

    with open(input_file, 'r') as fd:
        initial_state = list(fd.readline().split()[2].strip())

        for line in fd:
            if line.strip():
                rule_condition = line.split()[0].strip()
                rule_result = line.split()[2].strip()
                plant_rules[rule_condition] = rule_result
    return initial_state, plant_rules


def lpad_list(input_list):
    offset_shift = 0
    while ''.join(input_list[0:3]) != '...':
        input_list.insert(0, '.')
        offset_shift += 1
    return offset_shift


def rpad_list(input_list):
    while ''.join(input_list[len(input_list)-3:]) != '...':
        input_list.append('.')


def get_pot_sum(plant_states, start_offset):
    pot_sum = 0
    for index, plant in enumerate(plant_states):
        if plant == '#':
            pot_sum += (index - start_offset)
    return pot_sum


def get_sum_by_gen(plant_states, start_offset, gen, current_gen):
    pot_sum = 0
    for index, plant in enumerate(plant_states):
        if plant == '#':
            pot_sum += ((index + (gen - current_gen)) - start_offset)
    return pot_sum


def main():
    start_offset = 0
    plant_states, plant_rules = load_input_file()

    print('{}: {}'.format('0'.rjust(2), ''.join(plant_states)))
    for generation in range(1, 101):
        start_offset += lpad_list(plant_states)
        rpad_list(plant_states)

        next_gen = plant_states.copy()
        for pot_idx in range(2, len(plant_states)-2):
            pot_string = plant_states[pot_idx-2:pot_idx+3]
            try:
                next_gen_plant = plant_rules[''.join(pot_string)]
            except KeyError:
                # the following case handles the abbreviated rules for the example
                next_gen[pot_idx] = '.'
            else:
                next_gen[pot_idx] = next_gen_plant
            # print(''.join(pot_string))
        plant_states = next_gen.copy()
        print('gen {}, sum {}, idx {}: {}'.format(str(generation).rjust(4),
                                                  str(get_pot_sum(plant_states, start_offset)).rjust(5),
                                                  str(start_offset).rjust(2),
                                                  ''.join(plant_states)
                                                  )
              )
    # once the plants have reached the final orientation we can calculate the result for the 50000000000th generation
    print(get_sum_by_gen(plant_states, start_offset, 50000000000, 100))


if __name__ == "__main__":
    main()
