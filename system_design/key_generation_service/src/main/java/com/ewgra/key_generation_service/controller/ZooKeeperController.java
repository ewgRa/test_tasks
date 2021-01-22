package com.ewgra.key_generation_service.controller;

import com.ewgra.key_generation_service.ZooKeeperConnector;
import lombok.extern.slf4j.Slf4j;
import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.ZooDefs;
import org.apache.zookeeper.ZooKeeper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

import java.util.HashMap;

@Slf4j
@RestController
public class ZooKeeperController {
    @Autowired
    ZooKeeperConnector zooKeeperConnector;

    @Value("${kgs.zookeeper.node}")
    private String nodePath;

    @PutMapping("/counter-based/init")
    public ResponseEntity<Object> init() throws Exception {
        ZooKeeper zoo = zooKeeperConnector.getConnection();

        if (zoo.exists(nodePath, null) != null) {
            throw new ResponseStatusException(HttpStatus.CONFLICT, "Node has been already created");
        }

        zoo.create(nodePath, "0".getBytes(), ZooDefs.Ids.OPEN_ACL_UNSAFE, CreateMode.PERSISTENT);
        log.debug("ZooKeeper node created");

        HashMap<String, Object> response = new HashMap<>();
        response.put("success", true);

        return new ResponseEntity<>(response, HttpStatus.CREATED);
    }
}