package com.ewgra.key_generation_service;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.retry.annotation.EnableRetry;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@SpringBootApplication
@EnableRetry
public class KeyGenerationServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(KeyGenerationServiceApplication.class, args);
    }

    @Bean
    public WebMvcConfigurer corsConfigurer(@Value("${kgs.cors.origins}") String corsOrigins) {
        return new WebMvcConfigurer() {
            @Override
            public void addCorsMappings(CorsRegistry registry) {
                registry.addMapping("/**").allowedMethods("*").allowedOrigins(corsOrigins);
            }
        };
    }
}
