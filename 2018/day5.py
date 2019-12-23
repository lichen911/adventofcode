import string

# input_file = 'day5_input_test.txt'
input_file = 'day5_input.txt'

with open(input_file, 'r') as fd:
    input_string = fd.read()


def react(polymer_chain):
    post_reaction = polymer_chain

    for i in range(len(polymer_chain)-1):
        if (polymer_chain[i+1].isupper() and polymer_chain[i] == polymer_chain[i+1].lower()) or\
                (polymer_chain[i+1].islower() and polymer_chain[i] == polymer_chain[i+1].upper()):
            # print(polymer_chain[i], polymer_chain[i+1].lower())
            # print(polymer_chain[i], polymer_chain[i + 1].upper())
            # print('i:', i)
            # print(polymer_chain[:i], polymer_chain[i+2:])
            post_reaction = polymer_chain[:i] + polymer_chain[i+2:]
            return True, post_reaction

    return False, post_reaction


def remove_unit_and_react(unit, polymer_chain):
    modified_polymer = polymer_chain.replace(unit.upper(), '').replace(unit.lower(), '')

    reaction_occurred = True
    reaction_result = modified_polymer.strip()
    while reaction_occurred:
        reaction_occurred, reaction_result = react(reaction_result)
        # print(reaction_occurred, reaction_result)
        # print('polymer chain length: ', len(reaction_result))

    return reaction_result


# Part A
# print(input_string)
# reaction_occurred = True
# reaction_result = input_string.strip()
# while reaction_occurred:
#     reaction_occurred, reaction_result = react(reaction_result)
#     # print(reaction_occurred, reaction_result)
#     print('polymer chain length: ', len(reaction_result))
#
# print(reaction_result)

# Part B
minimum_count = len(input_string)
reaction_counts = {letter: None for letter in string.ascii_lowercase}
for letter in reaction_counts:
    reaction_counts[letter] = len(remove_unit_and_react(letter, input_string))
    if reaction_counts[letter] < minimum_count:
        minimum_count = reaction_counts[letter]
    print(letter, reaction_counts[letter])

print(minimum_count)
