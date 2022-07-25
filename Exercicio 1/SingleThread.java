package test;
import java.util.*;

public class SingleThread {
	
	public static boolean isPrimeNumber(int c) {
		for(int i = 2; i <= Math.sqrt(c); ++i) {
			if(c%i == 0) {
				return false;
			}
		}
		return true;
	}
	
	public static void task(int n) {
		int primes = 0;
		 for(int i = 2; i < n; ++i) {
			 if(isPrimeNumber(i)) {
				 ++primes;
			 }
		 }
		 System.out.println("There are " + primes + " primes between 1 and " + n);
	}
	
	public static void main(String[] args) {
		double st = System.nanoTime();
		task(1000000);
        double et = System.nanoTime();
		System.out.println(((et-st)/1000000));
	}
}
