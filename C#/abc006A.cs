using System;

class main{
  public static int Main(){
    int n = int.Parse(Console.ReadLine());
    int m = n % 3;
    if( m == 0 ){
      Console.WriteLine("YES");
    }else{
      Console.WriteLine("NO");
    }
    return 0;
  }
}