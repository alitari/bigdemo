<template>
  <div id="app">
    <img src="./assets/skaffold.png" height="100" width="150" align="center">
    <h1>{{ msg }}</h1>
    <h2><a href="/messageapp/form" >Write message</a></h2>
    <hr>

    <!-- <a href="#" v-on:click="fetchMessages('boy')">Fetch messages for word boy</a></h2> -->
    <h2>
      Search by word from {{messageCount}} messages<br>
    <input id="searchInpt" value=""><button id="searchBtn" v-on:click="fetchMessages()">Search</button>

    </h2>

    <table id="messages" align="center">
      <thead>
        <tr>
          <th width="5%">Id</th>
          <th width="20%">Time</th>
          <th width="15%">Author</th>
          <th width="60%">Text</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="m in messages">
          <td>{{m.id}}</td>
          <td>{{m.creationTime}}</td>
          <td>{{m.author}}</td>
          <td>{{m.text}}</td>
        </tr>
      </tbody>
    </table>
  
    
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "app",
  data() {
    return {
      msg: "Big Skaffold DEMO",
      messages: [],
      messageCount: 0
    };
  },
  methods: {
    fetchMessages() {
      var input = document.getElementById("searchInpt");
      var baseUrl = "http://" + window.location.hostname + "/messages";
      //console.log(baseUrl);

      axios
        .get(baseUrl.concat("?word=").concat(input.value))
        .then(response => {
          this.messages = [];
          for (var i = 0; i < response.data.length; i++) {
            var newDate = new Date();
            newDate.setTime(response.data[i].creationTime * 1000);
            this.messages.push({
              id: response.data[i].id,
              text: response.data[i].text,
              author: response.data[i].author,
              creationTime: newDate.toUTCString()
            });
            console.log(this.messages[i]);
          }
        })
        .catch(error => {
          console.log(error);
          this.messages = "error";
        });
      axios
        .get(baseUrl.concat("/count"))
        .then(response => {
          this.messageCount = response.data.text;
        })
        .catch(error => {
          console.log(error);
        });
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

table,
th,
td {
  border: 1px solid black;
  border-collapse: collapse;
}
</style>
