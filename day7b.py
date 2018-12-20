import string

# input_file = 'day7_input.txt'
input_file = 'day7_test_input.txt'


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
    # step_lookup = {letter: {'pre': set()} for letter in 'ABCDEF'}

    with open(input_file, 'r') as fd:
        step_list = fd.read().splitlines()

    for step in step_list:
        step_tokens = step.split()
        step_pre = step_tokens[1]
        step_letter = step_tokens[7]
        step_lookup[step_letter.upper()]['pre'].add(step_pre)

    step_list = []
    # worker_list = [{'step': None, 'time_remaining': None} for _ in range(5)]
    worker_list = [{'step': None, 'time_remaining': None} for _ in range(2)]
    # print(worker_list)

    # generate time lookup dict
    # base_time = 60
    base_time = 0
    time_lookup = {letter: base_time for letter in string.ascii_uppercase}
    for letter in time_lookup:
        time_lookup[letter] += string.ascii_uppercase.index(letter)+1

    second_counter = 0
    current_workers = 0
    while True:
        for worker in worker_list:
            if worker['time_remaining'] == 1:
                current_step = worker['step']
                step_list.append(current_step)
                remove_pre_req(step_lookup, current_step)
                worker['step'] = None
                worker['time_remaining'] = None
                current_workers -= 1
            elif worker['time_remaining']:
                worker['time_remaining'] -= 1

            if not worker['step']:
                new_job = find_first_with_no_pre_req(step_lookup)
                if new_job:
                    worker['step'] = new_job
                    worker['time_remaining'] = time_lookup[new_job]
                    step_lookup.pop(new_job, None)
                    current_workers += 1

        print('%s %s %s' % (second_counter, worker_list, ''.join(step_list)))
        if not find_first_with_no_pre_req(step_lookup) and current_workers == 0:
            break
        second_counter += 1


if __name__ == "__main__":
    main()
