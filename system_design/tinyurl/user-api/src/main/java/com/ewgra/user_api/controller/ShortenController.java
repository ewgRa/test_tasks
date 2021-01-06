package com.ewgra.user_api.controller;

import com.ewgra.user_api.model.UrlMap;
import com.ewgra.user_api.repository.UrlMapRepository;
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

    @Autowired
    UrlMapRepository repository;

    @PostMapping("/shorten")
    public ResponseEntity<Object> key() {
        String key;

        try {
            key = keyGenerationService.generate();
        } catch (Exception e) {
            // FIXME XXX: log
            return internalServerError("Can't generate a key");
        }

        UrlMap urlMap = new UrlMap();
        urlMap.setShortUrl(key);
        urlMap.setLongUrl("fff");

        // FIXME XXX: do it better
        // FIXME XXX: CircuitBreaker
        // FIXME XXX: Timeout?
        try {
            repository.save(urlMap);
        } catch (Exception e) {
            return internalServerError("Fail to save mapping");
        }

        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
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