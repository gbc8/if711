package test;
import java.util.*;

class Primes extends Thread{
	public static int currentNumber = 2;
	public static int primes = 0;
	public static int n = 1000000;
	public int threadId;
	
	public Primes(int id) {
		threadId = id;
	}
	
	public static synchronized int getNextNumber() {
		int number = currentNumber;
		currentNumber++;
		return number;
	}
	
	public static synchronized void incPrimes() {
		++primes;
	}
	
	public boolean isPrimeNumber(int n) {
		for(int i = 2; i <= Math.sqrt(n); ++i) {
			if(n%i == 0) {
				return false;
			}
		}
		return true;
	}
	
	public void run(){
		while(true) {
			int number = getNextNumber();
			if(number >= n) {
				break;
			}
			if(isPrimeNumber(number)) {
				incPrimes();
			}
		}
	}
}

public class MultiThread {
	
	public static void main(String[] args) {
		int threads = 64;
		Primes[] primes = new Primes[threads];
		double st = System.nanoTime();
		for(int i = 0; i < threads; ++i) {
			primes[i] = new Primes(i+1);
			primes[i].start();
		}
		for(int i = 0; i < threads; ++i) {
			try {
				primes[i].join();
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
        double et = System.nanoTime();
		System.out.println("There are " + Primes.primes + " primes between 1 and " + Primes.n);
		System.out.println(threads + " Threads: " + ((et-st)/1000000) + " ms");
	}
}
