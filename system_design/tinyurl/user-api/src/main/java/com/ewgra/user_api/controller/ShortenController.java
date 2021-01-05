package com.ewgra.user_api.controller;

import com.ewgra.user_api.service.KeyGenerationService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Slf4j
@RestController
public class ShortenController {

    @Value("${visitorApiUrl}")
    String visitorApiUrl;

    @Autowired
    KeyGenerationService keyGenerationService;

    @PostMapping("/shorten")
    public ResponseEntity<Object> key() {
        String key;

        try {
            key = keyGenerationService.generate();
        } catch (Exception e) {
            // FIXME XXX: log
            return internalServerError("Can't generate a key");
        }

        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
        // FIXME XXX: store mapping longurl -> key
        response.put("shortUrl", visitorApiUrl + "/" + key);

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    private ResponseEntity<Object> internalServerError(String message) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", false);
        response.put("message", message);

        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}