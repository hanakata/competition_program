m = int(input())
memo = [0] * m
memo[0:2] = [0,0,1]
for i in range(3,m):
    memo[i] = (memo[i - 1] + memo[i - 2] + memo[i - 3]) % 10007 
print(memo[m - 1])