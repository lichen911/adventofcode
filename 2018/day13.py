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


def advance_carts(track_grid, tick):
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
                    raise RuntimeError('Two carts have crashed at ({},{})'.format(new_x, new_y))

                dest_track_segment = track_grid[new_y][new_x]
                track_grid[new_y][new_x] = track_grid[y][x]
                track_grid[y][x] = track_grid[new_y][new_x].track_segment
                track_grid[new_y][new_x].track_segment = dest_track_segment
                track_grid[new_y][new_x].tick = tick


def main():
    track_grid = load_input_file()
    # print_track(track_grid)
    tick = 0
    while True:
        tick += 1
        try:
            advance_carts(track_grid, tick)
            # print_track(track_grid)
        except RuntimeError as err:
            print(err)
            break


if __name__ == "__main__":
    main()
