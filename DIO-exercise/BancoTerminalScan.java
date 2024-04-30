import java.util.Scanner;

public class BancoTerminalScan {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Digite o saldo inicial: ");
        double saldo = scanner.nextDouble();

        System.out.print("Digite o valor a ser sacado: ");
        double valorSolicitado = scanner.nextDouble();

        if (saldo >= valorSolicitado) {
            saldo -= valorSolicitado;
            System.out.println("Saque realizado com sucesso!");
        } else {
            System.out.println("Saldo insuficiente");
        }

        System.out.println("Saldo atual: " + saldo);

        scanner.close();
    }
}
