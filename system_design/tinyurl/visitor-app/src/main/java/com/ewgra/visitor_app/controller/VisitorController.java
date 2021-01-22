package com.ewgra.visitor_app.controller;

import com.ewgra.visitor_app.model.UrlMap;
import com.ewgra.visitor_app.service.CacheService;
import com.ewgra.visitor_app.service.UrlMapService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.web.servlet.view.RedirectView;

@Slf4j
@RestController
public class VisitorController {

    @Autowired
    private CacheService cacheService;

    @Autowired
    UrlMapService urlMapService;

    @GetMapping("/{key}")
    public RedirectView visit(@PathVariable String key) {
        String longUrl = cacheService.get(key);

        if (longUrl == null) {
            longUrl = getFromDatabase(key);
            cacheService.set(key, longUrl);
        }

        return new RedirectView(longUrl);
    }

    private String getFromDatabase(String key) {
        UrlMap urlMap = urlMapService.findByShortUrl(key);

        if (urlMap == null) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "map not found");
        }

        return urlMap.getLongUrl();
    }
}