package com.ewgra.key_generation_service.generator;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.util.DigestUtils;

import java.nio.charset.StandardCharsets;

// HashBasedGenerator takes data and returns a hash of data. Hash will be the same for two same URLs.
// If we have a requirement to generate unique keys even for the same URL we need to use HashBasedUniqueGenerator as
// a fallback, when the key already presented in our storage. In this case hash based approach transforms more or less
// to random based approach.
// To keep hash harder to guess, we will add to data some "secret" as a salt.
@Component
public class HashBasedGenerator {
    private final String salt;
    private final CounterBasedGenerator counterBasedGenerator;

    public HashBasedGenerator(CounterBasedGenerator counterBasedGenerator, @Value("${hashSalt}") String salt) {
        this.counterBasedGenerator = counterBasedGenerator;
        this.salt = salt;
    }

    public String generate(String data) {
        String hash = DigestUtils.md5DigestAsHex((data + salt).getBytes(StandardCharsets.UTF_8));
        // we can take max 64 bits (radix 16 - 4bit, 16*4 = 64), and to be safe with max value, we reduce it to 60 bits
        // that should be enough
        long number = Long.parseUnsignedLong(hash.substring(0, 15), 16);

        return counterBasedGenerator.generate(number);
    }
}