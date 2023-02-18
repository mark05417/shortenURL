<template>
  <div id="header">
    <img alt="URL logo" src="./assets/icon.png">
  </div>

  <div id="sidebar_left">
  </div>

  <div id="content">
    <ShortenURL @shorten-success="refreshList" msg="Shorten"/>
    <ListURL @refresh="refreshList" @hide="hideList" @clear="refreshList" :urls="urls" msg="Table"/>
  </div>

  <div id="footer">
    All Copyright Reserved, Mark Hsu, 2023
  </div>
</template>

<script>
import ShortenURL from './components/ShortenURL.vue'
import ListURL from './components/ListURL.vue'
import axios from 'axios'

export default {
  name: 'App',
  components: {
    ShortenURL,
    ListURL
  },
  data() {
    return {
      urls: [],
    };
  },
  methods: {
    async hideList() {
      this.urls = [];
    },
    async refreshList() {
      const res = await axios.get("http://localhost:8888/api/urls");
      this.urls = res.data;
    },
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* background-color: cadetblue; */
  margin-top: 60px;
}

#header{
  background-color:#ffffff;
  height:80px;
  text-align:center;
  line-height:80px;
}
#sidebar_left{
  margin-top: 20px;
  background-color:#e4fce4;
  text-align:center;
  width:300px;
  height:500px;
  float:left;
}
#content{
  margin-top: 20px;
  background-color:#fcf0d0;
  text-align:center;
  height:500px;
}
#footer{
  clear:both;
  float:unset;
  background-color:#e4e4e4;
  height:80px;
  text-align:center;
  line-height:80px;
}
</style>
