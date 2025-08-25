import os
import random
import zipfile


def contains_nearby_duplicate(nums, k):
    last = {}
    for i, num in enumerate(nums):
        if num in last and i - last[num] <= k:
            return 1
        last[num] = i
    return 0


# Tạo thư mục testcases
os.makedirs("testcases/input", exist_ok=True)
os.makedirs("testcases/output", exist_ok=True)

cases = []

# 3 example đầu
cases.append(([1, 2, 3, 1], 3))
cases.append(([1, 0, 1, 1], 1))
cases.append(([1, 2, 3, 1, 2, 3], 2))

# Thêm 47 random case
for _ in range(47):
    n = random.randint(1, 100)  # có thể tăng n lên lớn nếu muốn
    nums = [random.randint(-(10**9), 10**9) for _ in range(n)]
    k = random.randint(0, 100)
    cases.append((nums, k))

# Ghi file input/output
for idx, (nums, k) in enumerate(cases):
    fname_in = f"testcases/input/input{idx:02}.txt"
    fname_out = f"testcases/output/output{idx:02}.txt"

    with open(fname_in, "w") as f:
        f.write(str(len(nums)) + "\n")
        f.write(" ".join(map(str, nums)) + "\n")
        f.write(str(k) + "\n")

    expected = contains_nearby_duplicate(nums, k)
    with open(fname_out, "w") as f:
        f.write(str(expected) + "\n")

# Đóng gói zip
with zipfile.ZipFile("testcases.zip", "w") as z:
    for folder in ["input", "output"]:
        for fname in os.listdir(f"testcases/{folder}"):
            z.write(f"testcases/{folder}/{fname}", f"{folder}/{fname}")

print(f"✅ Generated testcases.zip with {len(cases)} cases (3 examples + 47 random)")
