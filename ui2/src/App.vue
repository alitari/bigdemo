<template>
  <div id="app">
    <img src="./assets/logo.png" height="100" width="150" align="left">
    <img src="./assets/skaffold.png" height="100" width="150" align="right">
    <h1>{{ msg }}</h1>
    <table align="center">
      <tr>
        <th width="20%" >Backend</th>
        <th width="20%">URL</th>
        <th width="60%">Response</th>
      </tr>
      <tr>
        <td><a href="#" v-on:click="getSpring()">Call Spring</a></td>
        <td>{{ springURL}}</td>
        <td>{{ springresult}}</td>
      </tr>
      <tr>
        <td><a href="#" v-on:click="getGo()">Call Go</a></td>
        <td>{{ goURL}}</td>
        <td>{{ goresult}}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import axios from 'axios'


export default {
  name: "app",
  data() {
    return {
      msg: "Big Microservices DEMO",
      springresult: null,
      springURL : 'http://bigdemo.com/productapp/',
      goresult: null,
      goURL : 'http://bigdemo.com/people'
    };
  },
  methods: {
    getSpring() {
      axios.get(this.springURL)
            .then(response => {
              this.springresult = response.data
            })
            .catch(error => {
              console.log(error)
              this.springresult = "error"
            })
    },
    getGo() {
      axios.get(this.goURL)
            .then(response => {
              this.goresult = response.data
            })
            .catch(error => {
              console.log(error)
              this.goresult = "error"
            })
    },
  }
}
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

table, th, td {
    border: 1px solid black;
    border-collapse: collapse;
}
</style>
