#include <iostream>
#include <string>
using namespace std;

int main()
{
  int a;
  cin >> a;
  if( a <= 59){
    cout << "Bad" << std::endl;
  }
  if(60 <= a && a <= 89){
    cout << "Good" << std::endl;
  }
  if(90 <= a && a <= 99){
    cout << "Great" << std::endl;
  }
  if(100 == a){
    cout << "Perfect" << std::endl;
  }
}
