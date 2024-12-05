def is_safe_with_dampener(levels):
    def check_safety(levels):
        is_increasing = all(
            1 <= levels[i + 1] - levels[i] <= 3 for i in range(len(levels) - 1)
        )
        is_decreasing = all(
            1 <= levels[i] - levels[i + 1] <= 3 for i in range(len(levels) - 1)
        )
        return is_increasing or is_decreasing

    if check_safety(levels):
        return True

    for i in range(len(levels)):
        modified_levels = levels[:i] + levels[i + 1:]
        if check_safety(modified_levels):
            return True

    return False


with open('input.txt', 'r', encoding='utf-8') as f:
    lines = f.read().splitlines()

safe_count = 0
for line in lines:
    levels = list(map(int, line.split()))
    if is_safe_with_dampener(levels):
        safe_count += 1

print(f"Number of safe reports: {safe_count}")
