#include <stdio.h>

int main(void){
  int n,i,j;
  char l[4];
  int c[4] = {8,4,2,1};
  scanf("%d",&n);
  while (n != 0) {
    if(c[j] <= n){
      n -= c[j];
      l[i] = c[j];
      i++;
    }
    j++;
  }
  printf("%d\n",i);
  n = 0;
  while (n < i ) {
    printf("%d\n",l[n]);
    n++;
  }
}