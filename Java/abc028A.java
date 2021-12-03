import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String a = scan.nextLine();
    int b = Integer.parseInt(a);
    if( b <= 59){
      System.out.println("Bad");
    }else if(60 <= b && b <= 89){
      System.out.println("Good");
    }else if(90 <= b && b <= 99){
      System.out.println("Great");
    }else{
      System.out.println("Perfect");
    }
  }
}
