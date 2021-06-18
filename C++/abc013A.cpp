#include <iostream>
#include <string>
using namespace std;
 
int main()
{
  std::string str("ABCDE");
  std::string a,b;
  int i;
  i = 0;
  cin >> a;
  b = str[i];
  while( b != a ){
  i++;
  b = str[i];
  }
  cout << i + 1 << endl;
}
