<template>
  <div class="dashboard">
    <div class="left-pane">
      <NavigationComponent />
    </div>
    <div class="right-pane">
      <!-- Main content goes here -->
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import NavigationComponent from './components/NavigationComponent.vue'
export default {
  components: {
    NavigationComponent,
  },
}

</script>

<style>
#app {
  font-family: monospace;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

html, body {
  margin:0;
  padding:0;
  height:100%;
}


.dashboard {
  display: flex;
  height: 100vh; /* 100% of the viewport height */
}

.left-pane {
  width: 18%;
  background-color: #f0f0f0;
  padding: 20px;
  box-sizing: border-box;
  position: sticky;
  height: 100vh;
  overflow-y: auto; /* Enable scrolling for the left pane */
  margin-right: 1px;
}

.right-pane {
  flex: 1;
  padding: 20px;
  box-sizing: border-box;
  background-color: #e4e4e4;
  overflow-y: auto;
}

</style>
