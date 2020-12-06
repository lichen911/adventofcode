input_file = 'day4_input.txt'


def load_input():
    passport_list = []
    passport = {}

    with open(input_file) as file:
        for line in file:
            if line == '\n':
                passport_list.append(passport)
                passport = {}
            else:
                key_value_pairs = line.strip().split(' ')
                for kv_pair in key_value_pairs:
                    passport[kv_pair.split(':')[0]] = kv_pair.split(':')[1]
        passport_list.append(passport)
    return passport_list


def validate_passport_key(passport, key):
    valid = False
    if key == 'byr':
        if 1920 <= int(passport[key]) <= 2002:
            valid = True
    elif key == 'iyr':
        if 2010 <= int(passport[key]) <= 2020:
            valid = True
    elif key == 'eyr':
        if 2020 <= int(passport[key]) <= 2030:
            valid = True
    elif key == 'hgt':
        try:
            hgt_value = int(passport[key][:-2])
        except ValueError:
            pass
        else:
            hgt_unit = passport[key][-2:]
            if hgt_unit == 'cm':
                if 150 <= hgt_value <= 193:
                    valid = True
            if hgt_unit == 'in':
                if 59 <= hgt_value <= 76:
                    valid = True
    elif key == 'hcl':
        if passport[key][0] == '#':
            try:
                int(passport[key][1:], 16)
            except ValueError:
                pass
            else:
                valid = True
    elif key == 'ecl':
        if passport[key] in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
            valid = True
    elif key == 'pid':
        try:
            int(passport[key])
        except ValueError:
            pass
        else:
            if len(passport[key]) == 9:
                valid = True

    return valid


def get_valid_passports(passport_list):
    required_keys = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
    valid_passports = 0

    for passport in passport_list:
        current_valid = True
        for key in required_keys:
            if key not in passport:
                current_valid = False
                break
            else:
                if current_valid:
                    current_valid = validate_passport_key(passport, key)
                else:
                    break
        if current_valid:
            print(f'valid passport: {passport}')
            valid_passports = valid_passports + 1
        else:
            print(f'invalid passport: {passport}')
    return valid_passports


def main():
    passport_list = load_input()
    print(f'passport list: {passport_list}')

    valid_passports = get_valid_passports(passport_list)
    print(f'valid passports: {valid_passports}')


if __name__ == '__main__':
    main()
