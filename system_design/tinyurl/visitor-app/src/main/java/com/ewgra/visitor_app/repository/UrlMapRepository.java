package com.ewgra.visitor_app.repository;

import com.ewgra.visitor_app.model.UrlMap;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UrlMapRepository extends JpaRepository<UrlMap, String> {
    UrlMap findByShortUrl(String shortUrl);
}