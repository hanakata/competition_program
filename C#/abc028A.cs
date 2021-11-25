using System;

class main{
  public static int Main(){
    string a = Console.ReadLine();
    int b = int.Parse(a)
    if(b <= 59){
      Console.WriteLine("Bad");
    }
    if(60 <= b && b <= 89){
      Console.WriteLine("Good");
    }
    if(90 <= b && b <= 99){
      Console.WriteLine("Great");
    }
    if(100 == b){
      Console.WriteLine("Perfect");
    }
    return 0;
  }
}
