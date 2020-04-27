package jvmgo.book.ch07;

public class Function {

    public static void main(String[] args) {
        long x = max(3,5);
        System.out.println(x);
    }

    private static long max(long a,long b) {
        if (a > b) {
			return a;
        } else {
			return b;
		}
    }

}