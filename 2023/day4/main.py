def part1(filename: str) -> int:
    with open(filename) as f:
        total_score = 0
        for line in f:
            line = line.strip().split(":")[1].split("|")
            winning = set(line[0].strip().split())
            my_numbers = set(line[1].strip().split())
            matches = winning & my_numbers

            score = 0
            for index in range(len(matches)):
                score = 1 << index

            total_score += score

    return total_score

if __name__ == "__main__":
    print(part1("input.txt"))

