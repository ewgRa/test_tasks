package com.ewgra.key_generation_service;

public class Base62Converter {
	private static final char[] ALPHABET = ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789").toCharArray();
	private static final long base = 62;

	public static String convert(long number) {
		StringBuilder sb = new StringBuilder("");

		while (number > 0) {
			sb.append(ALPHABET[(int) (number % base)]);
			number /= base;
		}

		return sb.toString();
	}
}