def create_matrix(input_path) -> list:
    with open(input_path, 'r') as f:
        lines = f.readlines()
        matrix = [list(line.strip()) for line in lines]
    return matrix

def part1(matrix: list) -> int:
    rows, cols = len(matrix), len(matrix[0])
    total_sum = 0

    for x in range(rows):
        y = 0
        while y < cols:
            if is_num(matrix[x][y]):
                number = matrix[x][y]
                length = 1
                while y + length < cols and is_num(matrix[x][y + length]):
                    number += matrix[x][y + length]
                    length += 1

                if adjacent_to_symbol(matrix, x, y, length, rows, cols):
                    total_sum += int(number)

                y += length
            else:
                y += 1

    return total_sum

def is_num(char) -> bool:
    return char.isdigit()

def adjacent_to_symbol(matrix, x, y, length, rows, cols) -> bool:
    directions = [(dx, dy) for dx in [-1, 0, 1] 
        for dy in [-1, 0, 1] if not (dx == 0 and dy == 0)]

    for i in range(length):
        for dx, dy in directions:
            nx, ny = x + dx, y + dy + i
            if 0 <= nx < rows and 0 <= ny < cols and is_symbol(matrix[nx][ny]):
                return True
    return False

def is_symbol(char: str) -> bool:
    return char not in \
    ['.', ' ', '\n', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9']

if __name__ == "__main__":
    matrix = create_matrix("input.txt")
    print(part1(matrix))

