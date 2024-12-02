with open('test.txt', 'r', -1, 'utf') as f:
    lines = f.read().splitlines()

left = []
right = []
for line in lines:
    l, r = line.split()
    left.append(int(l))
    right.append(int(r))
left.sort()
right.sort()

sum = 0
LENGTH = len(left)
for i in range(LENGTH):
    sum += abs(left[i] - right[i])
print(sum)
