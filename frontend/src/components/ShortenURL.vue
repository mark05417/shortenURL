<template>
  <div>
    <div> &emsp;  </div>
    <div>
      <label for="url-input">Enter URL : </label>
      <input type="url" id="url-input" v-model="url" required>
      <button @click="shortenURL">Shorten</button>
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
</style>