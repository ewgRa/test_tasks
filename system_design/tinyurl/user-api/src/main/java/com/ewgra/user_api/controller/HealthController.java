package com.ewgra.user_api.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Slf4j
@RestController
public class HealthController {

    @GetMapping("/health/liveness")
    public ResponseEntity<Object> liveness() {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    @GetMapping("/health/readiness")
    public ResponseEntity<Object> readiness() {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);

        // FIXME XXX: kgs connection
        // FIXME XXX postgresql connection

        return new ResponseEntity<>(response, HttpStatus.OK);
    }
}