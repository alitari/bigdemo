package de.alexkrieg.messageapp;

import static org.junit.Assert.fail;

import java.util.Arrays;

import com.fasterxml.jackson.databind.ObjectMapper;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.test.context.TestPropertySource;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.web.client.RestTemplate;

import de.alexkrieg.messageapp.domain.NewsFeed;

@RunWith(SpringRunner.class)
@SpringBootTest
@TestPropertySource(locations = "classpath:test.properties")
public class NewsFeedTest {

	@Test
	public void test1() throws Exception {
		RestTemplate restTemplate = new RestTemplate();
		HttpHeaders headers = new HttpHeaders();
		headers.setAccept(Arrays.asList(MediaType.APPLICATION_JSON));
		headers.set("x-api-key", "f470ddf783c94759b7d49f13293273e9");
		headers.set("content-type", "application/json; charset=utf-16");
		headers.set("Accept-Encoding", "identity");
		HttpEntity<String> entity = new HttpEntity<String>("parameters", headers);

		String url = "https://newsapi.org/v2/everything?q=Karlsruhe";
		ResponseEntity<String> result = restTemplate.exchange(url, HttpMethod.GET, entity, String.class);
		String body = result.getBody();
		ObjectMapper mapper = new ObjectMapper();
		// NewsFeed obj = mapper.readValue(body, NewsFeed.class);
		// System.out.println(result.getBody());
		// fail(obj.toString());
		// NewsFeed newsFeed = restTemplate.getForObject(url, NewsFeed.class);

	}

}
