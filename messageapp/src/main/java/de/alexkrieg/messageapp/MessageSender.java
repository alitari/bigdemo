package de.alexkrieg.messageapp;

import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import de.alexkrieg.messageapp.domain.Message;

@Service
public class MessageSender {

    @Value("${spring.rabbitmq.exchange}")
    public String exchange = "test";

    private final RabbitTemplate rabbitTemplate;

    @Autowired
    public MessageSender(RabbitTemplate rabbitTemplate) {
        this.rabbitTemplate = rabbitTemplate;
    }

    public void sendMessage(Message message) {
        this.rabbitTemplate.convertAndSend(exchange, "", message);
    }
}