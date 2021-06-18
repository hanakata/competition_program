#include <stdio.h>

int main(void){
  static char line[64];
  int n[3];
  int r[3] = {1,1,1};
  int a, i, j;
  i = 0;
  while (fgets(line, sizeof(line), stdin) != NULL) {
    sscanf(line, "%d", &a);
    n[i] = a;
    i++;
  }
  for (i = 0;i < 3; i++){
    for(j = 0;j < 3; j++){
      if(n[i] < n[j]){
        r[i]++;
      }
    }
  }
  printf("%d\n", r[0]);
  printf("%d\n", r[1]);
  printf("%d\n", r[2]);
}
