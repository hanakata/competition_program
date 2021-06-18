#include <stdio.h>

int main(void){
  int n,m;
  scanf("%d",&n);
  scanf("%d",&m);
  m = n % m;
  if( m == 0){
    printf("YES\n");
  }else{
    printf("NO\n");
  }
}