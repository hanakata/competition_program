#include <stdio.h>

int main(void){
  int i,n,m,s;
  i = 1;
  scanf("%d",&n);
  while (i <= n ){
    m += i * 10000;
    i++;
  }
  s = m / n;
  printf("%d?n", s);
}