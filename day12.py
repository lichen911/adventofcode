def load_input_file():
    # input_file = 'day12_input.txt'
    input_file = 'day12_input_test.txt'
    plant_rules = {}

    with open(input_file, 'r') as fd:
        initial_state = list(fd.readline().split()[2].strip())

        for line in fd:
            if line.strip():
                rule_condition = line.split()[0].strip()
                rule_result = line.split()[2].strip()
                plant_rules[rule_condition] = rule_result
    return initial_state, plant_rules


def main():
    plant_states, plant_rules = load_input_file()
    plant_states = ['.', '.', '.'] + plant_states

    print('{}: {}'.format('0'.rjust(2), ''.join(plant_states)))
    for generation in range(1, 21):
        left_two = ''.join(plant_states[0:2])
        if left_two in ['##', '#.']:
            plant_states = ['.', '.'] + plant_states
        elif left_two in ['.#']:
            plant_states = ['.'] + plant_states

        right_two = ''.join(plant_states[len(plant_states)-2:])
        if right_two in ['##', '.#']:
            plant_states += ['.', '.']
        elif right_two in ['#.']:
            plant_states += ['.']

        next_gen = plant_states.copy()
        for pot_idx in range(2, len(plant_states)-2):
            pot_string = plant_states[pot_idx-2:pot_idx+3]
            try:
                # print('trying: {}'.format(pot_string))
                next_gen_plant = plant_rules[''.join(pot_string)]
            except KeyError:
                # print('no match')
                pass
            else:
                print('updating {} with {}'.format(pot_idx, next_gen_plant))
                next_gen[pot_idx] = next_gen_plant
            # print(''.join(pot_string))
        plant_states = next_gen.copy()
        print('{}: {}'.format(str(generation).rjust(2), ''.join(plant_states)))


if __name__ == "__main__":
    main()
