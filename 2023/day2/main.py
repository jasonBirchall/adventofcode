def part1(input_file: str) -> int:
    max_red, max_green, max_blue = 12, 13, 14
    total = 0

    with open (input_file, 'r') as f:
        lines = f.readlines()
        for line in lines:
        
            game_id = int(line.split(':')[0].split(' ')[1])
            rounds = line.split(':')[1].split(';')

            possible = True

            print("Game: " + str(game_id))
            print("Rounds: " + str(rounds))

            for round in rounds:
                red = 0
                green = 0
                blue = 0
                cubes = round.strip().split(',')

                for cube in cubes:
                    number, color = cube.strip().split(' ')
                    number = int(number)                     

                    if color == 'red':
                        red += number
                    elif color == 'green':
                        green += number
                    elif color == 'blue':
                        blue += number

                    if red > max_red or green > max_green or blue > max_blue:
                        possible = False
                        break

                if not possible:
                    break
            
            if possible:
                total += game_id

        return total


def part2(input_file: str) -> int:
    total_power = 0

    with open(input_file, 'r') as f:
        for line in f:
            rounds = line.split(': ')[1].split(';')
            min_cubes = {'red': 0, 'green': 0, 'blue': 0}

            for round in rounds:
                for cube in round.strip().split(','):
                    number, color = cube.strip().split(' ')
                    min_cubes[color] = max(min_cubes[color], int(number))

            power = min_cubes['red'] * min_cubes['green'] * min_cubes['blue']
            total_power += power

    return total_power


if __name__ == '__main__':
    print(part1('input.txt'))
    print(part2('input.txt'))

