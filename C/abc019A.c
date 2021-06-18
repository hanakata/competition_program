#include <stdio.h>

int main(void){
  int a,b,c, i, j, t;
  scanf("%d %d %d", &a, &b, &c);
  int n[ ] = {a, b, c};
  for(i = 0;i < 2; i++){
    for(j = 2; j > i; j--){
      if(n[j] < n[j-1]){
      t = n[j-1];
      n[j-1] = n[j];
      n[j] = t;
      }
    }
  }
  printf("%d\n", n[1]);
}
