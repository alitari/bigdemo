package de.alexkrieg.messageapp.domain;

import java.time.Instant;

public class Message {

    private String text;
    private Instant creationTime;
    private String author;

    public String getText() {
        return text;
    }

    /**
     * @return the creationTime
     */
    public Instant getCreationTime() {
        return creationTime;
    }

    /**
     * @param creationTime the creationTime to set
     */
    public void setCreationTime(Instant creationTime) {
        this.creationTime = creationTime;
    }

    public String getAuthor() {
        return author;
    }

    public void setAuthor(String author) {
        this.author = author;
    }

    public void setText(String text) {
        this.text = text;
    }
}