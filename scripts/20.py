import os
import random
import zipfile


# Hàm kiểm tra valid parentheses
def is_valid(s):
    stack = []
    mapping = {")": "(", "}": "{", "]": "["}
    for c in s:
        if c in mapping.values():
            stack.append(c)
        else:  # c là ), }, ]
            if not stack or stack[-1] != mapping[c]:
                return 0
            stack.pop()
    return 1 if not stack else 0


# Tạo thư mục testcases
os.makedirs("testcases/input", exist_ok=True)
os.makedirs("testcases/output", exist_ok=True)

cases = []

# 5 ví dụ gốc
cases.extend(["()", "()[]{}", "(]", "([])", "([)]"])


# Hàm sinh string ngẫu nhiên gồm (), {}, []
def generate_random_parentheses(length):
    chars = ["(", ")", "{", "}", "[", "]"]
    return "".join(random.choice(chars) for _ in range(length))


# 50 test case random
for _ in range(50):
    length = random.randint(1, 20)  # tăng max 10000 nếu muốn edge case
    s = generate_random_parentheses(length)
    cases.append(s)

# Ghi file input/output
for idx, s in enumerate(cases):
    fname_in = f"testcases/input/input{idx:02}.txt"
    fname_out = f"testcases/output/output{idx:02}.txt"

    with open(fname_in, "w") as f:
        f.write(s + "\n")

    expected = is_valid(s)
    with open(fname_out, "w") as f:
        f.write(str(expected) + "\n")

# Đóng gói zip
with zipfile.ZipFile("testcases.zip", "w") as z:
    for folder in ["input", "output"]:
        for fname in os.listdir(f"testcases/{folder}"):
            z.write(f"testcases/{folder}/{fname}", f"{folder}/{fname}")

print(f"✅ Generated testcases.zip with {len(cases)} cases (5 examples + 50 random)")
