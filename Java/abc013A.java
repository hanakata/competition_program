import java.util.Scanner;

class Main{
  public static void main(String args[]){
    Scanner scan = new Scanner(System.in);
    String[] str = { "A", "B", "C", "D", "E" };
    int i = 0;
    String n = scan.nextLine();
    while(!(str[i].equals(n))){
      i ++;
    }
    System.out.println(i + 1);
  }
}
