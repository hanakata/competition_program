#include <stdio.h>

int main(void){
  int a,b,c;
  scanf("%d",&a);
  scanf("%d",&b);
  scanf("%d",&c);
  if(a == b){
    printf("%d\n", c);
    return;
  }else if(a == c){
    printf("%d\n", b);
    return;
  }else if(b == c){
    printf("%d\n", a);
    return;
  }
}
