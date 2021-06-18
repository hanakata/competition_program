#include <stdio.h>

int main(void){
  static char line[64];
  int a, b, n, m;
  m = 0;
  while (fgets(line, sizeof(line), stdin) != NULL) {
    sscanf(line, "%d %d", &a, &b);
    n = a * b / 10;
    m += n;
    }
    printf("%d\n",m);
}
