import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    int n = scan.nextInt();
    int l[] = new int[4];
    int c[] = {8, 4, 2, 1};
    int i = 0;
    int j = 0;
    while(n != 0){
      if(c[j] <= n){
       n = n - c[j];
       l[i] = c[j];
       i = i + 1;
      }
      j = j + 1;
    }
    j = 0;
    System.out.println(i);
    while(j < i){
      System.out.println(l[j]);
      j = j + 1;
    }
  }
}
