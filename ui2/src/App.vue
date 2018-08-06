<template>
  <div id="app">
    <img src="./assets/logo.png" height="100" width="150" align="left">
    <img src="./assets/skaffold.png" height="100" width="150" align="right">
    <h1>{{ msg }}</h1>
    <h2><a href="/messageapp" >Message Application</a></h2>
    <hr>

    <!-- <a href="#" v-on:click="fetchMessages('boy')">Fetch messages for word boy</a></h2> -->
    <h2>
    <input id="searchInpt" value="a word..."><button id="searchBtn" v-on:click="fetchMessages()">Search</button>

    </h2>

    <table id="messages" align="center">
      <thead>
        <tr>
          <th width="20%">Time</th>
          <th width="20%">User</th>
          <th width="60%">Text</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="m in messages">
          <td>?</td>
          <td>?</td>
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
      msg: "Big Microservices DEMO",
      messages: [],
      goURL: "http://bigdemo.com/messages"
    };
  },
  methods: {
    fetchMessages() {
      var input = document.getElementById("searchInpt");

      axios
        .get(this.goURL.concat("?word=").concat(input.value))
        .then(response => {
          this.messages = response.data;
        })
        .catch(error => {
          console.log(error);
          this.messages = "error";
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
