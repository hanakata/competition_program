import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String a = scan.nextLine();
    int b = Integer.parseInt(a) / 2;
    System.out.println(b * b);
  }
}
