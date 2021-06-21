#include <iostream>
#include <string>
using namespace std;

int main()
{
  int a,b,c,k,s,t,n,m,o;
  cin >> a >> b >> c >> k;
  cin >> s >> t;
  n = s * a;
  m = t * b;
  o = s + t;
  if( o >= k){
    cout << n + m - (o * c) << std::endl;
  }else{
    cout << n + m << std::endl;
  }
}
