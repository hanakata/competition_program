import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    int s = 0;
    String a = scan.nextLine();
    String b = scan.nextLine();
    int p = (Integer.parseInt(b) - 1) / 5;
    int q = Integer.parseInt(b) % 5;
    if(q == 0){
      s = 4;
    }else{
      s = q - 1;
    }
    StringBuffer buf = new StringBuffer();
    buf.append(a.charAt(p));
    buf.append(a.charAt(s));
    System.out.println(buf.toString());
  }
}
