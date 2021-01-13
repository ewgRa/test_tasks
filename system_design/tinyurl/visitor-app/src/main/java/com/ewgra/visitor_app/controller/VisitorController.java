package com.ewgra.visitor_app.controller;

import com.ewgra.visitor_app.model.UrlMap;
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
    UrlMapService urlMapService;

    @GetMapping("/{key}")
    public RedirectView visit(@PathVariable String key) {
        // FIXME XXX redis connection healthcheck

        UrlMap urlMap;

        try {
            urlMap = urlMapService.findByShortUrl(key);
        } catch (Exception e) {
            log.error("Unable to find short url", e);

            throw new ResponseStatusException(
                    HttpStatus.INTERNAL_SERVER_ERROR, "Internal server error"
            );
        }

        if(urlMap == null) {
            throw new ResponseStatusException(
                HttpStatus.NOT_FOUND, "map not found"
            );
        }

        return new RedirectView(urlMap.getLongUrl());
    }
}