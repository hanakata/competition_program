#include <iostream>
#include <string>
using namespace std;

int main()
{
  int a,b,c,w,i,m,j;
  j = 0;
  cin >> a >> b >> c;
  cin >> w;
  if( b <= w && w <= c){
    j++;
  }
  for(i = 0;i< a - 1;i++){
    cin >> m;
    w = w + m;
    if( b <= w && w <= c){
      j++;
    }
  }
  cout << j << std::endl;
}
