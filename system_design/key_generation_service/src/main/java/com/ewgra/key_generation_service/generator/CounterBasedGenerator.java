package com.ewgra.key_generation_service.generator;

import com.ewgra.key_generation_service.Base62Converter;
import org.springframework.stereotype.Component;

// CounterBasedGenerator takes a number and converts it to key string.
@Component
public class CounterBasedGenerator {
    public String generate(long number) {
        return Base62Converter.convert(number);
    }
}