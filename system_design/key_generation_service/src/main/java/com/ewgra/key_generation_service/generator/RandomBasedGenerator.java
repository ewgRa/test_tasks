package com.ewgra.key_generation_service.generator;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.w3c.dom.css.Counter;

import java.math.BigInteger;
import java.util.concurrent.ThreadLocalRandom;

// Random
@Component
public class RandomBasedGenerator {
	static final long MAX_VALUE = 56800235584L; // 62⁶

	@Autowired
	CounterBasedGenerator counterBasedGenerator;

	public String generate() {
		long number = ThreadLocalRandom.current().nextLong(MAX_VALUE);

		return counterBasedGenerator.generate(number);
	}

}