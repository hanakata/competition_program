import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String a = scan.nextLine();
    StringBuilder buf = new StringBuilder();
    buf.append(a);
    buf.append("s");
    System.out.println(buf.toString());
  }
}
