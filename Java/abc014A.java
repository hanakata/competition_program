import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    int n = scan.nextInt();
    int m = scan.nextInt();
    int i = n % m;
    if(i == 0){
      System.out.println(0);
    }else{
      System.out.println(m - i);
    }
  }
}
