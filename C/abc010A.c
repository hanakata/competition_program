#include <stdio.h>

int main(void)
{
  char str[32];
  scanf("%s",str);
  sprintf(str,"%s%s\n",str,"pp");
  printf(str);
}