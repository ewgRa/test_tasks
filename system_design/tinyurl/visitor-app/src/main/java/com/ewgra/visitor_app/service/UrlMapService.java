package com.ewgra.visitor_app.service;

import com.ewgra.visitor_app.model.UrlMap;
import com.ewgra.visitor_app.repository.UrlMapRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.QueryHints;
import org.springframework.retry.annotation.CircuitBreaker;
import org.springframework.stereotype.Repository;
import org.springframework.stereotype.Service;

import javax.persistence.QueryHint;

@Service
public class UrlMapService {
    @Autowired
    private UrlMapRepository repository;

    @CircuitBreaker
    public UrlMap findByShortUrl(String shortUrl) {
        return repository.findByShortUrl(shortUrl);
    }
}