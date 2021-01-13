package com.ewgra.key_generation_service.controller;

import com.ewgra.key_generation_service.ZooKeeperConnector;
import com.ewgra.key_generation_service.counter.ZooKeeperCounter;
import com.ewgra.key_generation_service.generator.CounterBasedGenerator;
import lombok.extern.slf4j.Slf4j;
import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.ZooDefs;
import org.apache.zookeeper.ZooKeeper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Slf4j
@RestController
public class CounterBasedController {
    @Autowired
    ZooKeeperCounter counter;

    @Autowired
    CounterBasedGenerator generator;

    @Autowired
    ZooKeeperConnector zooKeeperConnector;

    @Value("${kgs.zookeeper.node}")
    private String nodePath;

    @PutMapping("/counter-based/init")
    public ResponseEntity<Object> init() {
        ZooKeeper zoo = zooKeeperConnector.getConnection();

        try {
            if (zoo.exists(nodePath, null) != null) {
                HashMap<String, Object> response = new HashMap<>();
                response.put("success", false);

                return new ResponseEntity<>(response, HttpStatus.CONFLICT);
            }
        } catch (Exception e) {
            log.error("Error when getting stat for node: " + e.getMessage());

            return internalServerError("Internal server error when getting stat for node");
        }

        try {
            zoo.create(nodePath, "0".getBytes(), ZooDefs.Ids.OPEN_ACL_UNSAFE, CreateMode.PERSISTENT);
        } catch (Exception e) {
            log.error("Error when creating node: " + e.getMessage());

            return internalServerError("Internal server error when creating node");
        }

        log.debug("ZooKeeper node created");

        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);

        return new ResponseEntity<>(response, HttpStatus.CREATED);
    }

    @GetMapping("/counter-based/key")
    public ResponseEntity<Object> key() {
        long next;

        try {
            next = counter.next();
        } catch (Exception e) {
            log.error("Error when getting next key: " + e.getMessage());

            return internalServerError("Internal server error when getting next counter value");
        }

        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);
        response.put("key", generator.generate(next));

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    private ResponseEntity<Object> internalServerError(String message) {
        HashMap<String, Object> response = new HashMap<>();
        response.put("success", false);
        response.put("message", message);

        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}