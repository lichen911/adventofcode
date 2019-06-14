class Cart(object):
    def __init__(self, cart_icon, track_segment):
        self.cart_icon = cart_icon
        self.track_segment = track_segment
        self.last_turn = 'right'
        self.tick = 0

        self.intersect_lookup = {
            '<': {'left': 'v', 'straight': '<', 'right': '^'},
            '^': {'left': '<', 'straight': '^', 'right': '>'},
            '>': {'left': '^', 'straight': '>', 'right': 'v'},
            'v': {'left': '>', 'straight': 'v', 'right': '<'}
        }

        self.turn_lookup = {
            '/': {'^': '>', '<': 'v', '>': '^', 'v': '<'},
            '\\': {'v': '>', '<': '^', '>': 'v', '^': '<'}
        }

    def use_intersect(self):
        if self.last_turn == 'left':
            self.cart_icon = self.intersect_lookup[self.cart_icon]['straight']
            self.last_turn = 'straight'
        elif self.last_turn == 'straight':
            self.cart_icon = self.intersect_lookup[self.cart_icon]['right']
            self.last_turn = 'right'
        elif self.last_turn == 'right':
            self.cart_icon = self.intersect_lookup[self.cart_icon]['left']
            self.last_turn = 'left'

    def make_turn(self):
        self.cart_icon = self.turn_lookup[self.track_segment][self.cart_icon]


def load_input_file():
    input_filename = 'day13_input.txt'
    # input_filename = 'day13b_input_test.txt'
    track_grid = []

    with open(input_filename, 'r') as fd:
        for row in fd:
            track_row = list(row.rstrip('\n'))
            for i, track_segment in enumerate(track_row):
                if track_segment in ['<', '>']:
                    track_row[i] = Cart(track_segment, '-')
                elif track_segment.lower() in ['^', 'v']:
                    track_row[i] = Cart(track_segment, '|')
            track_grid.append(track_row)
    return track_grid


def print_track(track_grid):
    printed_grid = ''
    for y in track_grid:
        new_y = y.copy()
        for i, x in enumerate(new_y):
            if isinstance(x, Cart):
                new_y[i] = x.cart_icon
        printed_grid += ''.join(str(i) for i in new_y) + '\n'
    print(printed_grid)


def get_total_carts(track_grid):
    total_carts = 0
    for y in range(0, len(track_grid)):
        for x in range(0, len(track_grid[y])):
            if isinstance(track_grid[y][x], Cart):
                total_carts += 1
    return total_carts


def get_final_cart(track_grid):
    for y in range(0, len(track_grid)):
        for x in range(0, len(track_grid[y])):
            if isinstance(track_grid[y][x], Cart):
                return x, y
    return None


def advance_carts(track_grid, tick):
    crashes_in_tick = 0
    for y in range(0, len(track_grid)):
        for x in range(0, len(track_grid[y])):
            if isinstance(track_grid[y][x], Cart) and track_grid[y][x].tick < tick:

                if track_grid[y][x].track_segment in ['/', '\\']:
                    track_grid[y][x].make_turn()
                elif track_grid[y][x].track_segment == '+':
                    track_grid[y][x].use_intersect()

                new_x = x
                new_y = y
                if track_grid[y][x].cart_icon == '<':
                    new_x -= 1
                elif track_grid[y][x].cart_icon == '>':
                    new_x += 1
                elif track_grid[y][x].cart_icon == '^':
                    new_y -= 1
                elif track_grid[y][x].cart_icon == 'v':
                    new_y += 1

                # crash detection
                if isinstance(track_grid[new_y][new_x], Cart):
                    track_grid[y][x] = track_grid[y][x].track_segment
                    track_grid[new_y][new_x] = track_grid[new_y][new_x].track_segment
                    crashes_in_tick += 1
                else:
                    dest_track_segment = track_grid[new_y][new_x]
                    track_grid[new_y][new_x] = track_grid[y][x]
                    track_grid[y][x] = track_grid[new_y][new_x].track_segment
                    track_grid[new_y][new_x].track_segment = dest_track_segment
                    track_grid[new_y][new_x].tick = tick
    return crashes_in_tick



def main():
    track_grid = load_input_file()
    # print_track(track_grid)
    tick = 0
    total_carts = get_total_carts(track_grid)
    while True:
        tick += 1
        crashes_in_tick = advance_carts(track_grid, tick)
        total_carts -= crashes_in_tick * 2
        if total_carts <= 1:
            print('Remaining cart: {}'.format(get_final_cart(track_grid)))
            break


if __name__ == "__main__":
    main()
