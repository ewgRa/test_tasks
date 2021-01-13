package com.ewgra.shorten_api.actuator;

import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.stereotype.Component;

@Component
public class KeyGenerationServiceHealthIndicator implements HealthIndicator {

    @Override
    public Health health() {
        // FIXME XXX: implement me
        return Health.up().build();
    }

}