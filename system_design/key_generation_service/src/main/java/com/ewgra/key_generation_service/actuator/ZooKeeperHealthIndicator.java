package com.ewgra.key_generation_service.actuator;

import com.ewgra.key_generation_service.ZooKeeperConnector;
import org.apache.zookeeper.ZooKeeper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.stereotype.Component;

@Component
public class ZooKeeperHealthIndicator implements HealthIndicator {

    @Autowired
    ZooKeeperConnector zooKeeperConnector;

    @Override
    public Health health() {
        ZooKeeper.States state = zooKeeperConnector.getConnection().getState();

        if (!state.isConnected()) {
            return Health.down().build();
        }
        return Health.up().build();
    }

}