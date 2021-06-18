using System;

class main{
  public static int Main(){
      int n = int.Parse(Console.ReadLine());
      int m = int.Parse(Console.ReadLine());
      int i = n % m;
      if(i == 0){
        Console.WriteLine("{0}", 0);
      }else{
        Console.WriteLine("{0}", m - i);
      }
      return 0;
  }
}
