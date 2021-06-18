#include <iostream>
#include <string>
using namespace std;

int main()
{
  int n,m,i;
  cin >> n;
  cin >> m;
  i = n % m;
  if(i == 0){
    cout << 0 << endl;
  }else{
    cout << m - i << endl;
  }
}
