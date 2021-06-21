#include <iostream>
#include <string>
using namespace std;

int main()
{
  std::string a;
  int b,p,q,s;
  cin >> a;
  cin >> b;
  p = (b - 1) / 5;
  q = b % 5;
  if( q == 0){
    s = 4;
  }else{
    s = q - 1;
  }
  cout << a[p] << a[s] << std::endl;
}
