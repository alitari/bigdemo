package de.alexkrieg.messageapp.domain;

import java.util.Date;

public class Message {

    private String text;
    private Date creationTime;
    private String author;

    public String getText() {
        return text;
    }

    public String getAuthor() {
        return author;
    }

    public void setAuthor(String author) {
        this.author = author;
    }

    public Date getCreationTime() {
        return creationTime;
    }

    public void setCreationTime(Date creationTime) {
        this.creationTime = creationTime;
    }

    public void setText(String text) {
        this.text = text;
    }
}