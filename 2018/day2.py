<<<<<<< HEAD
from collections import Counter

input_file = 'day2_input.txt'
box_counts = {
    'twice': 0,
    'thrice': 0
}


def char_appears_twice(box_code):
    char_counts = Counter(box_code)
    for value in char_counts.values():
        if value == 2:
            return True
    return False


def char_appears_thrice(box_code):
    char_counts = Counter(box_code)
    for value in char_counts.values():
        if value == 3:
            return True
    return False

with open(input_file, 'r') as fd:
    box_list = fd.read().splitlines()

for box in box_list:
    if char_appears_twice(box):
        box_counts['twice'] += 1
    if char_appears_thrice(box):
        box_counts['thrice'] += 1

print(box_counts)
print('checksum:', box_counts['twice'] * box_counts['thrice'])

print(box_list)
box_list.sort()
print(box_list)

for i in range(1, len(box_list)):
    box1 = box_list[i-1]
    box2 = box_list[i]
    mismatch_count = 0
    for x, y in zip(box1, box2):
        if x != y:
            mismatch_count += 1
        if mismatch_count > 1:
            break
    if mismatch_count == 1:
        print(box1)
        print(box2)
=======
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
>>>>>>> ad2019
