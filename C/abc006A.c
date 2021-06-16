#include <stdio.h>

int main(void){
  int n,m;
  scanf("%d",&n);
  m = n % 3;
  if( m == 0){
    printf("YES\n");
  }else{
    printf("NO\n");
  }
}