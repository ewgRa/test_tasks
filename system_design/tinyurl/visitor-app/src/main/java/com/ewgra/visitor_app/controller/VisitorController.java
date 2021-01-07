package com.ewgra.visitor_app.controller;

import com.ewgra.visitor_app.repository.UrlMapRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Slf4j
@RestController
public class VisitorController {

    @Autowired
    UrlMapRepository repository;

    @GetMapping("/{key}")
    public ResponseEntity<Object> visit(@PathVariable String key) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
        response.put("key", key);

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    private ResponseEntity<Object> internalServerError(String message) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", false);
        response.put("message", message);

        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}