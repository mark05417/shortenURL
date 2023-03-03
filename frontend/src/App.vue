<template>
  <div id="app">
    <div id="header"> <img alt="URL logo" src="./assets/icon.png"> </div>
    <div style="display: flex;">
      <div id="sidebar_left"></div>
      <div id="content">
        <ShortenURL @shorten-success="refreshList" msg="Shorten"/>
        <ListURL @refresh="refreshList" @clear="refreshList" :urls="urls" msg="Table"/>
      </div>
    </div>
    <div id="footer"> All Copyright Reserved, Mark Hsu, 2023 </div>
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
    async refreshList() {
      const res = await axios.get("http://localhost:8888/api/urls");
      this.urls = res.data;
    },
  }
}
</script>

<style>
html, body {
  height: 100%;
  margin: 0;
}
#app {
  display: flex;
  flex-direction: column;
  height: 100%;
}

#header{
  background-color: #ffffff;
  height: 80px;
  text-align: center;
  line-height: 80px;
}
#sidebar_left{
  background-color: #e4fce4;
  text-align: center;
  width: 300px;
  flex-shrink: 0;
  overflow-y: auto;
}
#content{
  background-color: #fcf0d0;
  text-align: center;
  flex-grow: 1;
  overflow-y: auto;
}
#footer{
  background-color: #e4e4e4;
  height: 80px;
  text-align: center;
  line-height: 80px;
}
</style>
