#include <iostream>
#include <string>
using namespace std;

int main()
{
  int n,i,j;
  int l[4];
  int c[4] ={8,4,2,1};
  j = 0;
  i = 0;
  cin >> n;
  while (n != 0) {
    if(c[j] <= n){
      n -= c[j];
      l[i] = c[j];
      i++;
    }
    j++;
  }
  j = 0;
  cout << i << std::endl;
  while ( j < i ){
    cout << l[j] << std::endl;
    j++;
  }
}
