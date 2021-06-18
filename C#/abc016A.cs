using System;

class main{
  public static int Main(){
    string[] a = Console.ReadLine().Split(' ');
    int n = int.Parse(a[0]);
    int m = int.Parse(a[1]);
    int s = n % m;
    if( s == 0 ){
      Console.WriteLine("YES");
    }else{
      Console.WriteLine("NO");
    }
    return 0;
  }
}
