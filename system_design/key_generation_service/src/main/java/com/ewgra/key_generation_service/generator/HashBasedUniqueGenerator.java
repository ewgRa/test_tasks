package com.ewgra.key_generation_service.generator;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.concurrent.ThreadLocalRandom;
import java.util.function.LongSupplier;

// HashBasedUniqueGenerator we can use as fallback when key already presented in storage.
// There is several strategies of making key unique. We can add as a salt user IP,
// sequence, random data or current time, something that unique enough at this point of time for such key.
@Component
public class HashBasedUniqueGenerator {
    private final HashBasedGenerator hashBasedGenerator;

    private final LongSupplier randomLongSupplier;

    @Autowired
    public HashBasedUniqueGenerator(HashBasedGenerator hashBasedGenerator) {
        this.hashBasedGenerator = hashBasedGenerator;
        this.randomLongSupplier = () -> ThreadLocalRandom.current().nextLong();
    }

    public HashBasedUniqueGenerator(HashBasedGenerator hashBasedGenerator, LongSupplier randomLongSupplier) {
        this.hashBasedGenerator = hashBasedGenerator;
        this.randomLongSupplier = randomLongSupplier;
    }

    public String generate(String data) {
        String random = String.valueOf(randomLongSupplier.getAsLong());

        return hashBasedGenerator.generate(data + random);
    }
}