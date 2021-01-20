package com.ewgra.shorten_api.controller;

import com.ewgra.shorten_api.model.UrlMap;
import com.ewgra.shorten_api.repository.UrlMapRepository;
import org.json.JSONObject;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.request.MockHttpServletRequestBuilder;

import static org.assertj.core.api.Assertions.assertThat;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
@AutoConfigureMockMvc
@ActiveProfiles("test")
public class ShortenControllerTest {
    @Value("${shorten.api.visitor.app.url}")
    private String visitorAppUrl;

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private UrlMapRepository urlMapRepository;

    @BeforeEach
    private void setUp() {
        urlMapRepository.deleteAll();
    }

    @Test
    public void shorten() throws Exception {
        MockHttpServletRequestBuilder request = post("/shorten").contentType(MediaType.APPLICATION_JSON)
                .content("{\"longUrl\": \"http://test.com\"}");

        MvcResult result = this.mockMvc.perform(request).andExpect(status().isOk()).andReturn();
        assertThat(urlMapRepository.count()).isEqualTo(1);
        UrlMap urlMap = urlMapRepository.findAll().get(0);
        JSONObject response = new JSONObject(result.getResponse().getContentAsString());
        assertThat(response.getString("shortUrl")).isEqualTo(visitorAppUrl + "/" + urlMap.getShortUrl());
    }
}