#include <stdio.h>

int main(void){
  int a,b,c,k,s,t,n,m,o;
  scanf("%d %d %d %d",&a ,&b ,&c ,&k);
  scanf("%d %d",&s ,&t);
  n = s * a;
  m = t * b;
  o = s + t;
  if( o >= k){
    printf("%d\n", n + m - (o * c));
  }else{
    printf("%d\n", n + m);
  }
}
