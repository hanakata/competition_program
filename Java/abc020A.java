import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String n = scan.nextLine();
    if(Integer.parseInt(n) == 1){
      System.out.println("ABC");
    }else{
      System.out.println("chokudai");
    }
  }
}
