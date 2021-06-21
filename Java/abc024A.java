import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String p = scan.nextLine();
    String q = scan.nextLine();
    String[] a = p.split(" ", 0);
    String[] s = q.split(" ", 0);
    int n = Integer.parseInt(s[0]) * Integer.parseInt(a[0]);
    int m = Integer.parseInt(s[1]) * Integer.parseInt(a[1]);
    int o = Integer.parseInt(s[0]) + Integer.parseInt(s[1]);
    if(Integer.parseInt(a[3]) <= o){
      System.out.println(n + m - (o * Integer.parseInt(a[2])));    
    }else{
      System.out.println(n + m);
    }
  }
}
