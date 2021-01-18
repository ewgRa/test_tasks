package com.ewgra.key_generation_service.controller;

import com.ewgra.key_generation_service.ZooKeeperConnector;
import org.apache.zookeeper.ZooKeeper;
import org.apache.zookeeper.data.Stat;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.web.servlet.MockMvc;

import java.nio.charset.StandardCharsets;

import static org.assertj.core.api.Assertions.assertThat;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
@AutoConfigureMockMvc
@ActiveProfiles("test")
public class CounterBasedControllerTest {
    @Value("${kgs.zookeeper.node}")
    String zooKeeperNode;

    @Autowired
    ZooKeeperConnector zooKeeperConnector;

    @Autowired
    private MockMvc mockMvc;

    @BeforeEach
    public void setUp() throws Exception {
        ZooKeeper zoo = zooKeeperConnector.getConnection();
        Stat stat = zoo.exists(zooKeeperNode, null);

        if (stat != null) {
            zoo.delete(zooKeeperNode, stat.getVersion());
        }

        this.mockMvc.perform(put("/counter-based/init")).andExpect(status().isCreated());
    }

    @Test
    public void getKey() throws Exception {
        String[] expectedResponses = new String[] { "{\"success\":true,\"key\":\"b\"}",
                "{\"success\":true,\"key\":\"c\"}", "{\"success\":true,\"key\":\"d\"}",
                "{\"success\":true,\"key\":\"e\"}", };

        for (String expectedResponse : expectedResponses) {
            this.mockMvc.perform(get("/counter-based/key")).andExpect(status().isOk())
                    .andExpect(content().string(expectedResponse));
        }

        Stat stat = new Stat();
        byte[] data = zooKeeperConnector.getConnection().getData(zooKeeperNode, null, stat);
        String number = new String(data, StandardCharsets.UTF_8);

        assertThat(number).isEqualTo("6");
    }
}