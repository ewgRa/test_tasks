package com.ewgra.key_generation_service.generator;

import org.junit.jupiter.api.Test;
import org.springframework.test.context.ActiveProfiles;

import java.util.function.LongSupplier;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.*;

@ActiveProfiles("test")
public class RandomBasedGeneratorTest {
    @Test
    public void generate() {
        CounterBasedGenerator counterBasedGenerator = mock(CounterBasedGenerator.class);
        when(counterBasedGenerator.generate(anyLong())).thenAnswer(i -> i.getArguments()[0].toString());

        long randomLong = 1000L;
        LongSupplier randomLongSupplierMock = mock(LongSupplier.class);
        when(randomLongSupplierMock.getAsLong()).thenReturn(randomLong);

        RandomBasedGenerator generator = new RandomBasedGenerator(counterBasedGenerator, randomLongSupplierMock);
        String key = generator.generate();

        assertThat(key).isEqualTo(String.valueOf(randomLong));
    }

    @Test
    public void generateIsRealRandom() {
        CounterBasedGenerator counterBasedGenerator = mock(CounterBasedGenerator.class);
        when(counterBasedGenerator.generate(anyLong())).thenAnswer(i -> i.getArguments()[0].toString());

        RandomBasedGenerator generator = new RandomBasedGenerator(counterBasedGenerator);
        String key = generator.generate();
        String anotherKey = generator.generate();

        assertThat(key).isNotEqualTo(anotherKey);
    }
}