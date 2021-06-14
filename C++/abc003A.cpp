#include <iostream>
#include <string>
using namespace std;

int main()
{
  int i,n,m,s;
  i = 1;
  cin >> n;
  while(i <= n){
    m += i * 10000;
    i++
  }
  s = m / n;
  cout << s;
  cout << "\n";
}