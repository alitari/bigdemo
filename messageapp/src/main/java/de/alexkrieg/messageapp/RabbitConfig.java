package de.alexkrieg.messageapp;

import org.springframework.amqp.rabbit.connection.ConnectionFactory;

import org.springframework.amqp.core.Binding;
import org.springframework.amqp.core.BindingBuilder;
import org.springframework.amqp.core.Exchange;
import org.springframework.amqp.core.ExchangeBuilder;
import org.springframework.amqp.core.Queue;
import org.springframework.amqp.core.QueueBuilder;
import org.springframework.amqp.core.TopicExchange;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.amqp.support.converter.Jackson2JsonMessageConverter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RabbitConfig {
    public static final String QUEUE_MESSAGES = "messages-queue";
    public static final String EXCHANGE_MESSAGES = "messages-exchange";

    // @Bean
    // Queue messageQueue() {
    // return QueueBuilder.durable(QUEUE_MESSAGES).build();
    // }

    @Bean
    Exchange messageExchange() {
        return ExchangeBuilder.fanoutExchange(EXCHANGE_MESSAGES).build();
    }

    // @Bean
    // Binding binding(Queue messagesQueue, TopicExchange ordersExchange) {
    // return
    // BindingBuilder.bind(messagesQueue).to(ordersExchange).with(QUEUE_MESSAGES);
    // }

    @Bean
    public RabbitTemplate rabbitTemplate(final ConnectionFactory connectionFactory) {
        final RabbitTemplate rabbitTemplate = new RabbitTemplate(connectionFactory);
        rabbitTemplate.setMessageConverter(producerJackson2MessageConverter());
        return rabbitTemplate;
    }

    @Bean
    public Jackson2JsonMessageConverter producerJackson2MessageConverter() {
        return new Jackson2JsonMessageConverter();
    }
}