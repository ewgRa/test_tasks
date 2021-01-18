package com.ewgra.visitor_app.controller;

import com.ewgra.visitor_app.model.UrlMap;
import com.ewgra.visitor_app.repository.UrlMapRepository;
import com.ewgra.visitor_app.service.CacheService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.web.servlet.MockMvc;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.header;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
@AutoConfigureMockMvc
@ActiveProfiles("test")
public class VisitorControllerTest {
    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private CacheService cacheService;

    @Autowired
    private UrlMapRepository urlMapRepository;

    @BeforeEach
    private void setUp() {
        cacheService.clean();
        urlMapRepository.deleteAll();
        UrlMap urlMap = new UrlMap();
        urlMap.setLongUrl("http://long_url.com");
        urlMap.setShortUrl("short");
        urlMapRepository.save(urlMap);
    }

    @Test
    public void nonExistMap() throws Exception {
        this.mockMvc.perform(get("/non_exists_key")).andExpect(status().isNotFound());
    }

    @Test
    public void visit() throws Exception {
        this.mockMvc.perform(get("/short")).andExpect(status().isFound())
                .andExpect(header().string("location", "http://long_url.com"));
    }

    @Test
    public void cache() throws Exception {
        this.mockMvc.perform(get("/short"));
        urlMapRepository.deleteAll();

        this.mockMvc.perform(get("/short")).andExpect(status().isFound())
                .andExpect(header().string("location", "http://long_url.com"));
    }
}