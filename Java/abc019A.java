import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String n = scan.nextLine();
    String[] s = n.split(" ", 0);
    for(int i = 0;i < 2; i++){
      for(int j = 2; j > i; j--){
        if(Integer.parseInt(s[j]) < Integer.parseInt(s[j-1])){
          String t = s[j-1];
          s[j-1] = s[j];
          s[j] = t;
        }
      }
    }
    System.out.println(s[1]);
  }
}
