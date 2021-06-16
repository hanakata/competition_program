#include <iostream>
#include <string>
using namespace std;

int main()
{
  int n,m;
  cin >> n;
  m = n % 3;
  if( m == 0){
    cout << "YES" << endl;
  }else{
    cout << "NO" << endl;
  }
}