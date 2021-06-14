using System;

class main{
  public static int Main(){
    int i = 1;
    int m = 0;
    int n = int.Parse(Console.ReadLine());
    while(i <= n){
      m = m + i * 10000;
      i ++;
    }
    int s = m / n;
    Console.WriteLine("{0}", s);
    return 0;
  }
}