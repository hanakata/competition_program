import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String l = scan.nextLine();
    String s = scan.nextLine();
    if(l.length() < s.length()){
      System.out.println(s);
    }else{
      System.out.println(l);
    }
  }
}
