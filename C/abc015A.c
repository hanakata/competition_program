#include <stdio.h>
#include <string.h>
int main(void){
  char l[100],s[100];
  int n,m;

  scanf("%s",l);
  scanf("%s",s);
  n = strlen( l );
  m = strlen( s );

  if( n < m){
    printf( "%s", s );
  }else{
    printf( "%s", l );
  }
}
