using System;

class main{
  public static int Main(){
      string[] a = Console.ReadLine().Split(' ');
      int n = int.Parse(a[0]);
      int m = int.Parse(a[1]);
      Console.WriteLine("{0}", m - n + 1);
      return 0;
  }
}
