def print_recipe_list(recipe_list, elf1, elf2):
    recipe_string = ''
    for i, recipe in enumerate(recipe_list):
        if i == elf1:
            recipe = '({})'.format(recipe_list[i])
        elif i == elf2:
            recipe = '[{}]'.format(recipe_list[i])
        else:
            recipe = ' {} '.format(recipe_list[i])
        recipe_string += str(recipe)
    print(recipe_string)


def main():
    puzzle_input = '704321'

    recipe_list = [3, 7]
    elf1 = 0
    elf2 = 1
    # print_recipe_list(recipe_list, elf1, elf2)

    while True:
        current_recipe_sum = recipe_list[elf1] + recipe_list[elf2]
        new_recipe_list = [int(x) for x in list(str(current_recipe_sum))]
        recipe_list += new_recipe_list

        elf1 = (elf1 + recipe_list[elf1] + 1) % len(recipe_list)
        elf2 = (elf2 + recipe_list[elf2] + 1) % len(recipe_list)

        # print_recipe_list(recipe_list, elf1, elf2)
        # print(recipe_list)

        recipe_string = ''.join([str(x) for x in recipe_list[len(recipe_list)-(len(new_recipe_list)+len(puzzle_input)):]])
        if puzzle_input in recipe_string:
            break

    print('previous recipes: {}'.format(''.join([str(x) for x in recipe_list]).find(puzzle_input)))


if __name__ == "__main__":
    main()
