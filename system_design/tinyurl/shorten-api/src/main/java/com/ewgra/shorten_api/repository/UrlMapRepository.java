package com.ewgra.shorten_api.repository;

import com.ewgra.shorten_api.model.UrlMap;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.QueryHints;
import org.springframework.stereotype.Repository;

import javax.persistence.QueryHint;

@Repository
public interface UrlMapRepository extends JpaRepository<UrlMap, String> {
    @QueryHints(value = { @QueryHint(name = "org.hibernate.timeout", value = "1") })
    UrlMap findByShortUrl(String shortUrl);

    @QueryHints(value = { @QueryHint(name = "org.hibernate.timeout", value = "1") })
    @Override
    <S extends UrlMap> S save(S entity);
}