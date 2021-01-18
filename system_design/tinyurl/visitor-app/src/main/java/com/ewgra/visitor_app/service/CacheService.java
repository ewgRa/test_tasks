package com.ewgra.visitor_app.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.retry.annotation.CircuitBreaker;
import org.springframework.stereotype.Service;

@Service
public class CacheService {
    @Autowired
    private RedisTemplate<String, String> redisTemplate;

    @CircuitBreaker
    public String get(String key) {
        return redisTemplate.opsForValue().get(key);
    }

    @CircuitBreaker
    public void set(String key, String longUrl) {
        redisTemplate.opsForValue().set(key, longUrl);
    }

    public void clean() {
        redisTemplate.getConnectionFactory().getConnection().flushDb();
    }
}