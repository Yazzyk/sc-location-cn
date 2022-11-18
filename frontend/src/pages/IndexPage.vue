<template>
  <q-page class="flex">
    <q-card class="main-box">
      <q-form>
        <p>StarCitizen文件夹路径<small>(如: E:\RSI\StarCitizen)</small>: <span style="font-size: large">{{ dirPath }}</span></p>

        <p>
          <q-btn color="primary" @click="ChooseFile">选择文件</q-btn>
        </p>
      </q-form>
    </q-card>
    <q-card class="desc">
      <h4 style="text-align: center">说明</h4>
    </q-card>
  </q-page>
</template>

<style>
.main-box {
  margin: 10px;
  width: calc(70% - 20px);
  height: calc(100%);
  padding: 10px;
}

.desc {
  position: relative;
  width: calc(30% - 20px);
  height: calc(100%);
  /*float: right;*/
  /*right: 10px;*/
  margin: 10px;
  padding: 10px;
}
</style>

<script>
import {defineComponent} from 'vue'
import {ChoseDir, GetConf} from "app/wailsjs/go/main/App";

export default defineComponent({
  name: 'IndexPage',
  data() {
    return {
      dirPath: ''
    }
  },
  mounted() {
    GetConf().then(data => {
      this.dirPath = data.SCPath
    })
  },
  methods: {
    ChooseFile() {
      ChoseDir().then(data => {
        this.dirPath = data
      })
    }
  }
})
</script>
