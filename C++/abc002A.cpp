#include <iostream>
#include <string>
using namespace std;

int main()
{
  int n,m;
  cin >> n;
  cin >> m;
  if( n < m){
    cout << m;
    cout << "\n";
  }else{
    cout << n;
    cout << "\n";
  }
}