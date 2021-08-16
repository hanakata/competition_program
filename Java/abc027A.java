import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String a = scan.nextLine();
    String[] s = a.split(" ", 0);
    if(s[0].equals(s[1])){
      System.out.println(s[2]);
    }else if(s[0].equals(s[2]){
      System.out.println(s[1]);
    }else{
      System.out.println(s[0]);
    }
  }
}
