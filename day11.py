class FuelGrid(object):
    def __init__(self, height, width):
        self.height = height
        self.width = width
        self.fuel_grid = [['.' for _ in range(self.width)] for _ in range(self.height)]

    def __str__(self):
        printed_grid = ''
        for y in self.fuel_grid:
            printed_grid += ','.join(str(i) for i in y) + '\n'
        return printed_grid

    def update_cell(self, x, y, value):
        self.fuel_grid[y-1][x-1] = value

    def get_cell(self, x, y):
        return self.fuel_grid[y-1][x-1]

    def compute_power_level(self, serial_number):
        for x in range(1, self.width + 1):
            for y in range(1, self.height + 1):
                rack_id = x + 10
                power_level = rack_id * y
                power_level += serial_number
                power_level *= rack_id
                power_level = int((power_level/100) % 10)
                power_level -= 5
                self.update_cell(x, y, power_level)

    def compute_region(self, x, y, size):
        region_total = 0
        x, y = x-1, y-1

        for i in range(y, y+size):
            try:
                region_total += sum(self.fuel_grid[i][x:x+size])
            except IndexError:
                return None
        return region_total

    def find_greatest_region(self, region_size):
        max_region_coord = ()
        max_region_value = -999

        for x in range(1, self.width + 1):
            for y in range(1, self.height + 1):
                current_region_value = self.compute_region(x, y, region_size)
                if current_region_value:
                    if current_region_value > max_region_value:
                        max_region_value = current_region_value
                        max_region_coord = (x, y)
        return max_region_coord


def main():
    fuel_levels = FuelGrid(300, 300)

    fuel_levels.compute_power_level(serial_number=3999)
    # print(fuel_levels)
    # print(fuel_levels.get_cell(33, 45))
    # print(fuel_levels.compute_region(33, 45, 3))
    print(fuel_levels.find_greatest_region(3))

    max_region_value = 0
    max_region_coord = ()
    max_region_size = 0
    for i in range(1, 301):
        current_max_region_coord = fuel_levels.find_greatest_region(i)
        current_max_region_value = fuel_levels.compute_region(*current_max_region_coord, i)
        if current_max_region_value > max_region_value:
            max_region_value = current_max_region_value
            max_region_coord = current_max_region_coord
            max_region_size = i
    print(max_region_coord, max_region_size)


if __name__ == "__main__":
    main()
