package com.ewgra.shorten_api.service;

import lombok.extern.slf4j.Slf4j;
import org.apache.http.HttpResponse;
import org.apache.http.HttpStatus;
import org.apache.http.client.config.RequestConfig;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.util.EntityUtils;
import org.json.JSONObject;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.retry.annotation.CircuitBreaker;
import org.springframework.stereotype.Service;

@Slf4j
@Service
public class KeyGenerationService implements DisposableBean {
    private final int CONNECT_TIMEOUT = 100;
    private final int SOCKET_TIMEOUT = 500;
    private final int CONNECTION_REQUEST_TIMEOUT = 50;

    @Value("${shorten.api.kgs.url}")
    private String url;

    private CloseableHttpClient client;

    @CircuitBreaker
    public String generate() throws Exception {
        HttpResponse response = client().execute(new HttpGet(url + "/counter-based/key"));

        if (response.getStatusLine().getStatusCode() != HttpStatus.SC_OK) {
            throw new Exception("Key Generation Service respond with unexpected code.");
        }

        String responseBody = EntityUtils.toString(response.getEntity());
        JSONObject json = new JSONObject(responseBody);

        if (!json.has("key") || json.getString("key").isEmpty()) {
            throw new Exception("Key Generation Service respond with invalid key");
        }

        return json.getString("key");
    }

    private CloseableHttpClient client() {
        if (client == null) {
            RequestConfig config = RequestConfig.custom()
                    .setConnectTimeout(CONNECT_TIMEOUT)
                    .setSocketTimeout(SOCKET_TIMEOUT)
                    .setConnectionRequestTimeout(CONNECTION_REQUEST_TIMEOUT)
                    .build();

            client = HttpClientBuilder.create().setDefaultRequestConfig(config).build();
        }

        return client;
    }

    @Override
    public void destroy() throws Exception {
        if (client != null) {
            client.close();
        }
    }
}