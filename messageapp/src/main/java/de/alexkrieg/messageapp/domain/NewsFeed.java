package de.alexkrieg.messageapp.domain;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class NewsFeed {

    private String[] articles;
    private String status;
    private int totalResults;

    /**
     * @return the articles
     */
    public String[] getArticles() {
        return articles;
    }

    /**
     * @return the totalResults
     */
    public int getTotalResults() {
        return totalResults;
    }

    /**
     * @param totalResults the totalResults to set
     */
    public void setTotalResults(int totalResults) {
        this.totalResults = totalResults;
    }

    /**
     * @return the status
     */
    public String getStatus() {
        return status;
    }

    /**
     * @param status the status to set
     */
    public void setStatus(String status) {
        this.status = status;
    }

    /**
     * @param articles the articles to set
     */
    public void setArticles(String[] articles) {
        this.articles = articles;
    }

}