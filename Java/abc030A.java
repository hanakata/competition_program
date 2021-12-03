import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String a = scan.nextLine();
    String[] s = a.split(" ", 0);
    double b = Double.parseDouble(s[0]);
    double c = Double.parseDouble(s[1]);
    double d = Double.parseDouble(s[2]);
    double e = Double.parseDouble(s[3]);
    if(c/b < e/d){
      System.out.println("AOKI");
    }else if(c/b > e/d){
      System.out.println("TAKAHASHI");
    }else{
      System.out.println("DRAW");
    }
  }
}
