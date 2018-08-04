package de.alexkrieg.messageapp;

import org.springframework.stereotype.Component;
import java.util.List;
import java.util.Arrays;

@Component
class MessageService {
    public List<String> getProducts() {
        return Arrays.asList("iPad", "iPod", "iPhone");
    }
}