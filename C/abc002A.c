#include <stdio.h>

int main(void){
  int n,m;
  scanf("%d",&n);
  scanf("%d",&m);
  if( n < m){
    printf("%d\n", m);
  }else{
    printf("%d\n", n);
  }
}