#include <stdio.h>

int main()
{
    char str[] = "ABCDE";
    char s[2];
    int n;
    scanf("%s",s);
    n = 0;
    while (str[n] != s[0]) {
        n++;
    }
    printf("%d\n",n + 1);
}
