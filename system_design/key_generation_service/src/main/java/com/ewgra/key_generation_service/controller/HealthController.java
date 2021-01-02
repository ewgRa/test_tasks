package com.ewgra.key_generation_service.controller;

import com.ewgra.key_generation_service.ZooKeeperConnector;
import lombok.extern.slf4j.Slf4j;
import org.apache.zookeeper.ZooKeeper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Slf4j
@RestController
public class HealthController {
    @Autowired
    ZooKeeperConnector zooKeeperConnector;

    @GetMapping("/health/liveness")
    public ResponseEntity<Object> liveness() {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    @GetMapping("/health/readiness")
    public ResponseEntity<Object> readiness() {
        HashMap<String, Object> response = new HashMap<>();
        ZooKeeper.States state = zooKeeperConnector.getConnection().getState();
        response.put("success", state.isConnected());

        return new ResponseEntity<>(response, state.isConnected() ? HttpStatus.OK : HttpStatus.SERVICE_UNAVAILABLE);
    }
}