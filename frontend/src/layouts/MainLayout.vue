<template>
  <q-layout view="hHh lpr fFf">

    <q-header class="window-header text-white" height-hint="98">
      <q-toolbar>
        <q-toolbar-title style="--wails-draggable:drag">
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg">
          </q-avatar>
          星河避难所Lite
        </q-toolbar-title>

        <q-btn-group class="window-btn-group" rounded>
<!--          <font-awesome-icon icon="fa-solid fa-circle-xmark" />-->
          <q-btn icon="minimize" @click="smallWindow"></q-btn>
          <q-btn :icon="fullIcon" @click="fullWindow"></q-btn>
          <q-btn class="close-btn" icon="highlight_off" @click="closeWindow"></q-btn>
        </q-btn-group>
      </q-toolbar>

      <q-tabs align="left">
        <q-route-tab to="/" label="汉化" />
        <q-route-tab to="/" label="更新" />
        <q-route-tab to="/" label="设置" />
        <q-route-tab to="/" label="关于" />
      </q-tabs>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>

  </q-layout>
</template>

<style>
.close-btn:hover {
  background-color: red;
}
.window-header {
  z-index: 999;
}
/*.window-btn-group {*/
/*  background-color: rgb(0,47,167);*/
/*  color: black;*/
/*}*/
</style>

<script>
import { defineComponent, ref } from 'vue'
import EssentialLink from 'components/EssentialLink.vue'

const linksList = [

]

export default defineComponent({
  name: 'MainLayout',
  data() {
    return {
      fullIcon: 'open_in_full'
    }
  },
  components: {
    // EssentialLink
  },
  methods: {
    closeWindow() {
      window.runtime.Quit()
    },
    fullWindow() {
      console.log('全屏点击')
      window.runtime.WindowIsMaximised().then((t) => {
        if (t) {
          window.runtime.WindowUnmaximise()
          this.fullIcon = 'open_in_full'
        } else {
          window.runtime.WindowMaximise()
          this.fullIcon = 'close_fullscreen'
        }
      })
    },
    smallWindow() {
      window.runtime.WindowMinimise()
    }
  }
})
</script>
