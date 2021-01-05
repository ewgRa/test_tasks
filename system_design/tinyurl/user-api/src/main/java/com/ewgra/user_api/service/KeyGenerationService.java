package com.ewgra.user_api.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

@Slf4j
@Service
public class KeyGenerationService {
    @Value("${keyGenerationServiceUrl}") String url;

    public String generate() {
        // FIXME XXX: make request to real service
        return url;
    }
}