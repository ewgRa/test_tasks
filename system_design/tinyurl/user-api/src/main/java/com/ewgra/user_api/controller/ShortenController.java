package com.ewgra.user_api.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Slf4j
@RestController
public class ShortenController {

    @PostMapping("/shorten")
    public ResponseEntity<Object> key() {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
        response.put("shortUrl", "piupiu");

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    private ResponseEntity<Object> internalServerError(String message) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", false);
        response.put("message", message);

        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}