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