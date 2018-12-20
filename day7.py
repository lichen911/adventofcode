import string

input_file = 'day7_input.txt'


def find_first_with_no_pre_req(step_lookup):
    for letter in string.ascii_uppercase:
        try:
            if len(step_lookup[letter]['pre']) == 0:
                return letter
        except KeyError:
            pass
    return None


def remove_pre_req(step_lookup, pre_req):
    for letter in string.ascii_uppercase:
        try:
            if pre_req in step_lookup[letter]['pre']:
                step_lookup[letter]['pre'].remove(pre_req)
        except KeyError:
            pass


def main():
    step_lookup = {letter: {'pre': set()} for letter in string.ascii_uppercase}

    with open(input_file, 'r') as fd:
        step_list = fd.read().splitlines()

    for step in step_list:
        step_tokens = step.split()
        step_pre = step_tokens[1]
        step_letter = step_tokens[7]
        step_lookup[step_letter.upper()]['pre'].add(step_pre)

    print(step_lookup)

    step_list = []

    while True:
        current_step = find_first_with_no_pre_req(step_lookup)

        if current_step:
            step_list.append(current_step)
            remove_pre_req(step_lookup, current_step)
            step_lookup.pop(current_step, None)
        else:
            break

    print(''.join(step_list))


if __name__ == "__main__":
    main()
