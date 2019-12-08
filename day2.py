def get_intcode(input_file):
    intcode = []
    with open(input_file, 'r') as file:
        intcode = file.readline().split(',')

    intcode = [int(x) for x in intcode]
    return intcode


def run_intcode(intcode):
    for index in range(0, len(intcode), 4):
        # print('intcode line: {}'.format(intcode[index:index+4]))
        param1 = intcode[index + 1]
        param2 = intcode[index + 2]
        result_address = intcode[index + 3]

        opcode = intcode[index]
        if opcode == 99:
            break
        elif opcode == 1:
            opcode_result = intcode[param1] + intcode[param2]
        elif opcode == 2:
            opcode_result = intcode[param1] * intcode[param2]
        else:
            print('Unknown opcode.')
            break

        intcode[result_address] = opcode_result

def main():
    input_file = 'day2_input.txt'
    # input_file = 'day2_input_test.txt'
    intcode = get_intcode(input_file)

    noun = 0
    verb = 0
    for noun in range(0,100):
        for verb in range(0,100):
            intcode_copy = intcode[::]
            intcode_copy[1] = noun
            intcode_copy[2] = verb
            run_intcode(intcode_copy)
            if intcode_copy[0] == 19690720:
                break
        else:
            continue
        break
    print(intcode_copy)
    print(100 * noun + verb)


if __name__ == '__main__':
    main()
