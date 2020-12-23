package com.ewgra.key_generation_service.count.provider;

import com.ewgra.key_generation_service.ZooKeeperConnector;
import lombok.extern.slf4j.Slf4j;
import org.apache.zookeeper.*;
import org.apache.zookeeper.data.Stat;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import java.nio.charset.StandardCharsets;

// ZooKeeper counter based generator uses ZooKeeper to store counter value.
// If we will update counter in ZooKeeper for every increment - we will have a frequent
// race condition case.
// To reduce this chance, we acquire range with some capacity and we safe to increment counter within such range
// in memory, without committing it to ZooKeeper. As a result, we reduce race condition probability for RANGE_CAPACITY times.
@Slf4j
@Component
public class ZooKeeperCountProvider {
    @Value("${zooKeeperNode}")
    private final String nodePath = null;

    @Value("${counterRangeCapacity}")
    private final int rangeCapacity = 100;

    private long rangeEnd;
    private long current;

    @Autowired
    private ZooKeeperConnector connector;

    public synchronized long next() throws Exception {
        if (isRangeExhausted()) {
            acquireNewRange();
        } else {
            current++;
            log.debug("Counter increased to {} / {}", current, rangeEnd);
        }

        return current;
    }

    private boolean isRangeExhausted() {
        return current == rangeEnd;
    }

    private void acquireNewRange() throws Exception {
        log.debug("Trying acquire new range");

        Stat stat = new Stat();
        byte[] data = connector.getConnection().getData(nodePath, null, stat);
        current = Long.parseLong(new String(data, StandardCharsets.UTF_8))+1;
        rangeEnd = current + rangeCapacity - 1;

        try {
            connector.getConnection().setData(nodePath, String.valueOf(rangeEnd).getBytes(), stat.getVersion());
        } catch (KeeperException e) {
            if (e.code() == KeeperException.Code.BADVERSION) {
                log.debug("Bad version while trying to acquire new range. Retrying...");
                acquireNewRange();
                return;
            }

            log.debug("Fail to acquire new range: " + e.getMessage());
            throw e;
        }

        log.debug("New range acquired {} {}", current, rangeEnd);
    }
}