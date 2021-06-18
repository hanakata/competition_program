#include <iostream>
#include <string>
using namespace std;

int main()
{
  std::string l,s;
  cin >> l;
  cin >> s;

  if (l.size() < s.size()){
    cout << s << endl;
  }else{
    cout << l << endl;
  }
}
