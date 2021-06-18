#include <iostream>
#include <string>
using namespace std;

int main()
{
  int n,m,s;
  cin >> n;
  cin >> m;
  s = n % m;
  if( s == 0){
    cout << "YES" << endl;
  }else{
    cout << "NO" << endl;
  }
}