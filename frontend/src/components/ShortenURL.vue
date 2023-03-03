<template>
  <div>
    <div> &emsp;  </div>
    <div>
      <label for="url-input">Enter URL : </label>
      <input type="url" id="url-input" v-model="url" required>
      &emsp;<button class="button-10" @click="shortenURL">Shorten</button>
    </div>
    <div v-if="shortUrl">
      <h4>Short URL: <a :href="shortUrl" target="_blank">{{ shortUrl }}</a> </h4>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ShortenUrl',
  data() {
    return {
      url: '',
      shortUrl: '',
    }
  },
  methods: {
    async shortenURL() {
      const response = await axios.post('http://localhost:8888/api/shorten', {
        original: this.url
      })
      this.shortUrl = response.data.short
      this.$emit("shorten-success"); // 觸發 App 的 refreshList 方法
    }
  }
}
</script>


<style scoped>
.button-10 {
  flex-direction: column;
  align-items: center;
  padding: 6px 14px;
  font-family: -apple-system, BlinkMacSystemFont, 'Roboto', sans-serif;
  border-radius: 6px;
  border: none;

  color: #fff;
  background: linear-gradient(180deg, #4B91F7 0%, #367AF6 100%);
  background-origin: border-box;
  box-shadow: 0px 0.5px 1.5px rgba(54, 122, 246, 0.25), inset 0px 0.8px 0px -0.25px rgba(255, 255, 255, 0.2);
  user-select: none;
  -webkit-user-select: none;
  touch-action: manipulation;
}

.button-10:hover {
  box-shadow: inset 0px 0.8px 0px -0.25px rgba(255, 255, 255, 0.2), 0px 0.5px 1.5px rgba(54, 122, 246, 0.25), 0px 0px 0px 3.5px rgba(58, 108, 217, 0.5);
  outline: 0;
}
</style>