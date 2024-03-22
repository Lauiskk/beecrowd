import java.util.Scanner;

public class Main {
    public static void main(String[] args){
    Scanner sc1 = new Scanner(System.in);
    double raio = sc1.nextDouble();
    double n = 3.14159;
    double area = n * (raio * raio);
    System.out.printf("A=%.4f\n",area);
    }
 
}