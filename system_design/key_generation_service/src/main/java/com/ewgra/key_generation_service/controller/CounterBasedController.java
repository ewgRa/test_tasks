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
    public ResponseEntity<Object> key() {
        long next;

        try {
            next = counter.next();
        } catch (Exception e) {
            log.error("Error when getting next key: " + e.getMessage());

            return internalServerError("Internal server error when getting next counter value");
        }

        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
        response.put("key", generator.generate(next));

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    private ResponseEntity<Object> internalServerError(String message) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", false);
        response.put("message", message);

        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}