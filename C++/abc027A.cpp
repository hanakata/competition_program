#include <iostream>
#include <string>
using namespace std;

int main()
{
  std::string a,b,c;
  cin >> a >> b >> c;
  if( a == b){
    cout << c << std::endl;
  }else if( a == c ){
      cout << b << std::endl;
  }else if( b == c ){
      cout << a << std::endl;
  }
}
