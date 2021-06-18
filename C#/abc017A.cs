using System;

class main{
  public static int Main(){
    string line;
    int m = 0;
    while ((line = Console.ReadLine()) != null) {
      string[] l = line.Split(' ');
      int a = int.Parse(l[0]);
      int b = int.Parse(l[1]);
      int n = a * b / 10;
      m = m + n;
    }
    Console.WriteLine(m);
    return 0;
  }
}
