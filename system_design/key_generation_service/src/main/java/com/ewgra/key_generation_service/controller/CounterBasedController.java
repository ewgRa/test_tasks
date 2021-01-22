package com.ewgra.key_generation_service.controller;

import com.ewgra.key_generation_service.counter.ZooKeeperCounter;
import com.ewgra.key_generation_service.generator.CounterBasedGenerator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Slf4j
@RestController
public class CounterBasedController {
    @Autowired
    ZooKeeperCounter counter;

    @Autowired
    CounterBasedGenerator generator;

    @GetMapping("/counter-based/key")
    public ResponseEntity<Object> key() throws Exception {
        long next = counter.next();
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
        response.put("key", generator.generate(next));

        return new ResponseEntity<>(response, HttpStatus.OK);
    }
}