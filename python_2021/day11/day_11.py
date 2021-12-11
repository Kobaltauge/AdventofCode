import numpy as np
import scipy.signal

def get_adjacent_cells(self, cell):
        """Returns adjacent cells to a cell.

        Clockwise starting from the one on the right.

        @param cell get adjacent cells for this cell
        @returns adjacent cells list.
        """
        cells = []
        if cell.x < self.grid_width - 1:
            cells.append(self.get_cell(cell.x + 1, cell.y))
        if cell.y > 0:
            cells.append(self.get_cell(cell.x, cell.y - 1))
        if cell.x > 0:
            cells.append(self.get_cell(cell.x - 1, cell.y))
        if cell.y < self.grid_height - 1:
            cells.append(self.get_cell(cell.x, cell.y + 1))
        return cells 

with open("day_11_test.txt") as f:
    input = f.read().splitlines()
    octopus = [[int(i) for i in line] for line in input]
    octopus = np.array(octopus)

# list(map(int, (binnr)))
adder = np.full([10,10], 1, dtype='int32')
surround = np.full([3,3],1)
for step in range(0,100):
    octopus = octopus + adder
    flasher = np.argwhere(octopus==9)
    print(octopus)
    print(flasher)
    adjcells = get_adjacent_cells(octopus, cell)

    print("x")
print(f"result part 1: {map}")


