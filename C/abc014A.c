#include <stdio.h>

int main(void){
  int n,m,i;
  scanf("%d",&n);
  scanf("%d",&m);
  i = n % m;
  if( i == 0 ){
    printf("%d\n",0);
  }else{
    printf("%d\n",m - i);
  }
}
