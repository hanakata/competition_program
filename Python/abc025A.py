n = list(input())
m = int(input())
str_list = []

for i in range(5):
    for j in range(5):
        str_list.append(n[i] + n[j])
str_list.sort()
print(str_list[m - 1])
