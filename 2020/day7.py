input_file = 'day7_input.txt'


def load_rules():
    rules = {}

    with open(input_file) as file:
        for line in file:
            line = line.strip()
            outer_bag = ' '.join(line.split()[0:2])
            rules[outer_bag] = {}

            contains = ' '.join(line.split(' ')[4:]).split(',')
            for item in contains:
                item = item.strip()
                quantity = item.split(' ')[0]
                if quantity != 'no':
                    quantity = int(quantity)
                    inner_bag = ' '.join(item.split(' ')[1:3])
                    rules[outer_bag][inner_bag] = quantity
    return rules


def bag_search(bag_types, rules):
    outer_bags = []
    for outer_bag in rules:
        for inner_bag in rules[outer_bag]:
            if inner_bag in bag_types:
                outer_bags.append(outer_bag)
    return outer_bags


def find_contents(input_bag_type, input_rules):
    def _count_bags(bag_type, rules, bag_total):
        if not rules[bag_type]:
            return None
        else:
            bag_subtotal = 0
            for bag in rules[bag_type]:
                count_result = _count_bags(bag, rules, bag_total)
                if not count_result:
                    bag_subtotal = bag_subtotal + rules[bag_type][bag]
                else:
                    bag_subtotal = bag_subtotal + (count_result * rules[bag_type][bag])
            return bag_total + bag_subtotal + 1

    return _count_bags(input_bag_type, input_rules, 0)


def main():
    rules = load_rules()
    print(rules)

    new_outer_bags = bag_search(['shiny gold'], rules)
    outer_bags = new_outer_bags
    while True:
        new_outer_bags = bag_search(new_outer_bags, rules)
        if new_outer_bags:
            outer_bags = outer_bags + new_outer_bags
        else:
            break
    print(len(set(outer_bags)))

    bag_count = find_contents('shiny gold', rules)
    print(f'bag count: {bag_count - 1}')


if __name__ == '__main__':
    main()
