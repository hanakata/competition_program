using System;

class main{
  public static int Main(){
    int[] l = new int[4];
    int[] c = { 8, 4, 2, 1 };
    int i = 0;
    int j = 0;
    int n = int.Parse(Console.ReadLine());
    while(n != 0){
      if(c[j] <= n){
       n = n - c[j];
       l[i] = c[j];
       i = i + 1;
      }
      j = j + 1;
    }
    j = 0;
    Console.WriteLine(i);
    while(j < i){
      Console.WriteLine(l[j]);
      j = j + 1;
    }
    return 0;
  }
}
