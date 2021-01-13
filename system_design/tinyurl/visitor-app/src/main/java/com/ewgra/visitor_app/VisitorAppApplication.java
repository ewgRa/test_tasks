package com.ewgra.visitor_app;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.retry.annotation.EnableRetry;

@SpringBootApplication
@EnableRetry
public class VisitorAppApplication {

    // FIXME XXX: redis cache
    public static void main(String[] args) {
        SpringApplication.run(VisitorAppApplication.class, args);
    }
}
