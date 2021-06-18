#include <iostream>
#include <string>
using namespace std;

int main()
{
  int n[3];
  int r[3] ={1,1,1};
  int a, i, j;
  while (cin >> a) {
    n[i] = a;
    i++;
  }
  for (i = 0;i < 3; i++){
    for(j = 0; j < 3; j++){
      if(n[i] < n[j]){
        r[i] = r[i] + 1;
      }
    }
  }
  cout << r[0] << std::endl;
  cout << r[1] << std::endl;
  cout << r[2] << std::endl;
}
