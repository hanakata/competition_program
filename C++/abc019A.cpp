#include <iostream>
#include <string>
using namespace std;

int main()
{
  int a, b, c, i, j, t;
  cin >> a >> b >> c;
  int n[ ] = {a,b,c};
  for(i = 0;i < 2; i++){
    for(j = 2; j > i; j--){
      if(n[j] < n[j-1]){
      t = n[j-1];
      n[j-1] = n[j];
      n[j] = t;
      }
    }
  }
  cout << n[1] << std::endl;
}
