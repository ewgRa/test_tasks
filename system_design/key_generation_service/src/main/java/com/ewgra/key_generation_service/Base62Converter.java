package com.ewgra.key_generation_service;

public class Base62Converter {
    private static final String ALPHABET_STR = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    private static final char[] ALPHABET = ALPHABET_STR.toCharArray();
    private static final long BASE = 62;

    public static String convert(long number) {
        StringBuilder sb = new StringBuilder("");

        while (number > 0) {
            sb.append(ALPHABET[(int) (number % BASE)]);
            number /= BASE;
        }

        return sb.toString();
    }
}