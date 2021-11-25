using System;

class main{
  public static int Main(){
    string[] a = Console.ReadLine().Split(' ');

    double b = Double.Parse(a[1]) / Double.Parse(a[0]);
    double c = Double.Parse(a[3]) / Double.Parse(a[2]);
    if(b < c){
      Console.WriteLine("AOKI");
    }

    if(b > c){
      Console.WriteLine("TAKAHASHI");    
    }

    if(b == c){
      Console.WriteLine("DRAW");      
    }
    return 0;
  }
}
