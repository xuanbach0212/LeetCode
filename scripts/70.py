import os
import random
import zipfile


# Hàm tính số cách leo thang
def climbing_stairs(n):
    if n == 1:
        return 1
    if n == 2:
        return 2
    a, b = 1, 2
    for _ in range(3, n + 1):
        a, b = b, a + b
    return b


# Tạo thư mục testcases
os.makedirs("testcases/input", exist_ok=True)
os.makedirs("testcases/output", exist_ok=True)

cases = []

# 2 ví dụ gốc
cases.append(2)  # Example 1
cases.append(3)  # Example 2

# 50 test case random
for _ in range(50):
    n = random.randint(1, 45)
    cases.append(n)

# Ghi file input/output
for idx, n in enumerate(cases):
    fname_in = f"testcases/input/input{idx:02}.txt"
    fname_out = f"testcases/output/output{idx:02}.txt"

    with open(fname_in, "w") as f:
        f.write(str(n) + "\n")

    expected = climbing_stairs(n)
    with open(fname_out, "w") as f:
        f.write(str(expected) + "\n")

# Đóng gói zip
with zipfile.ZipFile("testcases.zip", "w") as z:
    for folder in ["input", "output"]:
        for fname in os.listdir(f"testcases/{folder}"):
            z.write(f"testcases/{folder}/{fname}", f"{folder}/{fname}")

print(f"✅ Generated testcases.zip with {len(cases)} cases (2 examples + 50 random)")
