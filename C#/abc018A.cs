using System;

class main{
  public static int Main(){
    string line;
    int[] n = new int[3];
    int[] r = { 1, 1, 1 };
    int i = 0;
    while ((line = Console.ReadLine()) != null) {
      string[] l = line.Split('\n');
      int a = int.Parse(l[0]);
      n[i] = a;
      i = i + 1;
    }
    for(i = 0;i < 3; i++){
      for (int j = 0; j < 3; j++){
        if(n[i] < n[j]){
          r[i] = r[i] + 1;
        }
      }
    }
    Console.WriteLine(r[0]);
    Console.WriteLine(r[1]);
    Console.WriteLine(r[2]);
    return 0;
  }
}
