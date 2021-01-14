package com.ewgra.shorten_api.service;

import com.ewgra.shorten_api.model.UrlMap;
import com.ewgra.shorten_api.repository.UrlMapRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.retry.annotation.CircuitBreaker;
import org.springframework.stereotype.Service;

@Slf4j
@Service
public class UrlMapService {
    @Autowired
    KeyGenerationService keyGenerationService;

    @Autowired
    UrlMapRepository repository;

    @CircuitBreaker
    public UrlMap insert(String longUrl) throws Exception {
        UrlMap urlMap = new UrlMap();
        urlMap.setShortUrl(keyGenerationService.generate());
        urlMap.setLongUrl(longUrl);
        repository.save(urlMap);

        return urlMap;
    }
}