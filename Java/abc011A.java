import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    int n = scan.nextInt();
    if(n == 12){
      System.out.println(1);
    }else{
      System.out.println(n + 1);
    }
  }
}
