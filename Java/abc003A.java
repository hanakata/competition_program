import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    int i = 1;
    int m = 0;
    int n = scan.nextInt();
    while(i <= n){
      m = m + i * 10000;
      i ++;
    }
    int s = m / n;
    System.out.println(s);
  }
}
Javascript(Node.js)