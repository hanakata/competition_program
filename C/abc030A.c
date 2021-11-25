#include <stdio.h>
 
int main(void)
{
    int a1, a2, b1, b2; 
    double a, b;
 
    scanf("%d%d%d%d", &a1, &a2, &b1, &b2);
    a = a2 * 1.0 / a1; 
    b = b2 * 1.0 / b1; 
 
    if(a < b){
        printf("AOKI\n");
    }
    if(a == b){
        printf("DRAW\n");
    }
    if(a > b){
        printf("TAKAHASHI\n");
    }
}