import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    int n = scan.nextInt();
    int m = scan.nextInt();
    if(n < m){
      System.out.println(m);
    }else{
      System.out.println(n);
    }
  }
}