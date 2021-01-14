package com.ewgra.shorten_api.controller.request;

import org.hibernate.validator.constraints.URL;

import javax.validation.constraints.NotNull;
import javax.validation.constraints.Pattern;

public class ShortenRequest {
    @NotNull
    @URL
    @Pattern(regexp = "^https?:.*")
    private String longUrl;

    public String getLongUrl() {
        return longUrl;
    }

    public void setLongUrl(String longUrl) {
        this.longUrl = longUrl;
    }
}