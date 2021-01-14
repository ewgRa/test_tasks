package com.ewgra.shorten_api.actuator;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.http.ResponseEntity;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

@Component
public class KeyGenerationServiceHealthIndicator implements HealthIndicator {
    @Value("${shorten.api.kgs.url}")
    private String url;

    @Override
    public Health health() {
        ResponseEntity<String> response = createRestTemplate().getForEntity(
            url+"/actuator/health/readiness",
            String.class
        );

        if (!response.getStatusCode().is2xxSuccessful()) {
            return Health.down().build();
        }

        return Health.up().build();
    }

    private RestTemplate createRestTemplate() {
        HttpComponentsClientHttpRequestFactory requestFactory = new HttpComponentsClientHttpRequestFactory();
        requestFactory.setConnectionRequestTimeout(300);
        requestFactory.setConnectTimeout(300);
        requestFactory.setReadTimeout(300);

        return new RestTemplate(requestFactory);
    }
}