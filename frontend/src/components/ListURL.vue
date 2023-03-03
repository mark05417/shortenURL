<template>
  <div class="list-container">
    <h6>&emsp;</h6>
    <table>
      <thead>
        <tr>
          <th>Original URL</th>
          <th>Short URL</th>
          <th>Count</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(url, index) in urls" :key="index">
          <td>{{ url.original }}</td>
          <td>{{ url.short }}</td>
          <td class="td-center">{{ url.count }}</td>
          <td class="td-center"><button class="delete-btn" @click="deleteURL(url.short)">X</button></td>
        </tr>
      </tbody>
    </table>

    <h6>
    <button class="button-10" @click="refreshURLs">Refresh</button>
    &emsp;
    <button class="button-10" @click="clearURLs">Clear</button>
    </h6>
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
    async clearURLs() {
      await axios.delete("http://localhost:8888/api/urls")
      this.$emit("refresh"); // 觸發 App 的 refreshList 方法
    },
    async deleteURL(shortURL) {
      var deleteShort = shortURL.split('/',2)[1]
      await axios.delete(`http://localhost:8888/api/${deleteShort}`);
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
.td-center {
  text-align: center;
}
.delete-btn {
  background: none;
  border: none;
  color: rgb(0, 0, 0);
  cursor: pointer;
  font-size: 1.0em;
  font-weight: bold;
  padding: 0;
  text-align: center;
  width: 1.0em;
}

.delete-btn:hover {
  color: rgb(255, 0, 0);
}

.button-10 {
  flex-direction: column;
  align-items: center;
  padding: 6px 14px;
  font-family: -apple-system, BlinkMacSystemFont, 'Roboto', sans-serif;
  border-radius: 6px;
  border: none;

  color: #fff;
  background: linear-gradient(180deg,  #4B91F7 0%, #367AF6 100%);
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