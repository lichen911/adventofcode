input_file = 'day8_input.txt'
# input_file = 'day8_test_input.txt'


class Node(object):
    def __init__(self):
        self.child_nodes = []
        self.metadata_entries = []


class Tree(object):
    def __init__(self, serial_tree):
        self.serial_tree = serial_tree
        self.root_node = Node()
        self.all_metadata_entries = []

    def deserialize_tree(self):
        serial_tree = self.serial_tree.copy()
        serial_tree.reverse()

        def deserializer():
            child_node_count = int(serial_tree.pop())
            metadata_count = int(serial_tree.pop())
            current_node = Node()

            for child in range(child_node_count):
                current_node.child_nodes.append(deserializer())

            for metadatum in range(metadata_count):
                new_metadatum = int(serial_tree.pop())
                current_node.metadata_entries.append(new_metadatum)
                self.all_metadata_entries.append(new_metadatum)

            return current_node
        self.root_node = deserializer()

    def get_root_node_value(self):
        def get_node_value(node):
            node_total = 0

            if len(node.child_nodes) == 0:
                return sum(node.metadata_entries)
            else:
                for metadatum in node.metadata_entries:
                    if metadatum <= len(node.child_nodes):
                        node_total += get_node_value(node.child_nodes[metadatum-1])
                return node_total
        return get_node_value(self.root_node)


def main():
    with open(input_file, 'r') as fd:
        serialized = fd.read().split()

    the_tree = Tree(serialized)
    the_tree.deserialize_tree()

    # print(the_tree.all_metadata_entries)
    print(sum(the_tree.all_metadata_entries))
    print(the_tree.get_root_node_value())


if __name__ == "__main__":
    main()
