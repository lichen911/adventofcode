import sys

input_file = 'day8_input.txt'


def load_input():
    instruction_list = []
    with open(input_file) as file:
        for line in file:
            instruction = {
                'command': line.strip().split(' ')[0],
                'argument': int(line.strip().split(' ')[1]),
                'run_count': 0
            }

            instruction_list.append(instruction)
    return instruction_list


def zero_run_counts(instruction_list):
    for instruction in instruction_list:
        instruction['run_count'] = 0


def run_program(instruction_list):
    current_location = 0
    accumulator = 0
    exited_normally = True
    while True:
        if current_location >= len(instruction_list):
            # print('Reached end of program.')
            break

        current_instruction = instruction_list[current_location]
        # print(f'Current location: {current_location}, Current instruction: {current_instruction}')

        if current_instruction['run_count'] >= 1:
            # print('Loop detected. Exiting program.')
            exited_normally = False
            break

        if current_instruction['command'] == 'nop':
            current_instruction['run_count'] += 1
            current_location += 1
            continue
        elif current_instruction['command'] == 'acc':
            accumulator += current_instruction['argument']
            current_instruction['run_count'] += 1
            current_location += 1
        elif current_instruction['command'] == 'jmp':
            current_location += current_instruction['argument']
            current_instruction['run_count'] += 1
        else:
            print('Unknown instruction. Halting.')
            sys.exit(1)
    return accumulator, exited_normally


def main():
    instruction_list = load_input()

    # print(instruction_list)

    accumulator, exited_normally = run_program(instruction_list)
    print(f'Accumulator value: {accumulator}, Exited normally: {exited_normally}')

    for instruction in instruction_list[::-1]:
        zero_run_counts(instruction_list)

        if instruction['command'] == 'nop':
            # print('Change nop to jmp')
            instruction['command'] = 'jmp'
        elif instruction['command'] == 'jmp':
            # print('Change jmp to nop')
            instruction['command'] = 'nop'

        accumulator, exited_normally = run_program(instruction_list)
        if exited_normally:
            print(f'Accumulator value: {accumulator}, Exited normally: {exited_normally}')
            break
        else:
            if instruction['command'] == 'nop':
                instruction['command'] = 'jmp'
            elif instruction['command'] == 'jmp':
                instruction['command'] = 'nop'


if __name__ == '__main__':
    main()
