package com.ewgra.key_generation_service.generator;

import com.ewgra.key_generation_service.Base62Converter;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.math.BigInteger;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Arrays;

// CounterBasedGenerator takes a number and converts it to key string.
@Component
public class CounterBasedGenerator {
	public String generate(long number) {
		return Base62Converter.convert(number);
	}
}