using System;

class main{
  public static int Main(){
      string[] a = Console.ReadLine().Split(' ');
      int n = int.Parse(a[0]);
      if( n == 12 ){
        Console.WriteLine("{0}", 1);
      }else{
        Console.WriteLine("{0}", n + 1);
      }
      return 0;
  }
}