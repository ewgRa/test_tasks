package com.ewgra.key_generation_service.generator;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.concurrent.ThreadLocalRandom;
import java.util.function.LongSupplier;

// RandomBasedGenerator randomly generates a number and converts it to the key via CounterBasedGenerator.
@Component
public class RandomBasedGenerator {
    static final long MAX_VALUE = 56800235584L; // 62⁶

    private final CounterBasedGenerator counterBasedGenerator;

    private final LongSupplier randomLongSupplier;

    @Autowired
    public RandomBasedGenerator(CounterBasedGenerator counterBasedGenerator) {
        this.counterBasedGenerator = counterBasedGenerator;
        this.randomLongSupplier = () -> ThreadLocalRandom.current().nextLong(MAX_VALUE);
    }

    public RandomBasedGenerator(CounterBasedGenerator counterBasedGenerator, LongSupplier randomLongSupplier) {
        this.counterBasedGenerator = counterBasedGenerator;
        this.randomLongSupplier = randomLongSupplier;
    }

    public String generate() {
        long number = this.randomLongSupplier.getAsLong();

        return counterBasedGenerator.generate(number);
    }

}