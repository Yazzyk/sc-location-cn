<template>
  <q-page class="flex">
   <q-card class="main-box">
     <q-form>
       <p>选择星际公民本地路径: <q-btn color="primary">选择文件</q-btn></p>
       <p>{{dirPath}}</p>
     </q-form>
   </q-card>
  </q-page>
</template>

<style>
.main-box {
  margin: 10px;
  width: 40%;
  height: 400px;
  padding: 10px;
}
</style>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'IndexPage',
  data(){
    return {
      dirPath: ''
    }
  }
})
</script>
