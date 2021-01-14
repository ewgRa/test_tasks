package com.ewgra.shorten_api.controller;

import com.ewgra.shorten_api.model.UrlMap;
import com.ewgra.shorten_api.controller.request.ShortenRequest;
import com.ewgra.shorten_api.service.UrlMapService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import javax.validation.Valid;
import java.util.HashMap;

@Slf4j
@RestController
public class ShortenController {

    @Value("${shorten.api.visitor.app.url}")
    String visitorAppUrl;

    @Autowired
    UrlMapService urlMapService;

    @PostMapping("/shorten")
    public ResponseEntity<Object> key(@Valid @RequestBody ShortenRequest request) {
        UrlMap urlMap;

        try {
            urlMap = urlMapService.insert(request.getLongUrl());
        } catch (Exception e) {
            log.error("Unable to insert map", e);

            return internalServerError("Fail to save mapping");
        }

        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
        response.put("shortUrl", visitorAppUrl + "/" + urlMap.getShortUrl());

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    private ResponseEntity<Object> internalServerError(String message) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", false);
        response.put("message", message);

        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}