package com.ewgra.user_api.repository;

import com.ewgra.user_api.model.UrlMap;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UrlMapRepository extends JpaRepository<UrlMap, String> {
    UrlMap findByShortUrl(String shortUrl);
}