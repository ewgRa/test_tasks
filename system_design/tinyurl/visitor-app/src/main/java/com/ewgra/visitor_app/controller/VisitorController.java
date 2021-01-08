package com.ewgra.visitor_app.controller;

import com.ewgra.visitor_app.model.UrlMap;
import com.ewgra.visitor_app.repository.UrlMapRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.web.servlet.view.RedirectView;

import java.util.HashMap;

@Slf4j
@RestController
public class VisitorController {

    @Autowired
    UrlMapRepository repository;

    @GetMapping("/{key}")
    public RedirectView visit(@PathVariable String key) {
        // FIXME XXX: Circuit Breaker
        // FIXME XXX: Exceptions
        // FIXME XXX: Timeout
        UrlMap urlMap = repository.findByShortUrl(key);

        if(urlMap == null) {
            throw new ResponseStatusException(
                HttpStatus.NOT_FOUND, "entity not found"
            );
        }

        return new RedirectView(urlMap.getLongUrl());
    }

    private ResponseEntity<Object> internalServerError(String message) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", false);
        response.put("message", message);

        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}