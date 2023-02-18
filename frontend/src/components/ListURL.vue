<template>
  <div class="list-container">
    <h6>&emsp;</h6>
    <table>
      <thead>
        <tr>
          <th>Original URL</th>
          <th>Short URL</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(url, index) in urls" :key="index">
          <td>{{ url.original }}</td>
          <td>{{ url.short }}</td>
        </tr>
      </tbody>
    </table>

    <h6> </h6>
    <button @click="refreshURLs">Refresh</button>
    &emsp;
    <button @click="hideURLs">Hide</button>
    &emsp;
    <button @click="clearURLs">Clear</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props: {
    urls: Array,
  },
  data() {
    return {}
  },
  methods: {
    async refreshURLs() {
      this.$emit("refresh");
    },
    async hideURLs() {
      this.$emit("hide");
    },
    async clearURLs() {
      await axios.delete("http://localhost:8888/api/urls")
      this.$emit("refresh"); // 觸發 App 的 refreshList 方法
    },
  }
};
</script>


<style scoped>
.list-container {
  text-align: center;
}

table {
  width: 40%;
  margin: 0 auto;
  /* border-collapse: collapse; */
}

th,
td {
  padding: 8px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

th {
  background-color: #ccebf7;
}
</style>