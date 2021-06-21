#include <stdio.h>

int main(void){
  char a[6];
  int b,p,q,s;
  scanf("%s",a);
  scanf("%d",&b);
  p = (b - 1) / 5;
  q = b % 5;
  if( q == 0){
    s = 4;
  }else{
    s = q - 1;
  }
  printf("%c%c\n",a[p],a[s]);
}
