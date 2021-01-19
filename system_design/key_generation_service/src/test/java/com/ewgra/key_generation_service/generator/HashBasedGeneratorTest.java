package com.ewgra.key_generation_service.generator;

import org.junit.jupiter.api.Test;
import org.springframework.test.context.ActiveProfiles;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.*;

@ActiveProfiles("test")
public class HashBasedGeneratorTest {
    @Test
    public void generate() {
        CounterBasedGenerator counterBasedGeneratorMock = mock(CounterBasedGenerator.class);
        when(counterBasedGeneratorMock.generate(anyLong())).thenAnswer(i -> i.getArguments()[0].toString());

        HashBasedGenerator generator = new HashBasedGenerator(counterBasedGeneratorMock, "some_salt");
        String key = generator.generate("http://foo.bar");
        assertThat(key).isEqualTo("276098692598697967");
    }
}