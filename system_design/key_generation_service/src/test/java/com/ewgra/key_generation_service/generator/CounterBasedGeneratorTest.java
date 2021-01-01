package com.ewgra.key_generation_service.generator;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;

import static org.assertj.core.api.Assertions.assertThat;

public class CounterBasedGeneratorTest {
    @ParameterizedTest
    @CsvSource({ "276098692598697967,lM2vEtkHyu", "839299365868340223,9999999999" })
    public void generate(Long number, String expected) {
        CounterBasedGenerator generator = new CounterBasedGenerator();
        String key = generator.generate(number);

        assertThat(key).isEqualTo(expected);
    }
}