import re

left = []
right = []
with open('input.txt', 'r', -1, 'utf') as f:
    for line in f:
        numbers = list(map(int, re.findall(r'\d+', line)))
        left.append(numbers[0])
        right.append(numbers[1])

left.sort()
right.sort()

sum = 0
LENGTH = len(left)
for i in range(LENGTH):
    sum += abs(left[i] - right[i])
print(sum)
