using System;

class main{
  public static int Main(){
    string[] str = { "A", "B", "C", "D", "E" };
    int i = 0;
    string n = Console.ReadLine();
    while(str[i] != n ){
      i ++;
    }
    Console.WriteLine("{0}", i + 1);
    return 0;
  }
}
