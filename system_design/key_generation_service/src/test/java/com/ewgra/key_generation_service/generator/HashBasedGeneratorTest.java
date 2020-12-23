package com.ewgra.key_generation_service.generator;

import com.ewgra.key_generation_service.Base62Converter;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.math.BigInteger;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Arrays;
import static org.assertj.core.api.Assertions.*;

@SpringBootTest
public class HashBasedGeneratorTest {
	@Test
	public void generate() {
		HashBasedGenerator generator = new HashBasedGenerator();
		String url = "http://foo.bar";
		String key = generator.generate(url);

		assertThat(key).isEqualTo("hOLYbryfIiwWcmlzTI260g");
	}
}