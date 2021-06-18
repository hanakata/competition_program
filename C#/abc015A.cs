using System;

class main{
  public static int Main(){
    string l = Console.ReadLine();
    string s = Console.ReadLine();
    if (l.Length < s.Length){
      Console.WriteLine("{0}", s);
    }else{
      Console.WriteLine("{0}", l);
    }
    return 0;
  }
}
