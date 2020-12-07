input_file = 'day2_input.txt'


def check_pass1(policy_letter, min, max, password):
    count = 0
    for letter in password:
        if letter == policy_letter:
            count = count + 1

    if min <= count <= max:
        return True
    else:
        return False


def check_pass2(policy_letter, idx1, idx2, password):
    first_letter = password[idx1-1]
    second_letter = password[idx2-1]

    if (first_letter == policy_letter) is not (second_letter == policy_letter):
        return True
    else:
        return False


def main():
    with open(input_file) as file:
        check1_count = 0
        check2_count = 0

        for line in file:
            line = line.strip()
            line = line.split(':')
            letter_min = int(line[0].split(' ')[0].split('-')[0])
            letter_max = int(line[0].split(' ')[0].split('-')[1])
            letter = line[0].split(' ')[1]
            password = line[1].strip()

            if check_pass1(letter, letter_min, letter_max, password):
                check1_count = check1_count + 1

            if check_pass2(letter, letter_min, letter_max, password):
                check2_count = check2_count + 1

    print(check1_count, check2_count)


if __name__ == '__main__':
    main()
