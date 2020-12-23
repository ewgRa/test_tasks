package com.ewgra.key_generation_service;

import lombok.extern.slf4j.Slf4j;
import org.apache.zookeeper.*;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.event.EventListener;
import org.springframework.stereotype.Component;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

@Slf4j
@Component
public class ZooKeeperConnector implements DisposableBean {
	private ZooKeeper zoo;

	@Value("${zooKeeperConnectString}")
	private final String connectString = null;

	private final int SESSION_TIMEOUT = 60000; // ms
	private final int CONNECTION_TIMEOUT = 5; // seconds

	public ZooKeeper getConnection() {
		return zoo;
	}

	@EventListener
	public void connect(ApplicationReadyEvent event) throws Exception {
		log.debug("Trying connect to ZooKeeper");

		CountDownLatch connectionLatch = new CountDownLatch(1);

		zoo = new ZooKeeper(connectString, SESSION_TIMEOUT, we -> {
			if (we.getState() == Watcher.Event.KeeperState.SyncConnected) {
				connectionLatch.countDown();
			}
		});

		if (!connectionLatch.await(CONNECTION_TIMEOUT, TimeUnit.SECONDS) || !zoo.getState().isConnected()) {
			log.debug("Failed connect to ZooKeeper");
			throw new Exception("Failed connect to ZooKeeper");
		}

		log.debug("Connected to ZooKeeper");
	}

	@Override
	public void destroy() throws InterruptedException {
		log.debug("Closing ZooKeeper connection");
		zoo.close();
	}
}