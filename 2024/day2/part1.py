with open('input.txt', 'r', -1, 'utf') as f:
    lines = f.read().splitlines()

number_of_safe_reports = len(lines)
print(number_of_safe_reports)
for line in lines:
    print(line)
    last_number = int(line.split()[0])
    IS_GREATER = last_number < int(line.split()[1])
    split_line = line.split()
    for num in split_line[1:]:
        num = int(num)
        if num == last_number:
            print('unsafe')
            print('same number', num, last_number)
            number_of_safe_reports -= 1
            break
        if IS_GREATER:
            if num < last_number:
                print('unsafe')
                print('less than')
                number_of_safe_reports -= 1
                break

            if num - last_number > 3:
                print('unsafe')
                print('too many steps')
                number_of_safe_reports -= 1
                break
        else:
            if num > last_number:
                print('unsafe')
                print('greater than')
                number_of_safe_reports -= 1
                break

            if last_number - num > 3:
                print('unsafe')
                print('too many steps')
                number_of_safe_reports -= 1
                break

        last_number = num

print(number_of_safe_reports)
