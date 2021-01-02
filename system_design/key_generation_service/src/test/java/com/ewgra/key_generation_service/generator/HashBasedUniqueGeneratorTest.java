package com.ewgra.key_generation_service.generator;

import org.junit.jupiter.api.Test;

import java.util.function.LongSupplier;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.*;

public class HashBasedUniqueGeneratorTest {
    @Test
    public void generate() {
        HashBasedGenerator hashBasedGeneratorMock = mock(HashBasedGenerator.class);
        when(hashBasedGeneratorMock.generate(anyString())).thenAnswer(i -> i.getArguments()[0]);

        long randomLong = 1000L;
        LongSupplier randomLongSupplierMock = mock(LongSupplier.class);
        when(randomLongSupplierMock.getAsLong()).thenReturn(randomLong);

        HashBasedUniqueGenerator generator = new HashBasedUniqueGenerator(hashBasedGeneratorMock,
                randomLongSupplierMock);
        String key = generator.generate("http://foo.bar");

        assertThat(key).isEqualTo("http://foo.bar" + randomLong);
    }

    @Test
    public void generateIsRealUnique() {
        HashBasedGenerator hashBasedGeneratorMock = mock(HashBasedGenerator.class);
        when(hashBasedGeneratorMock.generate(anyString())).thenAnswer(i -> i.getArguments()[0]);

        HashBasedUniqueGenerator generator = new HashBasedUniqueGenerator(hashBasedGeneratorMock);
        String key = generator.generate("http://foo.bar");
        String anotherKey = generator.generate("http://foo.bar");

        assertThat(key).isNotEqualTo(anotherKey);
    }
}