#include <iostream>
#include <string>
using namespace std;

int main()
{
  int a, b, n, m;
  m = 0;
  while (cin >> a >> b) {
    n = a * b / 10;
    m += n;
  }
  cout << m << std::endl;
}
