#include <stdio.h>

int main(void){
  int a,b,c,w,i,m,j;
  j = 0;
  scanf("%d %d %d",&a ,&b ,&c);
  scanf("%d",&w);
  if( b <= w && w <= c){
    j++;
  }
  for(i = 0;i < a - 1; i++){
    scanf("%d",&m);
    w = w + m;
    if( b <= w && w <= c){
      j++;
    }
  }
  printf("%d\n", j);
}
