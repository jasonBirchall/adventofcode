def part1(file_path: str) -> int:
    total = []
    with open(file_path) as f:
        for line in f:
            line = line.strip()
            digits = [int(char) for char in line if char.isdigit()]

            if len(digits) > 0:
                concatenated = str(digits[0]) + str(digits[-1])
            elif len(digits) == 1:
                concatenated = str(digits[0]) * 2
            else:
                print("skipping", line)
                continue

            total.append(int(concatenated))
            print("Line: ", line, "Concatenated: ", concatenated)


    return sum(total)

def extract_digit(substring, mapping):
    if substring in mapping:
        return mapping[substring]
    elif len(substring) == 1 and substring.isdigit():
        return substring
    return None

def find_first_digit(line, mapping):
    for i in range(len(line)):
        for j in range(i + 1, len(line) + 1):
            digit = extract_digit(line[i:j], mapping)
            if digit:
                return digit
    return None

def find_last_digit(line, mapping):
    for i in range(len(line), 0, -1):
        for j in range(0, i):
            digit = extract_digit(line[j:i], mapping)
            if digit:
                return digit
    return None

def part2(file_path):
    total = 0
    mapping = {
        "one": "1", "two": "2", "three": "3", "four": "4",
        "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"
    }

    with open(file_path) as f:
        for line in f:
            line = line.strip()

            print("Line: ", line)
            first_digit = find_first_digit(line, mapping)
            last_digit = find_last_digit(line, mapping)

            if first_digit and last_digit:
                total += int(first_digit + last_digit)
                print("Concatenated: ", first_digit + last_digit)

    return total



if __name__ == "__main__":
    print(part1("input.txt"))
    print(part2("input.txt"))
