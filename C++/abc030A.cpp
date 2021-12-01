#include <iostream>
#include <string>
using namespace std;

int main()
{
  double a, b, c, d, e, f;
  cin >> a >> b >> c >> d;
  e = b / a;
  f = d / c;
  if(e < f){
    cout << "AOKI" << std::endl;
  }
  if(e > f){
    cout << "TAKAHASHI" << std::endl;
  }
  if(e == f){
    cout << "DRAW" << std::endl;
  }

}
