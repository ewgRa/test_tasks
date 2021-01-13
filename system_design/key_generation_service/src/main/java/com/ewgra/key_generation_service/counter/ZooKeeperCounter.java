package com.ewgra.key_generation_service.counter;

import com.ewgra.key_generation_service.ZooKeeperConnector;
import lombok.extern.slf4j.Slf4j;
import org.apache.zookeeper.*;
import org.apache.zookeeper.data.Stat;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.retry.annotation.CircuitBreaker;
import org.springframework.stereotype.Component;
import java.nio.charset.StandardCharsets;

// ZooKeeper counter-based number generator uses ZooKeeper to store number.
// If we will update the number in ZooKeeper for every increment - we will face race conditions.
// To reduce the chance of race condition probability, we acquire range with some capacity.
// We safe to increment number within such range in memory, without committing it to ZooKeeper.
@Slf4j
@Component
public class ZooKeeperCounter {
    @Value("${kgs.zookeeper.node}")
    private String nodePath;

    @Value("${kgs.counter.range.capacity}")
    private int rangeCapacity;

    private long rangeEnd;
    private long current;

    @Autowired
    private ZooKeeperConnector connector;

    @CircuitBreaker
    public synchronized long next() throws Exception {
        if (isRangeExhausted()) {
            acquireNewRange();
        } else {
            current++;
            log.debug("The number increased to {} / {}", current, rangeEnd);
        }

        return current;
    }

    private boolean isRangeExhausted() {
        return current == rangeEnd;
    }

    private void acquireNewRange() throws Exception {
        log.debug("Trying to acquire a new range");

        Stat stat = new Stat();
        byte[] data = connector.getConnection().getData(nodePath, null, stat);
        long rangeEndCandidate = Long.parseLong(new String(data, StandardCharsets.UTF_8)) + rangeCapacity;

        try {
            connector.getConnection().setData(nodePath, String.valueOf(rangeEndCandidate).getBytes(),
                    stat.getVersion());
        } catch (KeeperException e) {
            if (e.code() == KeeperException.Code.BADVERSION) {
                log.debug("Bad version while trying to acquire a new range. Retrying...");
                acquireNewRange();
                return;
            }

            log.debug("Fail to acquire a new range: " + e.getMessage());
            throw e;
        }

        rangeEnd = rangeEndCandidate;
        current = rangeEndCandidate - rangeCapacity + 1;

        log.debug("New range acquired {} {}", current, rangeEnd);
    }
}