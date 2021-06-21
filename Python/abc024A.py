m = input().split()
n = input().split()
p = int(n[0]) * int(m[0]) + int(n[1]) * int(m[1])
t = int(n[0]) + int(n[1])
if( t >= int(m[3]) ):
    print( p - int(m[2]) * t )
else:
    print( p )
