using System;

class main{
  public static int Main(){
    string line;
    int m = 0;
    line = Console.ReadLine();
    string[] n = line.Split(' ');
    for(int i = 0;i < 2; i++){
      for(int j = 2; j > i; j--){
        if(int.Parse(n[j]) < int.Parse(n[j-1])){
          string t = n[j-1];
          n[j-1] = n[j];
          n[j] = t;
        }
      }
    }
    Console.WriteLine(n[1]);
    return 0;
  }
}
