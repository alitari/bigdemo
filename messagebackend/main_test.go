package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockDataAccess struct {
	data  []string
	count int
	keys  []string
}

func (r *mockDataAccess) getData(key string) (data []string, err error) {
	var er error
	if r.data == nil {
		er = errors.New("error getData")
	}
	return r.data, er
}

func (r *mockDataAccess) deleteData(key string) error {
	var er error
	if r.data == nil {
		er = errors.New("error getData")
	}
	return er
}

func (r *mockDataAccess) deleteAllData() error {
	var er error
	if r.data == nil {
		er = errors.New("error getData")
	}
	return er
}

func (r *mockDataAccess) getDataCount() (count int, err error) {
	var er error
	if r.count == -1 {
		er = errors.New("error getDataCount")
	}
	return r.count, er
}
func (r *mockDataAccess) queryData(query string) (keys []string, err error) {
	var er error
	if r.keys == nil {
		er = errors.New("error queryData")
	}
	return r.keys, er
}

func init() {
	dataAcc = &mockDataAccess{data: nil, count: -1, keys: nil}
	go serv()
	time.Sleep(500 * time.Millisecond)
}

func mockData() *mockDataAccess {
	return dataAcc.(*mockDataAccess)
}

func redisData() *redisDataAccess {
	return dataAcc.(*redisDataAccess)
}

func TestMock_message(t *testing.T) {
	require := require.New(t)
	key := "key"
	text := "myText"
	author := "myAuthor"
	unixTime := time.Now().Unix()
	instant := fmt.Sprintf("{ \"epochSecond\": %d, \"nano\": %d }", unixTime, 589580000)
	mockData().data = []string{text, author, instant}

	mess := message(key)

	//fmt.Printf("message { id: %s, text: %s, author: %s, creationTime: { epochSecond: %d 1539859457, nano: %d 589580000} }\n", mess.ID, mess.Text, mess.Author, mess.)
	require.Equal(key, mess.ID)
	require.Equal(text, mess.Text)
	require.Equal(author, mess.Author)
	require.Equal(strconv.FormatInt(unixTime, 10), mess.CreationTime)
}

func TestMock_messageNotFound(t *testing.T) {
	require := require.New(t)
	key := "key"
	mockData().data = []string{}

	mess := message(key)

	//fmt.Printf("message { id: %s, text: %s, author: %s, creationTime: %d }\n", mess.ID, mess.Text, mess.Author, mess.CreationTime)
	require.Equal(key, mess.ID)
	require.Equal("not found", mess.Text)
	require.Equal("no", mess.Author)
}

func TestMock_httpgetMessage(t *testing.T) {
	require := require.New(t)

	key := "key2"
	text := "myText2"
	author := "myAuthor2"
	unixTime := time.Now().Unix()
	instant := fmt.Sprintf("{ \"epochSecond\": %d, \"nano\": %d }", unixTime, 589580000)
	mockData().data = []string{text, author, instant}

	url := fmt.Sprintf("http://localhost:8000/messages/%s", key)
	resp, err := http.Get(url)
	require.Nil(err)
	require.NotNil(resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	require.Nil(err)
	require.NotNil(body)
	require.Equal(http.StatusOK, resp.StatusCode)

	fmt.Printf("body: %v \n", string(body))
	expectedBody := fmt.Sprintf("{\"id\":\"%s\",\"text\":\"%s\",\"author\":\"%s\",\"creationTime\":\"%d\"}\n", key, text, author, unixTime)
	require.Equal(expectedBody, string(body))
}

func Test_redis(t *testing.T) {
	require := require.New(t)
	os.Setenv("REDIS_HOST", "localhost")
	os.Setenv("REDIS_PASSWORD", "redis")
	dataAcc = &redisDataAccess{redisClient: createRedisClient()}

	key := "key3"
	text := "myText3"
	author := "myAuthor3"

	unixTime := time.Now().Unix()
	instant := fmt.Sprintf("{ \"epochSecond\": %d, \"nano\": %d }", unixTime, 589580000)

	redisData().redisClient.Del(key)
	redisData().redisClient.RPush(key, text, author, instant)

	mess := message(key)

	//fmt.Printf("message { id: %s, text: %s, author: %s, creationTime: %d }\n", mess.ID, mess.Text, mess.Author, mess.CreationTime)
	require.Equal(key, mess.ID)
	require.Equal(text, mess.Text)
	require.Equal(author, mess.Author)
	require.Equal(strconv.FormatInt(unixTime, 10), mess.CreationTime)
}
