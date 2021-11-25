#include <stdio.h>

int main(void){
  int a;
  scanf("%d",&a);
  if(a <= 59){
    printf("%s\n", "Bad");
    return;
  }
  if(60 <= a && a <= 89){
    printf("%s\n", "Good");
    return;
  }
  if(90 <= a && a<= 99){
    printf("%s\n", "Great");
    return;
  }
  if(100 == a){
    printf("%s\n", "Perfect");
    return;
  }
}